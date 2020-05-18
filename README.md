# SlidingWindow

A simple implementation of a sliding window structure. It's basically a
circular buffer that holds the last `n` elements that you add to it.

## Example usage
```go
sw, err := slidingwindow.New(3)
if err != nil {
    log.Fatal(err)
}

sw.Add(1)
sw.Add(2)
sw.Add(3)
sw.Add(4)
sw.Add(5)

fmt.Println(sw)
```

```
$ go run main.go
[3 4 5]
```
