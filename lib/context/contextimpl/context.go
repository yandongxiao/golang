package contextimpl

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

var Canceled = errors.New("context canceled")
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }

type emptyContext int

func (emptyContext) Deadline() (deadline time.Time, ok bool) { return }
func (emptyContext) Done() <-chan struct{}                   { return nil }
func (emptyContext) Err() error                              { return nil }
func (emptyContext) Value(key interface{}) interface{}       { return nil }

var (
	todo       = new(emptyContext)
	background = new(emptyContext)
)

func TODO() Context { return todo }

func Background() Context { return background }

type CancelFunc func()

type cancelContext struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelContext) Value(key interface{}) interface{}       { return ctx.Value(key) }
func (ctx *cancelContext) Deadline() (deadline time.Time, ok bool) { return ctx.Deadline() }
func (ctx *cancelContext) Done() <-chan struct{}                   { return ctx.done }
func (ctx *cancelContext) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelContext{
		Context: parent,
		done:    make(chan struct{}),
	}

	cancel := func() {
		ctx.cancel(Canceled)
	}

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.done:
		}
	}()

	return ctx, cancel
}

func (ctx *cancelContext) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}

type timeoutContext struct {
	Context
	deadline time.Time
}

func (ctx *timeoutContext) Deadline() (deadline time.Time, ok bool) { return ctx.deadline, true }

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	t := time.AfterFunc(d.Sub(time.Now()), func() {
		cctx.(*cancelContext).cancel(DeadlineExceeded)
	})

	stop := func() {
		t.Stop()
		cancel()
	}

	return &timeoutContext{
		Context:  cctx,
		deadline: d,
	}, stop
}

type valueContext struct {
	Context
	key, value interface{}
}

func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("key can not be nil")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key must be comparable")
	}

	return &valueContext{
		Context: parent,
		key:     key,
		value:   val,
	}
}

func (ctx *valueContext) Value(key interface{}) interface{} {
	if ctx.key == key {
		return ctx.value
	}
	return ctx.Context.Value(key)
}
