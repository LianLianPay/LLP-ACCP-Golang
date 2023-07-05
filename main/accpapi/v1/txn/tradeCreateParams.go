package txn

type TradeCreateParams struct {
	Timestamp  string
	OidPartner string
	TxnType    string
	UserID     string
	UserType   string
	NotifyURL  string
	ReturnURL  string
	PayExpire  string

	OrderInfo TradeCreateOrderInfo
	PayeeInfo []TradeCreatePayeeInfo
}
