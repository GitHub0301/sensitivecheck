package models

/**
 * @author:  wsb
 * @description: 数据库结构体
 * @version:  1.0.1
 * @Date:  11:09 2021/12/8
 */
type Sensitive struct {
	Id       int    `form:"Id"`
	Type     string `form:"Type"`
	Keywords string `form:"Keywords"`
}
