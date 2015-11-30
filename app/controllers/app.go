package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	//"io/ioutil"
	"log"
	//"strings"
)

// var (
// 	DBNAME = "Info"
// )

type Info struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var db *bolt.DB

type App struct {
	*revel.Controller
}

func init() {}

///Page for Index
func (c App) Index() revel.Result {
	return c.Render()
}

///Page for people success apply
func (c App) Thanks(info string) revel.Result {
	return c.Render(info)
}

///Page for backend to view all the data
func (c App) ViewTable() revel.Result {
	var table = c.viewall()
	log.Println(table)
	return c.Render(table)
}

///API
func (c App) Query() revel.Result {
	a := c.Request.URL.Query()
	if a.Get("key") == "" {
		return c.RenderText("The key not found")
	} else {
		return c.RenderText(string(c.check(a.Get("key"))))
	}
}

///API
func (c App) Apply() revel.Result {
	var name = c.Request.PostForm.Get("name")
	var phone = c.Request.PostForm.Get("phone")
	var address = c.Request.PostForm.Get("address")
	var email = c.Request.PostForm.Get("email")
	var info = Info{
		Name:    name,
		Phone:   phone,
		Address: address,
		Email:   email,
	}
	var id = name + "-" + phone
	var j []byte
	var err error
	if j, err = json.Marshal(info); err != nil {
		return c.RenderText(err.Error())
	}
	if err := c.save(id, j); err != nil {
		return c.RenderText(err.Error())
	}
	return c.Redirect("/thanks/" + name)
}

func (c App) save(key string, val []byte) error {
	err := app.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.DBNAME))
		err := b.Put([]byte(key), val)
		return err
	})
	return err
}

func (c App) check(key string) []byte {
	var result []byte
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.DBNAME))
		result = b.Get([]byte(key))
		fmt.Printf("The answer is: %s\n", result)
		return nil
	})
	return result
}

func (c App) viewall() map[string]string {
	m := make(map[string]string)
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.DBNAME))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			m[string(k)] = string(v)
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
	log.Println(m)
	return m
}

func (c App) apply(data []byte) error {
	var v Info
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	log.Println(v)
	return nil
}
