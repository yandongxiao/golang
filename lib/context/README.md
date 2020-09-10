context package 为我们提供了以下关键能力：
- cancellation
- cancellation propagation
    - cancellation propagation 的实现依赖：
    - Context对象在函数之间的传递
    - 所有的Context对象构成了Tree Of Dependency
    - 父节点的Cancel动作会递归地传播给所有子节点
- pass value
    - 谨慎使用WithValue，要求：
    - value必须是Request级别，比如RequestID
    - 即使是Request级别的value，如果它很重要，你应该通过参数传递。提高代码的可读性

### WithTimeout

```go
func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 1秒钟后cancel会自动执行
	// 可重入、协程安全的函数
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("work done")
	}

	// Output:
	// context deadline exceeded
}
```

### WithValue

```go
func ExampleWithValue() {
	ctx := context.WithValue(context.Background(), 1, 100)
	fmt.Println(ctx.Value(1))

	ctx2 := context.WithValue(ctx, 1, 200)
	fmt.Println(ctx.Value(1))
	fmt.Println(ctx2.Value(1))

	// Output:
	// 100
	// 100
	// 200
}
```