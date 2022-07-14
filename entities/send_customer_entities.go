package entities

type SendCustomer struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (*SendCustomer) TableName() string {
	return "SendCustomer"
}
