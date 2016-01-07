package controllers

import (
	"errors"
	"github.com/revel/revel"
	"net/http"
	"strings"
)

type Backend struct {
	*revel.Controller
}

func (c Backend) Index() revel.Result {
	if auth := c.Request.Header.Get("Authorization"); auth != "" {
		// Split up the string to get just the data, then get the credentials
		username, password, err := getCredentials(strings.Split(auth, " ")[1])
		if err != nil {
			return c.RenderError(err)
		}
		if username != correctUsername || password != correctPassword {
			c.Response.Status = http.StatusUnauthorized
			c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="revel"`)
			return c.RenderError(errors.New("401: Not authorized"))
		}
		return c.Render()
	} else {
		c.Response.Status = http.StatusUnauthorized
		c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="新盟"`)
		return c.RenderError(errors.New("401: Not authorized"))
	}
}
