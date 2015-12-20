package controllers

import "github.com/revel/revel"

type Backend struct {
	*revel.Controller
}

func (c Backend) Index() revel.Result {
	return c.Render()
}
