/* Solution1: Batasi Maksimal 4 Filsuf di Meja */

package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 5

func philosopher(id int, forks []sync.Mutex, room chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	left, right := id, (id+1)%N

	for round := 1; round <= 3; round++ {
		fmt.Printf("Filsuf %d berpikir\n", id)
		time.Sleep(10 * time.Millisecond)

		room <- struct{}{} // masuk ruangan (blok jika sudah 4 orang)

		forks[left].Lock()
		time.Sleep(100 * time.Millisecond)
		forks[right].Lock()

		fmt.Printf("Filsuf %d MAKAN (ronde %d)\n", id, round)
		time.Sleep(50 * time.Millisecond)

		forks[right].Unlock()
		forks[left].Unlock()

		<-room // keluar ruangan
	}
}

func main() {
	forks := make([]sync.Mutex, N)
	room := make(chan struct{}, N-1) // semaphore dengan 4 izin
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go philosopher(i, forks, room, &wg)
	}
	wg.Wait()
	fmt.Println("Selesai tanpa deadlock")
}
