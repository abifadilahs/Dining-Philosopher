package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 5

func philosopher(id int, forks []sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	left, right := id, (id+1)%N

	for round := 1; round <= 3; round++ {
		fmt.Printf("Filsuf %d berpikir\n", id)
		time.Sleep(10 * time.Millisecond)

		forks[left].Lock()
		fmt.Printf("Filsuf %d ambil garpu kiri (%d)\n", id, left)
		time.Sleep(100 * time.Millisecond) // memperbesar peluang deadlock

		forks[right].Lock()
		fmt.Printf("Filsuf %d ambil garpu kanan (%d) -> MAKAN\n", id, right)
		time.Sleep(50 * time.Millisecond)

		forks[right].Unlock()
		forks[left].Unlock()
	}
}

func main() {
	forks := make([]sync.Mutex, N)
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go philosopher(i, forks, &wg)
	}
	wg.Wait()
	fmt.Println("Selesai")
}
