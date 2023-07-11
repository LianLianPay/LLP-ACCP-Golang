package enterprise

type OpenacctApplyBusinessInfo struct {
	Scale             string `json:"scale"`
	IndustryCode      string `json:"industry_code"`
	RegisteredCapital string `json:"registered_capital"`
	BusinessScope     string `json:"business_scope"`
	OpenPermit        string `json:"open_permit"`
}
