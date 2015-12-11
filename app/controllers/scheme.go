package controllers

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	"log"
)

var testjson = `{
    "name":"111",
    "describe":"2222",
    "items":[{
    "label": "参会类别",
    "field_type": "select",
    "required": true,
    "name":"category",
    "field_options": {
        "options": [{
            "label": "免费申请"
        },{
            "label":"媒体参会"
        }, {
            "label": "VIP参会¥2880"
        }, {
            "label": "VIP参会¥1880"
        }],
        "include_other_option": true
    },
    "cid": "c10"
},{
    "label": "参会机构来源",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"origin",
    "placeholder":"来源",
    "cid": "c6"
},{
    "label": "您的姓名",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"name",
    "placeholder":"姓名",
    "cid": "c6"
},{
    "label": "您的身份证号",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"idcard",
    "validator":"idcard",
    "placeholder":"身份证号",
    "cid": "c6"
},{
    "label": "联系电话",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"phone",
    "validator":"phone",
    "placeholder":"电话",
    "cid": "c6"
},{
    "label": "电子邮箱",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"email",
    "placeholder":"邮箱",
    "validator":"email",
    "cid": "c6"
},{
    "label": "您是先生/女士",
    "field_type": "radio-inline",
    "required": true,
    "name":"sex",
    "field_options": {
        "options": [{
            "label": "男"
        }, {
            "label": "女"
        }],
        "include_other_option": true
    },
    "cid": "c10"
},{
    "label": "公司名称",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"company",
    "placeholder":"公司名字",
    "cid": "c6"
},{
    "label": "职务",
    "field_type": "text",
    "required": true,
    "field_options": {},
    "name":"position",
    "placeholder":"职务",
    "cid": "c6"
},{
    "label": "紧急联系人姓名",
    "field_type": "text",
    "required": false,
    "field_options": {},
    "name":"emergency_contact",
    "placeholder":"姓名",
    "cid": "c6"
},{
    "label": "紧急联系人电话",
    "field_type": "text",
    "required": false,
    "field_options": {},
    "name":"emergency_contact_phone",
    "placeholder":"电话",
    "cid": "c6"
}]
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

// var db *bolt.DB

func (c Scheme) newtask(data []byte) error {
	var t Task
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	if b := c.check(t.Name, "task"); len(b) != 0 {
		return errors.New("该任务已经存在")
	}

	return nil
}

func (c Scheme) gettask(name string) []byte {
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

// func NewTask(name string) *Scheme {
// 	var s = new(Scheme)
// 	app.DB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("aabb"))
// 		if b == nil {
// 			log.Println("The bucket is not exists")
// 		}
// 		// v := b.Get([]byte("answer"))
// 		// fmt.Printf("The answer is: %s\n", v)
// 		return nil
// 	})
// 	return s
// }

func (c Scheme) TryNew() revel.Result {
	c.newtask([]byte(testjson))
	return c.RenderText("ok")
}
