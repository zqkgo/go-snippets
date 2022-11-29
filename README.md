- [打印原始HTTP响应 dump raw HTTP response message](#打印原始http响应-dump-raw-http-response-message)
- [计算N的整数倍 rounding to N](#计算n的整数倍-rounding-to-n)
- [下一秒时间戳 next second](#下一秒时间戳-next-second)
- [浮点数保留精度 round to precision](#浮点数保留精度-round-to-precision)
- [打印x进制对应的十进制 print decimal of binary/octal/hexadecimal](#打印x进制对应的十进制-print-decimal-of-binaryoctalhexadecimal)

## 打印原始HTTP响应 dump raw HTTP response message

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

## 计算N的整数倍 rounding to N

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

## 下一秒时间戳 next second

```go
func main() {
	now := time.Now()
	round := now.Truncate(1 * time.Second)
	next := round.Add(1 * time.Second)
	fmt.Printf("now: %v\nround: %v\nnext: %v\n", now, round, next)
}
// output:
// now: 2022-11-23 11:01:29.566766 +0800 CST m=+0.000057801
// round: 2022-11-23 11:01:29 +0800 CST
// next: 2022-11-23 11:01:30 +0800 CST
```

## 浮点数保留精度 round to precision

```go
func main() {
	fmt.Println(roundToPrecision(3.1415926, 2)) // 3.14
	fmt.Println(roundToPrecision(3.1415926, 3)) // 3.142
	fmt.Println(roundToPrecision(3.1, 1))       // 3.1
}

func roundToPrecision(v float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(v*ratio) / ratio
}
```

## 打印x进制对应的十进制 print decimal of binary/octal/hexadecimal

```go
func main() {
	fmt.Println(0b11111111) // 255 二进制
	fmt.Println(0111)       // 73 八进制
	fmt.Println(0xa5c)      // 2652 十六进制
}
```