package dao

import (
	"github.com/astaxie/beego/client/orm"
	"sensitivecheck/models"
)

func InitModel() {

	orm.RegisterModel(new(models.Sensitive)) //这里需要注册model不然新增是会报错

}
