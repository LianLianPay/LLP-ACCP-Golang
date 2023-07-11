package txn

/**
 * 用户委托协议申请 请求参数
 */
type PapAgreeApplyAndModifyParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	TxnSeqNo   string `json:"txn_seqno"`
	TxnTime    string `json:"txn_time"`
	UserID     string `json:"user_id"`
	PapAgreeNo string `json:"pap_agree_no"`
	/*
	   交易发起渠道。
	   ANDROID
	   IOS
	   H5
	   PC
	   目前只支持H5。
	*/
	FlagChnl  string `json:"flag_chnl"`
	ReturnURL string `json:"return_url"`
	NotifyURL string `json:"notify_url"`
	// 签约信息
	PapSignInfo PapAgreeSignInfo `json:"pap_sign_info"`
}
