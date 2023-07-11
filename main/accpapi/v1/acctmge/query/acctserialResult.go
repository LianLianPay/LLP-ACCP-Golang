package query

/**
 * 资金流水详情列表查询 响应参数
 */
type AcctserialResult struct {
	RetCode     string              `json:"ret_code"`
	RetMsg      string              `json:"ret_msg"`
	OidPartner  string              `json:"oid_partner"`
	UserID      string              `json:"user_id"`
	PageNo      string              `json:"page_no"`
	TotalOutAmt string              `json:"total_out_amt"`
	TotalInAmt  string              `json:"total_in_amt"`
	TotalNum    string              `json:"total_num"`
	TotalPage   string              `json:"total_page"`
	AcctbalList []AcctserialAcctbal `json:"acctbal_list"`
}
