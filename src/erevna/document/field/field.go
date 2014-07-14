package field

/*import (
	"container/list"
)*/

type Field struct {
	Name string
	Data interface{}
}

func New(name string, data interface{}) *Field {//tokens list.List data interface{}
	var ret = new(Field)
	ret.Name = name
	ret.Data = data
	
	return ret
}

func (t *Field) DataAsString() string {
    return t.Data.(string)
}