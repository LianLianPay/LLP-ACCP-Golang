package enterprise

type OpenacctApplyUboInfos struct {
	UboName     string `json:"ubo_name"`
	UboNameEN   string `json:"ubo_name_en"`
	UboPhone    string `json:"ubo_phone"`
	IDType      string `json:"id_type"`
	IDNo        string `json:"id_no"`
	IDEmblem    string `json:"id_emblem"`
	IDPortrait  string `json:"id_portrait"`
	UnifiedCode string `json:"unified_code"`
	IDExp       string `json:"id_exp"`
	IDIssue     string `json:"id_issue"`
	Address     string `json:"address"`
	UboEvidence string `json:"ubo_evidence"`
	UboType     string `json:"ubo_type"`
}
