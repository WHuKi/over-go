package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 这个函数下面还会用到，请记住函数名
func Oauth(w http.ResponseWriter, r *http.Request) {

	var err error

	// 获取 code
	var code = r.URL.Query().Get("code")

	// 通过 code, 获取 token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	var token *Token
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}

	//  将用户信息返回前端
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println("在将用户信息([]byte)返回前端时发生错误，错误信息为:", err)
		return
	}

}

// 通过code获取token认证url
func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		"43f2dff544a002b3d28b", "560d4a75043f62304f6bde5faa7190af23fc574d", code,
	)
}

// 获取 token
func GetToken(url string) (*Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}
