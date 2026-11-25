package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	oddTurn := make(chan struct{})
	evenTrun := make(chan struct{})
	var wg sync.WaitGroup
	go odd(oddTurn, evenTrun, &wg)
	go even(oddTurn, evenTrun, &wg)
	oddTurn <- struct{}{}
	wg.Wait()
}

func odd(oddTurn chan struct{}, evenTurn chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(oddTurn)
		evenTurn <- struct{}{}
	}()
	wg.Add(1)
	count := 1
	for range oddTurn {
		fmt.Println(count)
		count += 2
		time.Sleep(1 * time.Second)
		if count > 10 {
			return
		} else {
			evenTurn <- struct{}{}
		}
	}
}

func even(oddTurn chan struct{}, evenTurn chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(evenTurn)
	}()
	wg.Add(1)
	count := 2
	for range evenTurn {
		fmt.Println(count)
		count += 2
		time.Sleep(1 * time.Second)
		if count > 10 {
			return
		} else {
			oddTurn <- struct{}{}
		}
	}
}
