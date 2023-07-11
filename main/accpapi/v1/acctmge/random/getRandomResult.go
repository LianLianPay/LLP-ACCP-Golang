package random

/**
 * 随机因子获取 响应参数
 */
type GetRandomResult struct {
	RetCode          string `json:"ret_code"`
	RetMsg           string `json:"ret_msg"`
	OidPartner       string `json:"oid_partner"`
	UserID           string `json:"user_id"`
	RandomKey        string `json:"random_key"`
	RandomValue      string `json:"random_value"`
	License          string `json:"license"`
	MapArr           string `json:"map_arr"`
	RSAPublicContent string `json:"rsa_public_content"`
	SM2KeyHex        string `json:"sm2_key_hex"`
}
