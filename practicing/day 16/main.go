package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) (string, error)
}

type Payment struct {
	Amount        float64
	PaymentMethod PaymentMethod
}

func (p Payment) Process() {
	msg, err := p.PaymentMethod.Pay(p.Amount)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(msg)

}

type CreditCard struct {
	Number string
}

func (cc CreditCard) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Paid %.2f using Credit Card", amount), nil
}

type Pix struct {
	Key string
}

func (p Pix) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Paid %.2f using Pix", amount), nil
}

type Boleto struct {
	Code string
}

func (b Boleto) Pay(amount float64) (string, error) {
	return fmt.Sprintf("Paid %.2f using Boleto", amount), nil
}

func main() {
	fmt.Println("Today's review:")
	fmt.Println("Understanding how Go uses interfaces in practical way (without useless theory)")

	payments := make([]Payment, 0, 3)
	payments = append(payments, Payment{Amount: 100, PaymentMethod: Pix{Key: "6165156156"}})
	payments = append(payments, Payment{Amount: 50, PaymentMethod: Boleto{}})
	payments = append(payments, Payment{Amount: 75, PaymentMethod: CreditCard{Number: "5456-5454-3243-2345"}})

	for _, item := range payments {
		item.Process()
	}

}
