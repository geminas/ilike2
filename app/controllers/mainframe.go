package controllers

import (
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
)

type MainFrame struct {
	*revel.Controller
}

func (c MainFrame) Index(name string) revel.Result {
	task := string(c.Gettask(name))
	return c.Render(task)

}

func (c MainFrame) Gettask(name string) []byte {
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
