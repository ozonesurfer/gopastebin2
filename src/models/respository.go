// respository
package models

import (
	//	"fmt"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"log"
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
}

func GetDb() (myDB *db.DB, err error) {
	myDB, err = db.OpenDB("/data/pastebin")
	return
}

func GetAll() (pastes []Paste) {
	myDB, err1 := GetDb()
	if err1 != nil {
		log.Fatal("Error on database start - GetAll():", err1)
	}
	col := myDB.Use(PASTES)
	var query interface{}
	result := make(map[int]struct{})
	json.Unmarshal([]byte(`"all"`), &query)
	db.EvalQuery(query, col, &result)
	var docs []Paste
	for id := range result {

		doc, _ := col.Read(id)
		docObj := Paste{Id: id, Title: doc[TITLE].(string), Content: doc[CONTENT].(string)}
		docs = append(docs, docObj)
	}
	return docs
}
