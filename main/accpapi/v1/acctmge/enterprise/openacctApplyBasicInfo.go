package enterprise

type OpenacctApplyBasicInfo struct {
	RegPhone         string `json:"reg_phone"`
	RegPhoneEvidence string `json:"reg_phone_evidence"`
	UserName         string `json:"user_name"`
	IDType           string `json:"id_type"`
	IDNo             string `json:"id_no"`
	IDExp            string `json:"id_exp"`
	IDStd            string `json:"id_std"`
	UnifiedCode      string `json:"unified_code"`
	Address          string `json:"address"`
	AddressPic       string `json:"address_pic"`
	RegEmail         string `json:"reg_email"`
	Password         string `json:"password"`
	RandomKey        string `json:"random_key"`
}
