package controllers

import "github.com/revel/revel"

type Bootstrap struct {
	*revel.Controller
}

/**
 *param
 */
func (c Bootstrap) Index() revel.Result {
	return c.Render()
}
