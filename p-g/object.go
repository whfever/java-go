package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// ========== 面向对象 ==========
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rect struct {
	Width, Height float64
}

func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

// ========== 网络编程 ==========
func startHTTPServer(wg *sync.WaitGroup) {
	defer wg.Done()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go Server! Time: %s", time.Now().Format(time.RFC3339))
	})

	fmt.Println("HTTP server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func tcpClient(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("TCP connect error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello TCP Server"))
	if err != nil {
		fmt.Println("TCP write error:", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("TCP read error:", err)
		return
	}
	fmt.Println("TCP response:", string(buf[:n]))
}

// ========== 并发编程 ==========
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // 模拟工作耗时
		results <- j * 2
	}
}

func main() {
	// 面向对象示例
	c := Circle{Radius: 5}
	printArea(c)
	r := &Rect{Width: 3, Height: 4}
	printArea(r)

	// 并发控制
	var wg sync.WaitGroup

	// 启动HTTP服务器
	wg.Add(1)
	go startHTTPServer(&wg)

	// 启动TCP服务器
	wg.Add(1)
	go func() {
		defer wg.Done()
		ln, err := net.Listen("tcp", ":8081")
		if err != nil {
			panic(err)
		}
		defer ln.Close()

		fmt.Println("TCP server starting on :8081")
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println("TCP received:", string(buf[:n]))
		conn.Write([]byte("ACK: " + string(buf[:n])))
	}()

	// 等待服务器启动
	time.Sleep(1 * time.Second)

	// 启动TCP客户端
	wg.Add(1)
	go tcpClient(&wg)

	// 并发工作池示例
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// 启动3个worker
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()

	// 打印结果
	for res := range results {
		fmt.Println("Result:", res)
	}

	wg.Wait()
}