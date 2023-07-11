package cnapscode

/**
 * 大额行号查询 请求参数
 */
type QueryParams struct {
	Timestamp   string `json:"timestamp"`
	OidPartner  string `json:"oid_partner"`
	BankCode    string `json:"bank_code,omitempty"`
	BrabankName string `json:"brabank_name,omitempty"`
	CityCode    string `json:"city_code,omitempty"`
}
