# unlimited-channel
A channel with unlimited buffer in golang

```go
import "unlimited-channel"
```

This package requires Go 1.18.

## Basic Usage
Init channel
```go
in, out := NewUnLimitedChannel[int]()
```
Send and receive message
```go
in <- 1
println(<-out)
```

## License
This package is licensed under the MIT License.