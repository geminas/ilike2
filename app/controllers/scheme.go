package controllers

import (
	//"encoding/json"
	"encoding/binary"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"time"
)

var testjson = `{
    "name":"活动名字",
    "describe":"活动描述",
    "code":"",
    "deleted":false,
    "bgimg":"/public/img/ljwl_background.jpg",
    "fields":[     {  
         "label":"下拉选项",
         "field_type":"dropdown",
         "required":true,
         "name":"c1",
         "field_options":{  
            "options":[  
               {  
                  "label":"选项1",
                  "checked":false
               },
               {  
                  "label":"选项2",
                  "checked":false
               },
               {  
                  "label":"选项3",
                  "checked":false
               }
            ],
            "include_blank_option":false
         },
         "cid":"c36"
      },
      {  
         "label":"这是输入框",
         "field_type":"text",
         "required":true,
         "name":"c2",
         "field_options":{  
            "size":"medium",
            "description":"",
            "minlength":"12",
            "maxlength":"22",
            "min_max_length_units":"words"
         },
         "cid":"c45"
      },
      {  
         "label":"这是单选框",
         "field_type":"radio",
         "required":true,
         "name":"c3",
         "field_options":{  
            "options":[  
               {  
                  "label":"选项1",
                  "checked":false
               },
               {  
                  "label":"选项2",
                  "checked":false
               },
               {  
                  "label":"选项3",
                  "checked":false
               }
            ],
            "description":"",
            "include_other_option":false
         },
         "cid":"c49"
      },
      {  
         "label":"这个是多选框",
         "field_type":"checkboxes",
         "required":true,
         "name":"c5",
         "field_options":{  
            "options":[  
               {  
                  "label":"选项1",
                  "checked":false
               },
               {  
                  "label":"选项2",
                  "checked":false
               },
               {  
                  "label":"选项3",
                  "checked":false
               },
               {  
                  "label":"选项4",
                  "checked":false
               }
            ],
            "description":""
         },
         "cid":"c57"
      },
      {  
         "label":"这个是下拉列表",
         "field_type":"dropdown",
         "required":true,
         "name":"c6",
         "field_options":{  
            "options":[  
               {  
                  "label":"选项1",
                  "checked":false
               },
               {  
                  "label":"选项2",
                  "checked":false
               },
               {  
                  "label":"选项3",
                  "checked":false
               },
               {  
                  "label":"选项4",
                  "checked":false
               }
            ],
            "include_blank_option":false,
            "description":""
         },
         "cid":"c61"
      }
]
}`

type Scheme struct {
	*revel.Controller
}

type Task struct {
	Name     string        `json:"name"`
	Describe string        `json:"describe"`
	Items    []interface{} `json:"items"`
}

func (c Scheme) Index() revel.Result {
	return c.Render()
}
func (c Scheme) GetTasks() revel.Result {
	//ts := c.check(key, table)
	return c.RenderJson(c.checktable("task"))
}
func (c Scheme) GetTask(name string) revel.Result {
	return c.RenderJson(string(c.Gettask(name)))
}

func (c Scheme) PostTask() revel.Result {
	b, err := ioutil.ReadAll(c.Request.Body)
	id := c.Params.Get("id")
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	if id == "" {
		return c.RenderJson(app.JsonResp{
			1,
			"The id is null",
			"",
			"",
		})
	}
	err = c.updatetask(id, b)
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	return c.RenderJson(app.JsonResp{
		0,
		id,
		"",
		testjson,
	})
}
func (c Scheme) NewTask() revel.Result {
	//var t = time.Now().Unix()
	//id := strconv.FormatInt(t, 10)
	id, err := c.newtaskautoid([]byte(testjson))
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	return c.RenderJson(app.JsonResp{
		0,
		"" + strconv.Itoa(id),
		"",
		testjson,
	})
}

func (c Scheme) SetPhone(name string, phone string) error {
	err := c.update(phone, []byte("true"), name+"-phone")
	return err
}

func (c Scheme) GetSetPhone(name string, phone string) revel.Result {
	err := c.SetPhone(name+"-phone", phone)
	if err != nil {
		return c.RenderError(err)
	} else {
		return c.RenderText("OK")
	}
}

func (c Scheme) PhoneExist(name string, phone string) revel.Result {
	p := c.check(phone, name+"-phone")
	// println(p)
	// println(name)
	// println(phone)
	if len(p) == 0 {
		return c.RenderJson(app.JsonResp{
			0,
			"null",
			"",
			"",
		})
	} else {
		return c.RenderJson(app.JsonResp{
			1,
			"exist",
			"",
			"",
		})
	}
}
func (c Scheme) PostTaskData(name string) revel.Result {
	p := c.Request.URL.Query().Get("phone")
	//println("phone")
	//println(p)

	if p != "" {
		if d := c.check(p, name+"-phone"); string(d) == "true" {
			return c.RenderJson(app.JsonResp{
				1,
				"exist",
				"",
				"",
			})
		}
		err := c.SetPhone(name, p)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		return c.RenderJson(app.JsonResp{
			1,
			"Phone Should Not Null",
			"",
			"",
		})
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	err = c.updatetaskdata(name, body)
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	return c.RenderJson(app.JsonResp{
		0,
		name,
		"",
		"",
	})

}

func (c Scheme) GetFile(name string) revel.Result {
	var imgpath = revel.Config.StringDefault("imgpath", "")
	var p = imgpath + name

	file, err := os.Open(p)
	if err != nil {
		return c.RenderError(err)
	}
	//defer file.Close()
	return c.RenderFile(file, revel.Inline)
}

func (c Scheme) GetTaskData(name string) revel.Result {
	table := c.checktable(name)
	return c.RenderJson(table)
}

func (c Scheme) Upload(name string) revel.Result {
	var imgpath = revel.Config.StringDefault("imgpath", "")

	log.Println("uploading")
	c.Request.ParseMultipartForm(32 << 20)
	file, handler, err := c.Request.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	defer file.Close()

	out, err := os.Create(imgpath + handler.Filename)
	if err != nil {
		log.Println(err)
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	return c.RenderJson(app.JsonResp{
		0,
		"/files/" + handler.Filename,
		"",
		"",
	})
}

func (c Scheme) CreateTask() revel.Result {
	b, err := ioutil.ReadAll(c.Request.Body)
	id := c.Params.Get("id")
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	if id == "" {
		return c.RenderJson(app.JsonResp{
			1,
			"The id is null",
			"",
			"",
		})
	}
	err = c.newtask(id, b)
	if err != nil {
		return c.RenderJson(app.JsonResp{
			1,
			err.Error(),
			"",
			"",
		})
	}
	return c.RenderJson(app.JsonResp{
		0,
		id,
		"",
		testjson,
	})

}

// func (c Scheme) Submit(id string) revel.Result {

// }

func (c Scheme) ViewData(name string) revel.Result {

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
		infomap := c.viewarray(name)
		log.Println(infomap)
		sc := string(c.Gettask(name))
		return c.Render(infomap, sc)

	} else {
		c.Response.Status = http.StatusUnauthorized
		c.Response.Out.Header().Set("WWW-Authenticate", `Basic realm="新盟"`)
		return c.RenderError(errors.New("401: Not authorized"))
	}

}

func (c Scheme) viewarray(table string) string {
	m := "["
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return nil
		}
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if string(v) != "" {
				m += string(v)
				m += ","
			}
			//fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
	var bm = []byte(m)
	if len(bm) > 1 {
		bm[len(bm)-1] = ']'
	} else {
		m += "]"
		bm = []byte(m)
	}
	return string(bm)
}

func (c Scheme) newtask(id string, scheme []byte) error {
	if b := c.check(id, "task"); len(b) != 0 {
		return errors.New("该任务已经存在")
	}
	err := c.update(id, scheme, "task")
	if err != nil {
		return err
	}
	return nil
}

func (c Scheme) newtaskautoid(scheme []byte) (int, error) {
	// if b := c.check(id, "task"); len(b) != 0 {
	// 	return errors.New("该任务已经存在")
	// }
	id, err := c.updatesequence("task", scheme)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (c Scheme) updatetask(id string, scheme []byte) error {
	err := c.update(id, scheme, "task")
	if err != nil {
		return err
	}
	return nil
}

func (c Scheme) updatetaskdata(id string, data []byte) error {
	_, err := c.updatesequence(id, data)
	return err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// func (c Scheme) updateautoid(table string, val []byte) err {
// 	err := app.DB.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists([]byte(table))
// 		if err != nil {
// 			return errors.New("create bucket:" + err.Error())
// 		}
// 		seq, err := b.NextSequence()
// 		var id = int(seq)
// 		if err != nil {
// 			return err
// 		}

// 		err = b.Put(itob(id), val)
// 		if err != nil {
// 			return errors.New("put into the table:" + err.Error())
// 		}
// 		return nil
// 	})
// 	return err
// }

func (c Scheme) updatesequence(table string, val []byte) (int, error) {
	var id = -1
	err := app.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return errors.New("create bucket:" + err.Error())
		}
		seq, err := b.NextSequence()
		id = int(seq)
		//println("id-------")
		//println(id)
		if err != nil {
			return err
		}

		err = b.Put([]byte(strconv.Itoa(id)), val)
		if err != nil {
			return errors.New("put into the table:" + err.Error())
		}
		return nil
	})
	return id, err
}

/*
func (this *BoltdbOutput) UpdateSequence(val []byte, table string) error {
	err := this.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return errors.New("create bucket:" + err.Error())
		}
		seq, err := b.NextSequence()
		if err != nil {
			return err
		}

		err = b.Put(itob(seq), val)
		if err != nil {
			return errors.New("put into the table:" + err.Error())
		}
		return nil
	})
	return err
}

*/

func (c Scheme) Gettask(name string) []byte {
	var result []byte
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("task"))
		if b == nil {
			return nil
		}
		result = b.Get([]byte(name))
		//fmt.Printf("The answer is: %s\n", result)
		return nil
	})

	return result
}

func (c Scheme) checktable(table string) map[string]string {
	m := make(map[string]string)
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return nil
		}
		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			m[string(k)] = string(v)
			return nil
		})
		return nil
	})
	return m
}

func (c Scheme) check(key string, table string) []byte {
	var result []byte
	app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return nil
		}
		result = b.Get([]byte(key))
		//fmt.Printf("The answer is: %s\n", result)
		return nil
	})
	return result
}

func (c Scheme) update(key string, val []byte, table string) error {
	log.Println("updating ", table, key)
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
