package secured

type SecuredConfirmOrderInfo struct {
	ConfirmSeqNo  string  `json:"confirm_seqno"`
	ConfirmTime   string  `json:"confirm_time"`
	ConfirmAmount float64 `json:"confirm_amount"`
	CouponAmount  float64 `json:"coupon_amount"`
}
