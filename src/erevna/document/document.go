package document

import (
	"erevna/document/field"
)

type Document struct {
	Uuid string
	Name string
	Fields []field.Field
	Content Content 
}

/**
	Creates a new document
**/
func NewDocument(name string) *Document {
	var ret = new(Document);
	ret.Fields = make([]field.Field, 0)
	ret.Name = name
	
	return ret;
}

func (doc *Document) SetContent(data string) {
	doc.Content = FromText(data)	
} 