package cnapscode

/**
 * 大额行号查询 请求参数
 */
type QueryResult struct {
	RetCode  string      `json:"ret_code"`
	RetMsg   string      `json:"ret_msg"`
	BankCode string      `json:"bank_code"`
	CardList []BankCnaps `json:"card_list"`
}
