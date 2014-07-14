package main 

import (
	"fmt"
	"erevna/document"
	"erevna/document/field"
	"erevna/store"
)

func main() {
	//var field = field.New("nome", "dados")
	//fmt.Printf("Hello, world.\n έρευνα")
	//fmt.Printf("Field teste:" + field.DataAsString())
	
	testeInsert()
}

func testeInsert() {
	doc := document.NewDocument("teste123")
	doc.SetContent("teste teste teste123 abraao isvi ")
	
	st := store.Create(store.DEFAULT)
	st.Open()
	st.Persist(doc)
	st.Close()
}

func testeSearch() {
	
}