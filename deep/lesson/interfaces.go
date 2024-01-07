package main

import (
	"fmt"
)

type Caller interface {
	Call(num int) error
}

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
	fmt.Printf("На адрес %v было направлено письмо", msg)
	return nil
}

func (p *Phone) Call(num int) error {
	fmt.Printf("На номер %v был совершен звонок, баланс составил %v рублей", num, p.Balance)
	return nil
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("На номер %v был совершен звонок, баланс составил %v рублей", msg, p.Balance)
	return nil
}

func Notify(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Int не поддерживается")
	case *Email:
		mail, ok := i.(*Email)
		if ok {
			fmt.Println("Отправлен на электронную почту -", mail.Address)
		}
		fmt.Println("Success!")
	case *Phone:
		phone, ok := i.(*Phone)
		if ok {
			fmt.Printf("Отправлен по телефону %v\n", phone.Number)
			fmt.Printf("Баланс составляет %v$", phone.Balance)
		}
		fmt.Println("Success!")
	}
}

func main() {
	e := &Email{Address: "dodexpwnz@gmail.com"}
	Notify(e)

	p := &Phone{Number: 95210706, Balance: 120}
	Notify(p)
}
