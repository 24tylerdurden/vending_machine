package payment

import "errors"

type CoinPayment struct {
	balance int
}

func NewCoinPayment() *CoinPayment {
	return &CoinPayment{}
}

func (c *CoinPayment) Pay(amount int) (bool, error) {
	if amount > c.balance {
		return false, errors.New("Insufficient fund")
	}

	c.balance -= amount

	return true, nil
}

func (c *CoinPayment) AddMoney(amount int) {
	c.balance += amount
}

func (c *CoinPayment) GetBalance() int {
	return c.balance
}
