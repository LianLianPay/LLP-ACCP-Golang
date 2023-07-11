package cashier

/**
 * 账户+收银台 请求参数
 */
type CashierPayCreateParams struct {
	Timestamp  string                      `json:"timestamp"`
	OidPartner string                      `json:"oid_partner"`
	TxnType    string                      `json:"txn_type"`
	UserID     string                      `json:"user_id"`
	UserType   string                      `json:"user_type"`
	NotifyURL  string                      `json:"notify_url"`
	ReturnURL  string                      `json:"return_url"`
	PayExpire  string                      `json:"pay_expire"`
	FlagChnl   string                      `json:"flag_chnl"`
	RiskItem   string                      `json:"risk_item"`
	Extend     string                      `json:"extend"`
	OrderInfo  CashierPayCreateOrderInfo   `json:"orderInfo"`
	PayeeInfo  []CashierPayCreatePayeeInfo `json:"payeeInfo"`
	Style      CashierPayCreateStyle       `json:"style"`
	PayerInfo  CashierPayCreatePayerInfo   `json:"payerInfo"`
}
