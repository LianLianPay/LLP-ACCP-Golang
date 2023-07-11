package random

/**
 * 随机因子获取 请求参数
 */
type GetRandomParams struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	/*
	   交易发起渠道。
	   ANDROID
	   IOS
	   H5
	   PC
	*/
	FlagChnl string `json:"flag_chnl"`
	PkgName  string `json:"pkg_name"`
	AppName  string `json:"app_name"`
	/*
	   加密算法。
	   RSA：国际通用RSA算法
	   SM2 ：国密算法
	   默认 RSA算法
	*/
	EncryptAlgorithm string `json:"encrypt_algorithm"`
}
