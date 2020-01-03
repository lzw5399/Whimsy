package main

import "fmt"

func main() {
	var emailService IEmailService
	emailService = &EmailService{}
	status := emailService.Send("这是内容")
	fmt.Println(status)
}

type IEmailService interface {
	Send(content string) bool
}

type EmailService struct {
}

func (e EmailService) Send(content string) (status bool) {
	fmt.Println(content)
	status = true
	return
}
