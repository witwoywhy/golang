package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func inti() {
	rand.Seed(time.Now().Unix())
}

const Timeout = 10
const FormatTime = "2006-01-02 15:04:05"

type Task struct {
	ID          int
	ProcessTime int
	Start       time.Time
}

func CreateTask(ch *chan Task) {
	id := 1
	for {
		n := rand.Intn(20 + 1)
		task := Task{
			ID:          id,
			ProcessTime: n,
		}

		*ch <- task
		id++

		log.Printf("Task ID: %-3d => Create Process Time: %-3d", task.ID, task.ProcessTime)

		if id > 10 {
			time.Sleep(time.Second * 1)
		}
	}
}

func Do(task *Task, result chan int) {
	log.Printf("Task ID: %-3d => Do", task.ID)

	task.Start = time.Now()
	time.Sleep(time.Second * time.Duration(task.ProcessTime))
	result <- rand.Intn(100 + 1)
}

func Worker(wg *sync.WaitGroup, ch *chan Task) {
	for {
		select {
		case task := <-*ch:
			log.Printf("Task ID: %-3d => Recive", task.ID)

			result := make(chan int, 1)
			go Do(&task, result)

			select {
			case r := <-result:
				log.Printf("Task ID: %-3d => Got Result: %-3d", task.ID, r)
			case <-time.After(time.Second * time.Duration(Timeout)):
				log.Printf("Task ID: %-3d => Time Out Start Task: %s", task.ID, task.Start.Format(FormatTime))
			}
		}
	}
}

func main() {
	ch := make(chan Task, 10)
	go CreateTask(&ch)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(&wg, &ch)
	}
	wg.Wait()
}