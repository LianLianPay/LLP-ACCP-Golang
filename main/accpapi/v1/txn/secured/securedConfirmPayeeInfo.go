package secured

type SecuredConfirmPayeeInfo struct {
	PayeeID         string `json:"payee_id"`
	PayeeType       string `json:"payee_type"`
	PayeeOidPartner string `json:"payee_oid_partner"`
	PayeeAmount     string `json:"payee_amount"`
}
