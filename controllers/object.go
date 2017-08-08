package controllers

import (
	"math/rand"

	"github.com/astaxie/beego"
)

// Operations about object
type DocentePlantaController struct {
	beego.Controller
}

func (c *DocentePlantaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

func (o *DocentePlantaController) GetOne() {
	if(rand.Intn(2)==1){
		o.Data["json"] = true
	}else{
		o.Data["json"] = false
	}
	o.ServeJSON()
}