package main

import (
	"fmt"

	fake "github.com/brianvoe/gofakeit/v6"

	"quocbang/go-queue/queue"
)

func main() {
	q := queue.BuildQueue()

	for i := 0; i < 5; i++ {
		q.Push(fake.Company())
	}

	for i := 0; i < 5; i++ {
		message := q.Pop()
		fmt.Println(message)
	}
}
