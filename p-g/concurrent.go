package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 1. 基础并发机制
func basicConcurrency() {
	var wg sync.WaitGroup

	// Goroutine基本使用
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 executed")
	}()

	// 带参数的Goroutine
	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("Goroutine 2 with message")

	wg.Wait()
}

// 2. 通道(Channel)使用
func channelDemo() {
	// 无缓冲通道
	ch := make(chan int)
	go func() {
		ch <- 42 // 发送数据
	}()
	fmt.Println("Received:", <-ch) // 接收数据

	// 缓冲通道
	bufCh := make(chan int, 2)
	bufCh <- 1
	bufCh <- 2
	fmt.Println("Buffered 1:", <-bufCh)
	fmt.Println("Buffered 2:", <-bufCh)

	// 关闭通道检查
	close(bufCh)
	if _, ok := <-bufCh; !ok {
		fmt.Println("Channel closed")
	}
}

// 3. 工作池模式
func workerPool() {
	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// 创建工作池
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := range jobs {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				results <- j * 2
				fmt.Printf("Worker %d processed job %d\n", workerID, j)
			}
		}(i)
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

	// 输出结果
	for res := range results {
		fmt.Println("Result:", res)
	}
}

// 4. 竞态条件解决方案
func raceSolution() {
	var (
		counter int32
		wg      sync.WaitGroup
	)

	// 原子操作解决竞态
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	// Mutex解决竞态
	var (
		mu      sync.Mutex
		counter2 int
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter2++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic counter: %d\nMutex counter: %d\n", atomic.LoadInt32(&counter), counter2)
}

// 5. 高级并发模式
func advancedPatterns() {
	// Select多路复用
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-time.After(1500 * time.Millisecond):
			fmt.Println("timeout")
		}
	}

	// Once用法
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	for i := 0; i < 5; i++ {
		once.Do(onceBody)
	}

	// 条件变量
	var (
		mu      sync.Mutex
		cond    = sync.NewCond(&mu)
		isReady bool
	)
	go func() {
		time.Sleep(1 * time.Second)
		mu.Lock()
		isReady = true
		cond.Broadcast()
		mu.Unlock()
	}()

	mu.Lock()
	for !isReady {
		cond.Wait()
	}
	mu.Unlock()
	fmt.Println("Condition met")
}

// 6. 并发安全设计
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func concurrencySafety() {
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("key")
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", c.Value("key"))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置使用所有CPU核心

	fmt.Println("=== Basic Concurrency ===")
	basicConcurrency()

	fmt.Println("\n=== Channel Demo ===")
	channelDemo()

	fmt.Println("\n=== Worker Pool ===")
	workerPool()

	fmt.Println("\n=== Race Solution ===")
	raceSolution()

	fmt.Println("\n=== Advanced Patterns ===")
	advancedPatterns()

	fmt.Println("\n=== Concurrency Safety ===")
	concurrencySafety()
}