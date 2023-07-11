package individual

type OpenacctApplyBasicInfo struct {
	RegPhone    string `json:"reg_phone"`
	UserName    string `json:"user_name"`
	IDType      string `json:"id_type"`
	IDNo        string `json:"id_no"`
	IDExp       string `json:"id_exp"`
	IDStd       string `json:"id_std"`
	IDAuthority string `json:"id_authority"`
	AreaCode    string `json:"area_code"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
	IDEmblem    string `json:"id_emblem"`
	IDPortrait  string `json:"id_portrait"`
}
