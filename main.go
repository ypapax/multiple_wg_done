package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		log.Println("waited 1")
	}()
	go func() {
		wg.Wait()
		log.Println("waited 2")
	}()
	wg.Wait()
	log.Println("waited 3")
	time.Sleep(3 * time.Second)
}
