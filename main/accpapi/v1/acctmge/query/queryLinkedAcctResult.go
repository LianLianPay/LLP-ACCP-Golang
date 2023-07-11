package query

/**
 * 绑卡信息查询 请求参数
 */
type QueryLinkedAcctResult struct {
	RetCode        string          `json:"ret_code"`
	RetMsg         string          `json:"ret_msg"`
	OidPartner     string          `json:"oid_partner"`
	UserID         string          `json:"user_id"`
	LinkedAcctList []LinkedAccount `json:"linked_acctlist"`
}
