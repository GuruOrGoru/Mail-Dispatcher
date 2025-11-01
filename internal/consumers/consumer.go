package consumers

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"sync"
	"time"

	"github.com/guruorgoru/email-dispatcher/internal/models"
	"golang.org/x/time/rate"
)

func WorkWithEmail(id int, ch chan models.Reciever, wg *sync.WaitGroup, limiter *rate.Limiter) {
	defer wg.Done()
	smtpHost := "localhost"
	smtpPort := "1025"

	for reciever := range ch {
		if err := limiter.Wait(context.TODO()); err != nil {
			log.Println("Rate limiter error:", err)
			continue
		}
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		message, err := execTemp(reciever)
		if err != nil {
			log.Printf("❌ Failed to execute the template: %v\n", err)
			continue
		}
		fmt.Printf("Worker %d processing email for %+v\n", id, reciever)
		if err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "guruorgoru1@gmail.com", []string{reciever.Email}, []byte(message)); err != nil {
			log.Printf("❌ Worker %d failed to send to %s: %v\n", id, reciever.Email, err)
			continue
		}

		fmt.Printf("Worker %d processed email for %+v\n", id, reciever)
	}
}

func execTemp(r models.Reciever) (string, error) {
	temp, err := template.ParseFiles("email.templ")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = temp.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
