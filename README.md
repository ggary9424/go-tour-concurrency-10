# A Tour of Go - Concurrency 10

This repository is for me to practice `Go`(golang). I wrote the code for the exercise for [A Tour of Go](https://tour.golang.org/concurrency/10), and did the easy benchmark to understand the efficiency that `Goroutines` brings us.

## Execution

```bash
$ go run ./main.go
```

## Benchmark Results

```bash
$ DEPTH=1 go test ./normal_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/normal_crawler	0.008s
$ DEPTH=10 go test ./normal_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/normal_crawler	0.408s
$ DEPTH=15 go test ./normal_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/normal_crawler	11.458s

$ DEPTH=1 go test ./concurrent_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/concurrent_crawler	0.016s
$ DEPTH=10 go test ./concurrent_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/concurrent_crawler	0.016s
$ DEPTH=15 go test ./concurrent_crawler --count 100
# ok  	github.com/ggary9424/go-tour-concurrency-10/concurrent_crawler	0.015s
```
