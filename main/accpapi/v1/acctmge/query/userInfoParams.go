package query

/**
 * 用户信息查询 请求参数
 */
type UserInfoParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
}
