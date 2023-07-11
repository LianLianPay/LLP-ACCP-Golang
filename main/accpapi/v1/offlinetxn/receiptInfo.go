package offlinetxn

type ReceiptInfo struct {
	ReceiptFilename string  `json:"receipt_filename"`
	Amount          float64 `json:"amount"`
	/*
	   付款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	PayerType string `json:"payer_type"`
	PayerID   string `json:"payer_id"`
	Method    string `json:"method"`
	PayeeType string `json:"payee_type"`
	PayeeID   string `json:"payee_id"`
}
