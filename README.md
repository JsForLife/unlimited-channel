# unlimited-channel
A channel with unlimited buffer in golang

```go
import "unlimited-channel"
```

This package requires Go 1.18.

## Basic Usage
Init channel
```go
ch := NewUnLimitedChannel[int]()
```
Send and receive message
```go
ch.In() <- 1
println(<-ch.Out)
```
Close
```go
ch.Close()
```

## License
This package is licensed under the MIT License.