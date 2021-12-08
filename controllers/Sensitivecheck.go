package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"sensitivecheck/dao"
	"sensitivecheck/util"
	"strings"
)

/**
 * @author:  wsb
 * @description: 方法分为两种 methodCheck sqlCheck
 * @version:  1.0.1
 * @Date:  11:10 2021/12/8
 */
type SensitiveCheckController struct {
	beego.Controller
}

func (c *SensitiveCheckController) Get() {
	c.TplName = "index.tpl"
}

func (c *SensitiveCheckController) Post() {

	str := c.GetString("txt")

	fmt.Println(str)

	genreType, genreWord, b := c.methodCheck(str)
	//genreType, genreWord, b := c.sqlCheck(str)

	fmt.Println("敏感类型 -> ", genreType)
	fmt.Println("敏感词 -> ", genreWord)
	fmt.Println("是否敏感 -> ", b)

	c.Data["result"] = b

	c.TplName = "index.tpl"
}

/*
genreType  敏感类型
genreWord  敏感词
b          是否敏感
*/

func (c *SensitiveCheckController) methodCheck(str string) (genreType string, genreWord string, b bool) {

	genre, keywords := util.Sensitivewords()

	for k := 0; k < len(keywords); k++ {
		if strings.Contains(str, keywords[k]) {
			return genre[k], keywords[k], true
		}
	}

	return "", "", false
}

func (c *SensitiveCheckController) sqlCheck(str string) (genreType string, genreWord string, b bool) {

	sens, _ := dao.QuerySensitive()

	//fmt.Println("输出结果", sens)

	for k := 0; k < len(sens); k++ {
		if strings.Contains(str, sens[k].Keywords) {
			return sens[k].Type, sens[k].Keywords, true
		}
	}

	return "", "", false
}
