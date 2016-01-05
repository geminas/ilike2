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
	"os"
	"strconv"
	"time"
)

var testjson = `{
    "name":"活动名字",
    "describe":"活动描述",
    "deleted":false,
    "bgimg":"/public/img/ljwl_background.jpg",
    "fields":[     {  
         "label":"下拉选项",
         "field_type":"dropdown",
         "required":true,
         "field_options":{  
            "options":[  
               {  
                  "label":"选项1",
                  "checked":true
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
         "field_options":{  
            "size":"medium",
            "description":"这里是一些额外数据",
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
            "description":"这里是一些额外数据",
            "include_other_option":true
         },
         "cid":"c49"
      },
      {  
         "label":"这个也是单选",
         "field_type":"checkboxes",
         "required":true,
         "field_options":{  
            "options":[  
               {  
                  "label":"选择a",
                  "checked":false
               },
               {  
                  "label":"选择b",
                  "checked":true
               }
            ],
            "description":"在这里输入一些额外的数据",
            "include_other_option":true
         },
         "cid":"c53"
      },
      {  
         "label":"这个是多选框",
         "field_type":"radio",
         "required":true,
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
            "description":"在这里输入一些东西"
         },
         "cid":"c57"
      },
      {  
         "label":"这个是下拉列表",
         "field_type":"dropdown",
         "required":true,
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
            "description":"在这里输入一些东西"
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
	var t = time.Now().Unix()
	id := strconv.FormatInt(t, 10)
	err := c.newtask(id, []byte(testjson))
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

func (c Scheme) PostTaskData(name string) revel.Result {
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

func (c Scheme) GetTaskData(name string) revel.Result {
	table := c.checktable(name)
	return c.RenderJson(table)
}

func (c Scheme) Upload(name string) revel.Result {
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

	out, err := os.Create("/Users/deepglint/gocode/src/github.com/geminas/ilike2/public/img/" + handler.Filename)
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
		"/public/img/" + handler.Filename,
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
	infomap := c.viewarray(name)
	log.Println(infomap)
	sc := string(c.Gettask(name))
	return c.Render(infomap, sc)
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

func (c Scheme) updatetask(id string, scheme []byte) error {
	err := c.update(id, scheme, "task")
	if err != nil {
		return err
	}
	return nil
}

func (c Scheme) updatetaskdata(id string, data []byte) error {
	err := c.updatesequence(id, data)
	return err
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (c Scheme) updatesequence(table string, val []byte) error {
	err := app.DB.Update(func(tx *bolt.Tx) error {
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
