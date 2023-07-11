package txn

type PaymentPayerInfo struct {
	PayerType string `json:"payer_type"`
	PayerID   string `json:"payer_id"`
	UserID    string `json:"user_id"`
	Password  string `json:"password"`
	RandomKey string `json:"random_key"`
}
