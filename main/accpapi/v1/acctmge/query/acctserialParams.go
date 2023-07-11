package query

/**
 * 资金流水详情列表查询 请求参数
 */
type AcctserialParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	UserType   string `json:"user_type"`
	AcctType   string `json:"acct_type"`
	DateStart  string `json:"date_start"`
	DateEnd    string `json:"date_end"`
	FlagDC     string `json:"flag_dc"`
	PageNo     string `json:"page_no"`
	PageSize   string `json:"page_size"`
	SortType   string `json:"sort_type"`
}
