package baiducloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
* @author:  wsb
* @description: 这里只是简单返回一个状态而已,详细的可以自己去研究,无非就是多返回几个字段而已
                最后面有百度云返回结果示例
* @version:  1.0.1
* @Date:  11:24 2021/12/8
*/

type Res struct {
	Conclusion     string `form:"Conclusion"`
	ConclusionType int    `form:"ConclusionType"`
}

func Check(txt string) int {

	res := Res{}

	var host = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"
	var accessToken = GetToken()
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("text", txt)
	sendData := sendBody.Form.Encode()
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	json.Unmarshal(result, &res)

	fmt.Println("校验结果", res)

	return res.ConclusionType
}

/*  正确的返回结果
{
   "log_id": 15556561295920002,
   "conclusion": "合规",
   "conclusionType": 1
}

或者

{
    "log_id": 15572142621780024,
    "conclusion": "合规",
    "conclusionType": 1,
    "data": [{
        "type": 14,
        "subType": 0,
        "conclusion": "合规",
        "conclusionType": 1,
        "msg": "自定义文本白名单审核通过",
        "hits": [{
            "datasetName": "SLK-测试-自定义文本白名单",
            "words": ["张三"]
        }]
    }]
}

*/

/*  错误返回结果
{
   "log_id": 15656780617612718,
   "conclusion": "不合规",
   "conclusionType": 2,
   "data": [{
	   "type": 11,
	   "subType": 0,
	   "conclusion": "不合规",
	   "conclusionType": 2,
	   "msg": "存在百度官方默认违禁词库不合规",
	   "hits": [{
		   "datasetName": "百度默认黑词库",
		   "words": ["免费翻墙"]
	   }]
   }, {
	   "type": 12,
	   "subType": 2,
	   "conclusion": "不合规",
	   "conclusionType": 2,
	   "msg": "存在文本色情不合规",
	   "hits": [{
		   "datasetName": "百度默认文本反作弊库",
		   "probability": 1.0,
		   "words": ["电话 找小姐"]
	   }]
   }, {
	   "type": 12,
	   "subType": 3,
	   "conclusion": "不合规",
	   "conclusionType": 2,
	   "msg": "存在政治敏感不合规",
	   "hits": [
			   {
				  "modelHitPositions":[
					   [
						   0,
						   6,
						   0.9998
					   ]
				   ],
				   "wordHitPositions":[
					   {
						   "keyword":"法轮功",
						   "positions":[
							   [
								   0,
								   2
							   ]
						   ],
						   "label":"201102"
					   }
				   ],
				   "probability":"1.0",
				   "datasetName":"百度默认文本反作弊库",
				   "words":[
					   "法轮功"
				   ]
			   }
		   ]
   }, {
	   "type": 12,
	   "subType": 4,
	   "conclusion": "不合规",
	   "conclusionType": 2,
	   "msg": "存在恶意推广不合规",
	   "hits": [{
		   "probability": 1.0,
		   "datasetName": "百度默认文本反作弊库",
		   "words": [""]
	   }]
   }, {
	   "type": 13,
	   "subType": 0,
	   "conclusion": "不合规",
	   "conclusionType": 2,
	   "msg": "存在自定义文本黑名单不合规",
	   "hits": [{
		   "datasetName": "SLK-测试-自定义黑名单",
		   "words": ["我是你爹", "他妈的"]
	   }]
   }]
}
*/
