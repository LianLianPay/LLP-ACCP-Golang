package txn

type WithDrawalPayerInfo struct {
	PayerType     string `json:"payer_type"`
	PayerID       string `json:"payer_id"`
	PayerAcctType string `json:"payer_accttype"`
	Password      string `json:"password"`
	RandomKey     string `json:"random_key"`
	PapAgreeNo    string `json:"pap_agree_no"`
}
