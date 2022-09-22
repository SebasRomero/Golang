package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {
	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	fakeIDsChan := make(chan string)
	closedChans := make(chan int)

	wg.Add(3)
	go generateIDS(IDsChan, wg, closedChans)
	go generateFakeIDs(fakeIDsChan, wg, closedChans)
	go logIDs(IDsChan, fakeIDsChan, wg, closedChans)

	wg.Wait()
}

func generateFakeIDs(fakeIDsChan chan<- string, wg *sync.WaitGroup, closedChan chan<- int) {
	for i := 0; i < 50; i++ {
		id := uuid.New()
		fakeIDsChan <- fmt.Sprintf("%d. %s", i+1, id.String())

	}
	close(fakeIDsChan)
	closedChan <- 1

	wg.Done()
}

func generateIDS(idsChan chan<- string, wg *sync.WaitGroup, closedChan chan<- int) {
	for i := 0; i < 100; i++ {
		id := uuid.New()
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())

	}
	close(idsChan)
	closedChan <- 1

	wg.Done()
}

func logIDs(idsChan <-chan string, fakeIDsChan <-chan string, wg *sync.WaitGroup, closedChan chan int) {

	closedCounter := 0
	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("ID: ", id)
			}
		case id, ok := <-fakeIDsChan:
			if ok {
				fmt.Println("Fake ID: ", id)
			}
		case count, ok := <-closedChan:
			if ok {
				closedCounter += count
			}
		}
		if closedCounter == 2 {
			close(closedChan)
			break
		}
	}
	wg.Done()
}
