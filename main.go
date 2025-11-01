package main

import (
	"log"
	"sync"

	"github.com/guruorgoru/email-dispatcher/internal/consumers"
	"github.com/guruorgoru/email-dispatcher/internal/models"
	"github.com/guruorgoru/email-dispatcher/internal/producers"
	"golang.org/x/time/rate"
)

func main() {
	ch := make(chan models.Reciever, 50)
	var wg sync.WaitGroup
	go func() {
		if err := producers.LoadRecievers("./emails.csv", ch); err != nil {
			log.Fatalln("Error loading csv file:", err)
		}
	}()

	workerCount := 200
	limiter := rate.NewLimiter(rate.Limit(1), 200)
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go func() {
			consumers.WorkWithEmail(i, ch, &wg, limiter)
		}()
	}

	wg.Wait()
}
