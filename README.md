- [dump raw HTTP response message](#dump-raw-http-response-message)
- [rounding to N](#rounding-to-n)
- [next second](#next-second)

## dump raw HTTP response message

```go
func main() {
	resp, err := http.Get("https://www.github.com")
	if err != nil {
		panic(err)
	}
	// without body
	nobody, err := httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}
	// with body
	withbody, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp line and header len: %d, body len: %d\n", len(nobody), len(withbody))
}
```

## rounding to N

```go
func main() {
	// rounding to n
	n := 32
	fmt.Println(31 &^ (n - 1)) // 0
	fmt.Println(32 &^ (n - 1)) // 32
	fmt.Println(55 &^ (n - 1)) // 32
	fmt.Println(70 &^ (n - 1)) // 64
}
```

## next second

```go
func main() {
	now := time.Now()
	round := now.Truncate(1 * time.Second)
	next := round.Add(1 * time.Second)
	fmt.Printf("now: %v\nround: %v\nnext: %v\n", now, round, next)
}
```