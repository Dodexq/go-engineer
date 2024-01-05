package main

import (
	"fmt"
)

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

type Phone struct {
	Number  int
	Balance float64
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по почте на адрес %v \n", msg, e.Address)
	return nil
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по номеру телефона %v \n", msg, p.Number)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch s.(type) {
	case *Email:
		fmt.Println("Сообщение по почте")
	case *Phone:
		fmt.Println("Сообщение по телефону")

	}
	phone, ok := s.(*Phone)
	if ok {
		fmt.Println("Balance is -", phone.Balance)
	}
	fmt.Println("Success!")
}

func main() {
	a := &Email{Address: "dodexpwnz@gmail.com"}
	Notify(a)

	b := &Phone{Number: 95210706, Balance: 102.5}
	Notify(b)
}
