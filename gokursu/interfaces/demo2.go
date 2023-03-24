package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {

	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{creditPaymentTotal: 100000, address: "Adana", rate: 10}
	credit2 := Mortgage{creditPaymentTotal: 50000, address: "Mersin", rate: 12}
	credit3 := Car{creditPaymentTotal: 60000, carInfo: "polo", rate: 15}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println(total)
}
