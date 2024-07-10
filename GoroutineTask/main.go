package main

import (
	"fmt"
	"sync"
)

func parallelExecute(tasks []func() error, N int, maxErrors int) error {
	var wg sync.WaitGroup
	taskChan := make(chan func() error, len(tasks))
	errorChan := make(chan error, len(tasks))
	quitChan := make(chan struct{})
	var errorsCount int
	var errorsCountMutex sync.Mutex

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case task, ok := <-taskChan:
					if !ok {
						return
					}
					if err := task(); err != nil {
						errorChan <- err
					} else {
						errorChan <- nil
					}
				case <-quitChan:
					return
				}
			}
		}()
	}
	go func() {
		for _, task := range tasks {
			select {
			case taskChan <- task:
			case <-quitChan:
				return
			}
		}
		close(taskChan)
	}()

	go func() {
		for err := range errorChan {
			if err != nil {
				errorsCountMutex.Lock()
				errorsCount++
				errorsCountMutex.Unlock()
			}
			if errorsCount >= maxErrors {
				close(quitChan)
				return
			}
		}
	}()

	wg.Wait()
	close(errorChan)

	if errorsCount >= maxErrors {
		return fmt.Errorf("processing stopped after reaching the max numbers of errors: %d", errorsCount)
	}
	return nil
}

func main() {
	tasks := []func() error{
		func() error {
			fmt.Println("Task 1 executed")
			return nil
		},
		func() error {
			fmt.Println("Task 2 executed")
			return nil
		},
		func() error {
			fmt.Println("Task 3 executed")
			return nil
		},
		func() error {
			fmt.Println("Task 4 executed")
			return nil
		},
		func() error {
			fmt.Println("Task 5 executed")
			return nil
		},
	}

	maxParallelsTasks := 2
	maxErrors := 2

	err := parallelExecute(tasks, maxParallelsTasks, maxErrors)
	if err != nil {
		fmt.Println("Execution stopped: ", err)
	} else {
		fmt.Println("All tasks executed successfully")
	}
}
