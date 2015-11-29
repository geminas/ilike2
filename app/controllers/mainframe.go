package controllers

import "github.com/revel/revel"

type MainFrame struct {
	*revel.Controller
}

func (c MainFrame) Index() revel.Result {
	return c.Render()
}
