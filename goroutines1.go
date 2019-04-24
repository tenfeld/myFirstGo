package main

import (
	"fmt"
	"time"
	"sync"
)

func say (s string) {
	for i:=0; i<3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum (s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// c チャネルに sum を送信
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i <n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	{
		// go fn() と書けば、新しい goroutine が実行される
		go say ("1-1. Hello")
		say("1-2. World")
	}

	{
		// 共有メモリへのアクセスは、チャネル(Channel)を用いて行う
		s := []int {7, 2, 8, -9, 4, 0}

		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		// c チャネルからデータを受け取り
		x, y := <-c, <-c

		fmt.Println("2. ", x, y, x+y)
	}

	{
		// チャネルはバッファバッファとして使える。
		// バッファをもつチャネルを初期化するには、make の2つ目の引数にバッファの長さを与える
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		//ch <- 3
		fmt.Println("3. ", <-ch)
		fmt.Println("3. ", <-ch)
		//fmt.Println(<-ch)
	}

	{
		// データの送り手は、これ以上送信する値がないことを示すため、チャネルをcloseできる
		// 受け取り側は、受信の式に2つ目のパラメータを割り当てることで、そのチャネルがcloseされているか確認できる
		// v, ok := <- ch
		// 受けて側はチャネルをcloseしてはいけない
		c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c {
			fmt.Println("4. ", i)
		}
	}

	{
		// select ステートメントは、goroutineを複数の通信操作で待たせる
		fn := func(c, quit chan int) {
			x, y:= 0, 1
			for {
				select {
				case c <- x:
					fmt.Println("5. ", "send to channel", x)
					x, y = y, x+y
				case <- quit:
					fmt.Println("5. ", "quit")
					return
				default:
					time.Sleep(100)
				}
			}
		}

		c:= make(chan int)
		quit := make(chan int)
		go func() {
			for i:= 0; i <10; i++ {
				v := <-c
				fmt.Println("5. ", "receive from channel", v)
			}
			quit <- 0
		}()
		fn(c, quit)
	}

	{
		// 排他制御 (mutex)
		// Lock() と Unlock() で囲むことで、排他制御で実行する
		c:= SafeCounter{v: make(map[string]int)}
		for i :=0; i <1000; i++ {
			go c.Inc("somekey")
		}
		time.Sleep(time.Second)
		fmt.Println("6. ", c.Value("somekey"))
	}
}