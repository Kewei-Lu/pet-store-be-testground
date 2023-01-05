package model

type MoneyQueryInput struct {
	UserName string
}

type MoneyAddInput struct {
	UserName           string
	DestinationAccount string
	Amount             int
	Comment            string
}

type CreateMoneyAccountInput struct {
	UserName          string
	CreationTimeStamp string
}

type MoneyTransferInput struct {
	UserName           string
	DestinationAccount string
	Amount             int
	Comment            string
}
