package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	//"io/ioutil"
	"errors"
	"log"
	//"strings"
)

// var (
// 	DBNAME = "Info"
// )

type Info struct {
	Name                  string `json:"name"`
	Phone                 string `json:"phone"`
	Company               string `json:"company"`
	Email                 string `json:"email"`
	Address               string `json:"address"`
	Category              string `josn:"category"`
	Origin                string `json:"origin"`
	Sex                   string `json:"sex"`
	Position              string `json:"position"`
	EmergencyContact      string `json:"emergency_contact"`
	EmergencyContactPhone string `json:"emergency_contact_phone"`
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
	var infomap = c.viewarray()
	//log.Println(table)
	// b, err := json.Marshal(table)
	// if err != nil {
	// 	return c.RenderError(err)
	// }
	// log.Println(string(b))
	// var infomap = string(b)
	return c.Render(infomap)
}

///API
func (c App) Query() revel.Result {
	a := c.Request.URL.Query()
	if a.Get("key") == "" {
		return c.RenderText("The key not found")
	} else {
		return c.RenderText("ok")
	}
}

///API
func (c App) Apply() revel.Result {
	var name = c.Request.PostForm.Get("name")
	var phone = c.Request.PostForm.Get("phone")
	var address = c.Request.PostForm.Get("address")
	var email = c.Request.PostForm.Get("email")
	var category = c.Request.PostForm.Get("category")
	var origin = c.Request.PostForm.Get("origin")
	var sex = c.Request.PostForm.Get("sex")
	var company = c.Request.PostForm.Get("company")
	var position = c.Request.PostForm.Get("position")
	var emergencycontact = c.Request.PostForm.Get("emergency_contact")
	var emergencyphone = c.Request.PostForm.Get("emergency_contact_phone")
	//log.Println(c.Request.PostForm)
	//log.Println(name, phone, address, email, category, origin, sex, company, position, emergencycontact, emergencyphone)
	var info = Info{
		Name:                  name,
		Phone:                 phone,
		Address:               address,
		Email:                 email,
		Category:              category,
		Origin:                origin,
		Sex:                   sex,
		Company:               company,
		Position:              position,
		EmergencyContact:      emergencycontact,
		EmergencyContactPhone: emergencyphone,
	}
	var id = name + "-" + phone
	if b := c.check(id, "info"); len(b) != 0 {
		return c.RenderError(errors.New("用户被重复申请"))
	}
	if b := c.check(email, "email"); len(b) != 0 {
		return c.RenderError(errors.New("邮箱被重复使用"))
	}
	if b := c.check(phone, "phone"); len(b) != 0 {
		return c.RenderError(errors.New("电话被重复使用"))
	}

	var j []byte
	var err error
	if j, err = json.Marshal(info); err != nil {
		return c.RenderError(err)
	}
	if err := c.update(id, j, "info"); err != nil {
		c.update(id, []byte{}, "info")
		return c.RenderError(err)
	}
	if err := c.update(email, j, "email"); err != nil {
		c.update(id, []byte{}, "info")
		c.update(email, []byte{}, "email")
		return c.RenderError(err)
	}
	if err := c.update(phone, j, "phone"); err != nil {
		c.update(id, []byte{}, "info")
		c.update(email, []byte{}, "email")
		c.update(phone, []byte{}, "phone")
		return c.RenderError(err)
	}

	return c.Redirect("/thanks/" + name)
}

func (c App) update(key string, val []byte, table string) error {
	log.Println("updating ", table)
	err := app.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			log.Println(err)
			return nil
		}
		err = b.Put([]byte(key), val)
		return err
	})
	return err
}

// func (c App) save(key string, val []byte) error {
// 	err := app.DB.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(app.DBNAME))
// 		c := b.Get([]byte(key))
// 		log.Println(c)
// 		if len(c) != 0 {
// 			return errors.New("无法提交,因为已经被提交过")
// 		}
// 		err := b.Put([]byte(key), val)
// 		return err
// 	})
// 	return err
// }

func (c App) check(key string, table string) []byte {
	var result []byte
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return nil
		}
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
			//fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
	//log.Println(m)
	return m
}

func (c App) viewarray() string {
	m := "["
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("info"))
		if b == nil {
			return nil
		}
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			m += string(v)
			m += ","
			//fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
	var bm = []byte(m)
	bm[len(bm)-1] = ']'
	return string(bm)
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
