package payment

import "errors"

type CardPayment struct {
	balance int
}

func NewCardPayment(balance int) *CardPayment {
	return &CardPayment{
		balance: balance,
	}
}

func (c *CardPayment) Pay(amount int) (bool, error) {
	if amount > c.balance {
		return false, errors.New("Insufficent Balance")
	}

	return true, nil
}

func (c *CardPayment) AddMoney(amount int) {

	// Handle this case as well
	if amount < 0 {

	}

	c.balance += amount
}

func (c *CardPayment) GetBalance() int {
	return c.balance
}
