package query

/**
 * 绑卡信息查询 请求参数
 */
type QueryLinkedAcctParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
}
