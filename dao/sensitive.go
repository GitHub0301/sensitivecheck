package dao

import (
	"github.com/astaxie/beego/client/orm"
	"sensitivecheck/models"
)

// Query Description
func QuerySensitive() (sens []models.Sensitive, err error) {

	o := orm.NewOrm()

	sql := "select * from `sensitive`"

	_, err = o.Raw(sql).QueryRows(&sens)

	if err != nil {
		return
	}

	return sens, nil
}
