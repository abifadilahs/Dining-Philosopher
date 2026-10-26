/* Solution 2: Hierarki Sumber Daya (Lock Ordering) */
package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 5

func philosopher(id int, forks []sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	a, b := id, (id+1)%N

	// Urutkan: garpu bernomor kecil diambil lebih dulu
	first, second := a, b
	if first > second {
		first, second = second, first
	}

	for round := 1; round <= 3; round++ {
		fmt.Printf("Filsuf %d berpikir\n", id)
		time.Sleep(10 * time.Millisecond)

		forks[first].Lock()
		time.Sleep(100 * time.Millisecond)
		forks[second].Lock()

		fmt.Printf("Filsuf %d MAKAN (ronde %d) dengan garpu %d dan %d\n", id, round, first, second)
		time.Sleep(50 * time.Millisecond)

		forks[second].Unlock()
		forks[first].Unlock()
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
	fmt.Println("Selesai tanpa deadlock")
}
