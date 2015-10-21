package controllers

import (
	"github.com/astaxie/beego"
)

type TrainTicketsController struct {
	beego.Controller
}

func (c *TrainTicketsController) Get() {
	c.TplNames = "traintickets.html"
}

func (c *TrainTicketsController) Post() {

}
