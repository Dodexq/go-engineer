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

type Phone struct {
	Number  int
	Balance float64
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
	}

	s, ok := i.(Sender)
	if !ok {
		fmt.Println("Ошибка пустого интерфейса")
		return
	}
	err := s.Send("Сообщение пустого интерфейса")
	if err != nil {
		fmt.Println("Ошибка")
	}
	fmt.Println("Sucess")

}

func main() {
	Notify(&Phone{Number: 23})

}
