package enterprise

type OpenacctApplyContactsInfo struct {
	ContactsName   string `json:"contacts_name"`
	ContactsPhone  string `json:"contacts_phone"`
	ContactsIDType string `json:"contacts_id_type"`
	ContactsIDNo   string `json:"contacts_idno"`
}
