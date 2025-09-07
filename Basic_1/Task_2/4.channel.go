package Task_2

import (
	"fmt"
	"sync"
)

// Generator 生成整数
func Generator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// Receiver 接收并打印整数
func Receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
}

// HundredIntegersGenerator 生成 100 个整数
func HundredIntegersGenerator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
