package webapge

type OpenacctApplyBasicInfo struct {
	RegPhone   string `json:"reg_phone"`
	UserName   string `json:"user_name"`
	IDType     string `json:"id_type"`
	IDNo       string `json:"id_no"`
	IDExp      string `json:"id_exp"`
	RegEmail   string `json:"reg_email"`
	Address    string `json:"address"`
	Occupation string `json:"occupation"`
}
