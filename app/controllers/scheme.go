package controllers

import (
	"github.com/boltdb/bolt"
	"github.com/geminas/ilike2/app"
	"github.com/revel/revel"
	"log"
)

type Scheme struct {
	*revel.Controller
	SchemeJson string
	Name       string
}

func (c Scheme) Index() revel.Result {
	return c.Render()
}

// var db *bolt.DB

func NewScheme(name string) (*Scheme, error) {
	var s = new(Scheme)
	err := app.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("aabb"))
		if b == nil {
			log.Println("The bucket is not exists")
		}
		// v := b.Get([]byte("answer"))
		// fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	return s
}

func (c Scheme) TryNew() revel.Result {
	NewScheme("hello")
	return c.RenderText("ok")
}
