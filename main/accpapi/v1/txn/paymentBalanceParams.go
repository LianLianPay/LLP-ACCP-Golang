package txn

/**
 * 余额支付 请求参数
 */
type PaymentBalanceParams struct {
	Timestamp          string           `json:"timestamp"`
	OidPartner         string           `json:"oid_partner"`
	TxnSeqNo           string           `json:"txn_seqno"`
	TotalAmount        float64          `json:"total_amount"`
	CouponAmount       float64          `json:"coupon_amount"`
	DirectionalPayFlag string           `json:"directionalpay_flag"`
	RiskItem           string           `json:"risk_item"`
	PayerInfo          PaymentPayerInfo `json:"payerInfo"`
}
