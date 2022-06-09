package main

import (
	"fmt"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// これだと最後まで"＊"が表示されないことがある。 -> syncを使用してforを回し切って残りのmain()の処理を行う様にする必要がある
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// func main() {
// 	fmt.Println(time.Now())
// 	fmt.Println()
// 	count := 0
// 	for i := 0; i < 1000; i ++ {
// 		go func() {
// 			count ++
// 			fmt.Print("*")
// 		}()
// 	}

// 	time.Sleep(1 * time.Second)
// 	fmt.Println()
// 	fmt.Printf("\n%d\n", count)
// 	fmt.Println(time.Now())
// }

func main() {
	fmt.Println(time.Now())
	fmt.Println()
	count := 0
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			count++
			fmt.Print("*")
			// mu.Unlock()
			wg.Done()
		}()
	}

	// time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println()
	fmt.Printf("\n%d\n", count)
	fmt.Println(time.Now())
}
