package refund

/**
 * 退款申请 请求参数
 */
type MorePayeeRefundParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	NotifyURL  string `json:"notify_url"`
	// 退款原因
	RefundReason string `json:"refund_reason"`
	// 原商户订单信息
	OriginalOrderInfo OriginalOrderInfo `json:"originalOrderInfo"`
	// 退款订单信息
	RefundOrderInfo RefundOrderInfo `json:"refundOrderInfo"`
	// 原收款方退款信息
	PyeeRefundInfos []PyeeRefundInfo `json:"pyeeRefundInfos"`
	// 原付款方式
	RefundMethods []RefundMethod `json:"refundMethods"`
}
