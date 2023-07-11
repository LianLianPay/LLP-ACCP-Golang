package enterprise

type OpenacctApplyLegalreptInfo struct {
	LegalreptName   string `json:"legalrept_name"`
	LegalreptPhone  string `json:"legalrept_phone"`
	LegalreptIDType string `json:"legalrept_id_type"`
	LegalreptIDNo   string `json:"legalrept_idno"`
	IDEmblem        string `json:"id_emblem"`
	IDPortrait      string `json:"id_portrait"`
	LegalreptIDExp  string `json:"legalrept_idexp"`
	LegalreptStd    string `json:"legalrept_std"`
}
