// respository
package models

import (
	//	"fmt"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"log"
	"sort"
	"time"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func init() {
	myDB, err := GetDb()
	defer myDB.Close()
	if err != nil {
		log.Fatal("Error on database start:", err)
	} else {
		log.Println("Database start was successful")
	}
	collections := myDB.AllCols()
	exists := false
	for _, col := range collections {
		if col == PASTES {
			exists = true
		}
	}
	if exists == false {
		myDB.Create(PASTES)
		log.Println("Collection was successfully created")
	} else {
		log.Println("Collection already exists")
	}
	col := myDB.Use(PASTES)
	col.Index([]string{"title"})
}

func GetDb() (myDB *db.DB, err error) {
	myDB, err = db.OpenDB("/data/pastebin")
	return
}

type ByCreated []Paste

func (this ByCreated) Len() int           { return len(this) }
func (this ByCreated) Less(i, j int) bool { return this[i].CreatedOn.After(this[j].CreatedOn) }
func (this ByCreated) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func GetAll() (pastes []Paste) {
	myDB, err1 := GetDb()
	if err1 != nil {
		log.Fatal("Error on database start - GetAll():", err1)
	}
	col := myDB.Use(PASTES)
	var query interface{}
	result := make(map[int]struct{})
	err := json.Unmarshal([]byte(`"all"`), &query)
	//	err := json.Unmarshal([]byte(`{"limit": 5}`), &query)
	if err != nil {
		log.Fatal("json error:", err)
	}
	db.EvalQuery(query, col, &result)
	var docs []Paste
	for id := range result {

		doc, _ := col.Read(id)
		theTime, _ := time.Parse(time.RFC3339, doc[CREATED].(string))
		docObj := Paste{Id: id, Title: doc[TITLE].(string), Content: doc[CONTENT].(string), CreatedOn: theTime}
		docs = append(docs, docObj)
	}
	sort.Sort(ByCreated(docs))
	q := docs[0:5]
	return q
}
