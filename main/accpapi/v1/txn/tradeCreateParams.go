package txn

/**
 * 支付统一创单 请求参数
 */
type TradeCreateParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnType    string `json:"txn_type"`
	UserID     string `json:"user_id"`
	UserType   string `json:"user_type"`
	NotifyURL  string `json:"notify_url"`
	ReturnURL  string `json:"return_url"`
	PayExpire  string `json:"pay_expire"`

	OrderInfo TradeCreateOrderInfo
	PayeeInfo []TradeCreatePayeeInfo
}
