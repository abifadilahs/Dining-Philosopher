/* Solution 4: Monitor dengan Status (Gaya Tanenbaum) */
package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 5

const (
	Thinking = iota
	Hungry
	Eating
)

type Table struct {
	mu    sync.Mutex
	cond  *sync.Cond
	state [N]int
}

func NewTable() *Table {
	t := &Table{}
	t.cond = sync.NewCond(&t.mu)
	return t
}

func left(i int) int  { return (i + N - 1) % N }
func right(i int) int { return (i + 1) % N }

func (t *Table) TakeForks(i int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.state[i] = Hungry
	// Tunggu selama ada tetangga yang sedang makan
	for t.state[left(i)] == Eating || t.state[right(i)] == Eating {
		t.cond.Wait() // tidur & lepas mutex otomatis
	}
	t.state[i] = Eating
}

func (t *Table) PutForks(i int) {
	t.mu.Lock()
	t.state[i] = Thinking
	t.cond.Broadcast() // bangunkan yang menunggu, mereka cek ulang kondisinya
	t.mu.Unlock()
}

func philosopher(id int, t *Table, wg *sync.WaitGroup) {
	defer wg.Done()
	for round := 1; round <= 3; round++ {
		fmt.Printf("Filsuf %d berpikir\n", id)
		time.Sleep(10 * time.Millisecond)

		t.TakeForks(id)
		fmt.Printf("Filsuf %d MAKAN (ronde %d)\n", id, round)
		time.Sleep(50 * time.Millisecond)
		t.PutForks(id)
	}
}

func main() {
	table := NewTable()
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go philosopher(i, table, &wg)
	}
	wg.Wait()
	fmt.Println("Selesai tanpa deadlock")
}
