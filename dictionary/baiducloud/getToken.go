package baiducloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * @author:  wsb
 * @description: 调用百度云检测需要进行两步,第一步就是先获取Token,然后根据Token进行校验
 * @version:  1.0.1
 * @Date:  11:30 2021/12/8
 */
type Token struct {
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	SessionKey   string `json:"session_key"`
}

func GetToken() (token string) {

	t := Token{}

	var host = "https://aip.baidubce.com/oauth/2.0/token"
	var param = map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     "免费申请的 API Key", // 申请地址 https://cloud.baidu.com/product/textcensoring 需要实名认证才可以 然后新建应用即可
		"client_secret": "免费申请的 Secret Key",
	}

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(result, &t)
	fmt.Println(string(result))

	fmt.Println("token", t.AccessToken)

	return t.AccessToken

}
