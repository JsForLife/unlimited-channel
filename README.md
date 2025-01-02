# chann ![example workflow](https://github.com/golang-design/chann/actions/workflows/chann.yml/badge.svg) ![](https://changkun.de/urlstat?mode=github&repo=golang-design/chann)

a unified channel package in Go

```go
import "golang.design/x/chann"
```

This package requires Go 1.18.

## Basic Usage

Different types of channels:

```go
ch := chann.New[int]()                  // unbounded, capacity unlimited
ch := chann.New[func()](chann.Cap(0))   // unbufferd, capacity 0
ch := chann.New[string](chann.Cap(100)) // buffered,  capacity 100
```

Send and receive operations:

```go
ch.In() <- 42
println(<-ch.Out()) // 42
```

Close operation:

```go
ch.Close()
```

Channel properties:

```go
ch.Len() // the length of the channel
ch.Cap() // the capacity of the channel
```

See https://golang.design/research/ultimate-channel for more details of
the motivation of this abstraction.

## License


MIT | &copy; 2021 The [golang.design](https://golang.design) Initiative Authors, written by [Changkun Ou](https://changkun.de).

# unlimited-channel
A channel with unlimited buffer in golang

```go
import "unlimited-channel"
```

This package requires Go 1.18.

## Basic Usage

To use this package, first import it in your Go code:


Then, you can create a new unlimited channel by calling the NewUnLimitedChannel function:

```go
    in, out := NewUnLimitedChannel[int]()
	var wg sync.WaitGroup
	wg.Add(2)

    // producer
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

    // consumer
	go func() {
		defer wg.Done()
		for val := range out {
			fmt.Println(val)
		}
	}()

	wg.Wait()
```

## License
This package is licensed under the MIT License.