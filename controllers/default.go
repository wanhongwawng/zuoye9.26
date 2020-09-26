package controllers

import (
	"beego02/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段：
}

func (c *MainController) Get() {
	//1，获取请求数据

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post() {
	var person models.Person
    dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
    if err!=nil{
    	c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err=json.Unmarshal(dataBytes,&person)
	if err!=nil {
		c.Ctx.WriteString("数据解释失败，请重试")
		return
	}
	c.Ctx.WriteString("数据解析成功")
	fmt.Println("姓名",person.Name)
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
}
