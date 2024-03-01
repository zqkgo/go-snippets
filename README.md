- [打印原始HTTP响应 dump raw HTTP response message](#打印原始http响应-dump-raw-http-response-message)
- [计算N的整数倍 rounding to N](#计算n的整数倍-rounding-to-n)
- [下一秒时间戳 next second](#下一秒时间戳-next-second)
- [浮点数保留精度 round to precision](#浮点数保留精度-round-to-precision)
- [打印x进制对应的十进制 print decimal of binary/octal/hexadecimal](#打印x进制对应的十进制-print-decimal-of-binaryoctalhexadecimal)
- [昨天的日期 print the date of yesterday](#昨天的日期-print-the-date-of-yesterday)
- [接口、实现与配置 interface, implementaton and options](#接口实现与配置-interface-implementation-and-options)
- [避免参数零值 prevent zero value in args](#避免参数零值-prevent-zero-value-in-args)
- [浅复制结构体 shallow copy struct](#浅复制结构体-shallow-copy-struct)
- [重定向标准错误到文件 redirect stderr to a file](#重定向标准错误到文件-redirect-stderr-to-a-file)
- [覆盖结构体的同时保留旧地址 override a struct while keeping the old address](#覆盖结构体的同时保留旧地址-override-a-struct-while-keeping-the-old-address)
- [清空slice但保留已分配内存 clear a slice but keep the allocated memory](#清空slice但保留已分配内存-clear-a-slice-but-keep-the-allocated-memory)
- [并发监听多个 channel listening on multiple channels concurrently](#并发监听多个channel-listening-on-multiple-channels-concurrently)
- [执行系统命令 execute system command](#执行系统命令-execute-system-command)
- [字符串切片编解码成字符串 convert between string and string slice](#字符串切片编解码成字符串-convert-between-string-and-string-slice)

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

## 昨天的日期 print the date of yesterday

```go
func main() {
	t, err := time.Parse("2006-01-02 15:04:05", "2022-11-01 13:00:00")
	if err != nil {
		panic(err)
	}
	yesterday := t.AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println(yesterday) // 2022-10-31

	t, err = time.Parse("2006-01-02 15:04:05", "2022-11-02 13:00:00")
	if err != nil {
		panic(err)
	}
	yesterday = t.AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println(yesterday) // 2022-11-01
}
```

## 接口、实现与配置 interface, implementation and options

```go
// options used by all implementations
type walkerOptions struct {
	dur time.Duration
}

type Walker interface {
	Walk()
}

type dog struct {
	opts walkerOptions
}

func newDog(opts walkerOptions) *dog {
	return &dog{
		opts: opts,
	}
}

func (d *dog) Walk() {
	for {
		println("dog is walking")
		time.Sleep(d.opts.dur)
	}
}

type human struct {
	opts walkerOptions
}

func newHuman(opts walkerOptions) *human {
	return &human{
		opts: opts,
	}
}

func (h *human) Walk() {
	for {
		println("human is walking")
		time.Sleep(h.opts.dur)
	}
}

func main() {
	var walkers []Walker
	walkers = append(walkers, newDog(walkerOptions{
		dur: 500 * time.Millisecond,
	}))
	walkers = append(walkers, newHuman(walkerOptions{
		dur: 1 * time.Second,
	}))
	for _, w := range walkers {
		go w.Walk()
	}
	select {}
}
```

## 避免参数零值 prevent zero value in args

```go
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostTasksArgs struct {
	Seq    *int   `json:"seq"`
	Owner  string `json:"owner"`
	Enable *bool  `json:"enable"`
	Task   *Task  `json:"task"`
}

func main() {
	s := `{"owner":"root", "enable":true}`
	var args PostTasksArgs
	err := json.Unmarshal([]byte(s), &args)
	if err != nil {
		panic(err)
	}
	if args.Enable != nil {
		println("enable: ", *args.Enable) // true
	}
	fmt.Printf("%+v\n", args) // {Seq:<nil> Owner:root Enable:0xc00001428c Task:<nil>}
}

```

## 浅复制结构体 shallow copy struct

```go
type Deployment struct {
	ID      string
	Status  int
	StartAt time.Time
}

func main() {
	d1 := &Deployment{
		ID:      "1",
		Status:  2,
		StartAt: time.Now().Add(1 * time.Hour),
	}
	d2 := *d1
	// d1 addr: 0xc00007c180, d2 addr: 0xc00007c1b0, d1 val == d2 val: true
	fmt.Printf("d1 addr: %p, d2 addr: %p, d1 val == d2 val: %v\n", d1, &d2, *d1 == d2)
}
```

## 重定向标准错误到文件 redirect stderr to a file

```go
func main() {
	name := "tmp.txt"
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	// redirect stderr to fd
	err = syscall.Dup2(int(f.Fd()), 2)
	if err != nil {
		panic(err)
	}
	f.Close()

	msg := "message written to stderr"
	os.Stderr.WriteString(msg)

	f, err = os.Open(name)
	if err != nil {
		panic(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(msg == string(bs)) // should be true

	os.Remove(name)
} 
```

## 覆盖结构体的同时保留旧地址 override a struct while keeping the old address

```go
type Foo struct {
	Bar string
	Biz int
}

func main() {
	p1 := &Foo{Bar: "Hello", Biz: 1024}
	p2 := &Foo{Bar: "World", Biz: 2048}
	*p2 = *p1
	fmt.Println(p2) // &{Hello 1024}
}
```

## 清空slice但保留已分配内存 clear a slice but keep the allocated memory

```go
func main() {
	var nums []int
	for i := 0; i < 2048; i++ {
		nums = append(nums, i)
	}
	cap1 := cap(nums)
	nums = nums[:0]
	cap2 := cap(nums)
	fmt.Println(len(nums), cap1 == cap2, nums[99:103]) // 0 true [99 100 101 102]
}
```

## 并发监听多个 channel listening on multiple channels concurrently

```go
func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func() {
		for {
			var s string
			select {
			case s = <-ch1:
			case s = <-ch2:
			}
			// to ch1
			// to ch2
			// to ch1
			// to ch2
			fmt.Println(s)
		}
	}()
	var n int
	for {
		time.Sleep(1 * time.Second)
		if n%2 == 0 {
			ch1 <- "to ch1"
		} else {
			ch2 <- "to ch2"
		}
		n++
	}
}
```

## 执行系统命令 execute system command

```go
func main() {
	bs, err := execCmd("ps -eo pid,ppid,command")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func execCmd(line string) ([]byte, error) {
	return exec.Command("bash", "-c", line).Output()
}
```

## 字符串切片编解码成字符串 convert between string and string slice
```go
type whateverKey string

func (k *whateverKey) enc(domains []string) {
	sort.SliceStable(domains, func(i, j int) bool {
		return domains[i] < domains[j]
	})
	*k = whateverKey(strings.Join(domains, ","))
}

func (k *whateverKey) dec() []string {
	return strings.Split(string(*k), ",")
}

func main() {
	var k whateverKey
	k.enc([]string{"c.com", "d.com", "a.com", "b.com", "e.com"})

	// a.com,b.com,c.com,d.com,e.com
	fmt.Println(k)
	// [a.com b.com c.com d.com e.com]
	fmt.Println(k.dec())
}
```