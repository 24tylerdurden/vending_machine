package interfaces

type UserActions interface {
	SelectItem(item string) error
	InsertCoin(amount int) error
	DispenseItems() error
	CancelTransaction() error
}

type PaymentStrategy interface {
	Pay(amount int) (bool, error)
	AddMoney(amount int)
	GetBalance() int
}
