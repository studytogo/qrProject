package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//qrUrl := `static/img/t.png`
	//err := qrcode.WriteFile(`http://106.12.38.239:9999/?name=zzg`, qrcode.Medium, 256, qrUrl)
	//if err != nil {
	//	fmt.Println("write error")
	//}
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "qr.tpl"
}

func (c *MainController) FormData() {
	fmt.Println("到这里了。。。。", c.GetString("tel"))
	tel := c.GetString("tel")
	qrUrl := `static/img/` + tel + `.png`
	err := qrcode.WriteFile(`http://106.12.38.239:8888/tel?tel=`+tel, qrcode.Medium, 256, qrUrl)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["tel"] = tel
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "result.tpl"
}

func (c *MainController) TelData() {
	c.Data["tel"] = c.GetString("tel")
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "show.tpl"
}
