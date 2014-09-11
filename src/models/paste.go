// paste
package models

import (
	//"fmt"
	//"github.com/HouzuoGuo/tiedot/db"
	"log"
	"time"
)

type MapDoc map[string]interface{}

type Paste struct {
	Id             int
	Title, Content string
	CreatedOn      time.Time
}

const TITLE = "title"
const CONTENT = "content"
const CREATED = "created"
const PASTES = "pastes"

func (this Paste) ToMap() (myMap MapDoc) {
	myMap = map[string]interface{}{TITLE: this.Title, CONTENT: this.Content, CREATED: this.CreatedOn}
	return
}

func (this Paste) Add() (id int, err error, createdOn time.Time) {
	myDB, err1 := GetDb()
	if err1 != nil {
		log.Fatal("Error on database start - Add():", err1)
	}
	col := myDB.Use(PASTES)
	now := time.Now()
	myMap := map[string]interface{}{TITLE: this.Title, CONTENT: this.Content, CREATED: now}
	id, err = col.Insert(myMap)
	createdOn = now
	return
}

func GetPaste(id int) (paste Paste) {
	myDB, err1 := GetDb()
	if err1 != nil {
		log.Fatal("Error on database start - GetPaste():", err1)
	}
	col := myDB.Use(PASTES)
	doc, _ := col.Read(id)
	theTime, _ := time.Parse(time.RFC3339, doc[CREATED].(string))
	paste = Paste{Id: id, Title: doc[TITLE].(string), Content: doc[CONTENT].(string), CreatedOn: theTime}
	return
}

/*
func main() {
	fmt.Println("Hello World!")
}
*/
