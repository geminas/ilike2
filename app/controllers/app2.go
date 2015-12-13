package controllers

import (
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	//"io/ioutil"
	"errors"
	"log"
	"net/http"
	"strings"
)

var (
// correctUsername string = "admin"
// correctPassword string = "alliance"
)

type Info2 struct {
	Source                string `json:"source"`
	Name                  string `json:"name"`
	Phone                 string `json:"phone"`
	Email                 string `json:"email"`
	Company               string `json:"company"`
	Position              string `json:"position"`
	Interest			  string `json:"interest"`
}

// var db *bolt.DB

type App2 struct {
	*revel.Controller
}

func init() {}

///Page for Index
func (c App2) Index() revel.Result {
	return c.Render()
}

func (c App2) ErrorInfo(info string) revel.Result {
	return c.Render(info)
}

///Page for people success apply
func (c App2) Thanks(info string) revel.Result {

	return c.Render(info)
}

///Page for backend to view all the data
func (c App2) ViewTable() revel.Result {
	var infomap = c.viewarray()
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
		return c.Render(infomap)
	} else {
		c.Response.Status = http.StatusUnauthorized
		c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="新盟"`)
		return c.RenderError(errors.New("401: Not authorized"))
	}

}

// func (c App2) BasicAuth() revel.Result {
// 	//c.Response.ContentType = "Application/text"
// 	if auth := c.Request.Header.Get("Authorization"); auth != "" {
// 		// Split up the string to get just the data, then get the credentials
// 		username, password, err := getCredentials(strings.Split(auth, " ")[1])
// 		if err != nil {
// 			return c.RenderError(err)
// 		}
// 		if username != correctUsername || password != correctPassword {
// 			c.Response.Status = http.StatusUnauthorized
// 			c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="revel"`)
// 			return c.RenderError(errors.New("401: Not authorized"))
// 		}
// 		return c.RenderText("username: %s\npassword: %s", username, password)
// 	} else {
// 		c.Response.Status = http.StatusUnauthorized
// 		c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="revel"`)
// 		return c.RenderError(errors.New("401: Not authorized"))
// 	}
// }

// func getCredentials(data string) (username, password string, err error) {
// 	decodedData, err := base64.StdEncoding.DecodeString(data)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	strData := strings.Split(string(decodedData), ":")
// 	username = strData[0]
// 	password = strData[1]
// 	return
// }

///API
func (c App2) Query() revel.Result {
	a := c.Request.URL.Query()
	if a.Get("key") == "" {
		return c.RenderText("The key not found")
	} else {
		return c.RenderText("ok")
	}
}

///API
func (c App2) Apply() revel.Result {
	var source = c.Request.PostForm.Get("source")
	var name = c.Request.PostForm.Get("name")
	var phone = c.Request.PostForm.Get("phone")
	var email = c.Request.PostForm.Get("email")
	var company = c.Request.PostForm.Get("company")
	var position = c.Request.PostForm.Get("position")
	var interest = c.Request.PostForm.Get("interest")
	//log.Println(c.Request.PostForm)
	//log.Println(name, phone, address, email, category, origin, sex, company, position, emergencycontact, emergencyphone)
	var info = Info2{
		Source: 			   source,
		Name:                  name,
		Phone:                 phone,
		Email:                 email,
		Company:               company,
		Position:              position,
		Interest:			   interest,
	}
	var id = name + "-" + phone
	if b := c.check(id, "info"); len(b) != 0 {
		//return c.RenderError(errors.New("用户被重复申请"))
		return c.Redirect("/errorinfo/用户被重复申请")
	}
	if b := c.check(email, "email"); len(b) != 0 {
		return c.Redirect("/errorinfo/邮箱被重复使用")
	}
	if b := c.check(phone, "phone"); len(b) != 0 {
		return c.Redirect("/errorinfo/电话被重复使用")
	}

	var j []byte
	var err error
	if j, err = json.Marshal(info); err != nil {
		return c.RenderError(err)
	}
	if err := c.update(id, j, "info2"); err != nil {
		c.update(id, []byte{}, "info2")
		return c.RenderError(err)
	}
	if err := c.update(email, j, "email2"); err != nil {
		c.update(id, []byte{}, "info2")
		c.update(email, []byte{}, "email2")
		return c.RenderError(err)
	}
	if err := c.update(phone, j, "phone2"); err != nil {
		c.update(id, []byte{}, "info2")
		c.update(email, []byte{}, "email2")
		c.update(phone, []byte{}, "phone2")
		return c.RenderError(err)
	}
	var r = name
	return c.Redirect("/thanks/" + r)
}

func (c App2) update(key string, val []byte, table string) error {
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

// func (c App2) save(key string, val []byte) error {
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

func (c App2) check(key string, table string) []byte {
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

func (c App2) viewall() map[string]string {
	m := make(map[string]string)
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("info2"))
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

func (c App2) viewarray() string {
	m := "["
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("info2"))
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

func (c App2) apply(data []byte) error {
	var v Info
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	log.Println(v)
	return nil
}
