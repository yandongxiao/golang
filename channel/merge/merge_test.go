package merge

import (
	"fmt"
	"reflect"
	"testing"
)

func write(begin, end, step int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- begin:
				begin += step
			}

			if begin >= end {
				close(ch)
				break
			}
		}
	}()
	return ch
}

func aggragate(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for i := range inputs {
			for v := range inputs[i] {
				output <- v
			}
		}
		close(output)
	}()
	return output
}

func Test_merge(t *testing.T) {
	getClosedChan := func() <-chan int {
		ch := make(chan int)
		close(ch)
		return ch
	}

	type args struct {
		ch1 <-chan int
		ch2 <-chan int
	}
	tests := []struct {
		name string
		args args
		want <-chan int
	}{
		{name: "channel is nil", args: args{ch1: nil, ch2: nil}, want: getClosedChan()},
		{name: "ch1 is nil", args: args{ch1: nil, ch2: write(0, 10, 2)}, want: write(0, 10, 2)},
		{name: "ch2 is nil", args: args{ch1: write(0, 10, 2), ch2: nil}, want: write(0, 10, 2)},
		{name: "ch1 is bigger", args: args{ch1: write(10, 20, 2), ch2: write(0, 10, 2)}, want: write(0, 20, 2)},
		{name: "ch2 is bigger", args: args{ch1: write(0, 10, 2), ch2: write(10, 20, 2)}, want: write(0, 20, 2)},
		{name: "ch1 is longer", args: args{ch1: write(0, 20, 2), ch2: write(1, 11, 2)}, want: aggragate(write(0, 10, 1), write(10, 20, 2))},
		{name: "ch2 is longer", args: args{ch1: write(0, 10, 2), ch2: write(1, 19, 2)}, want: aggragate(write(0, 9, 1), write(9, 19, 2))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.ch1, tt.args.ch2)
			var gotDatas []int
			for v := range got {
				fmt.Print(v)
				gotDatas = append(gotDatas, v)
			}

			fmt.Println()

			var expectDatas []int
			for v := range tt.want {
				fmt.Print(v)
				expectDatas = append(expectDatas, v)
			}

			if !reflect.DeepEqual(gotDatas, expectDatas) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
