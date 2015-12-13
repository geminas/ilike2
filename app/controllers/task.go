package controllers

import "github.com/revel/revel"

var (
	folder = "../../../"
)

type Task struct {
	*revel.Controller
}

func (c Task) Index() revel.Result {
	return c.Render()
}

func (c Task) NewTask() {}

func (c Task) UpdateTask() {}

func (c Task) ShowData() {}
