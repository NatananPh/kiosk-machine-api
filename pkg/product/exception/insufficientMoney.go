package exception

type InsufficientMoney struct{}

func (i *InsufficientMoney) Error() string {
	return "Insufficient money"
}