package baidu

const (
	apiGetAccessToken = "/oauth/2.0/token"
)

// TokenResponse 获取 access_token 成功返回数据
type TokenResponse struct {
	CommonError
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是2592000秒之内的值。
}

// GetAccessToken 获取百度开放平台全局唯一后台接口调用凭据（access_token）。
// 调调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存，注意缓存。
func GetAccessToken(appID, secret string) (*TokenResponse, error) {
	api := baseURL + apiGetAccessToken
	return getAccessToken(appID, secret, api)
}

func getAccessToken(appID, secret, api string) (*TokenResponse, error) {
	queries := requestQueries{
		"client_id": 		appID,
		"client_secret":  	secret,
		"grant_type":  		"client_credentials",
	}

	url, err := encodeURL(api, queries)

	if err != nil {
		return nil, err
	}

	res := new(TokenResponse)

	if err := getJSON(url, res); err != nil{
		return nil, err
	}

	return res, nil
}