package document

import (
	"text/scanner"
	"erevna/analysis"
)

type Content struct {
	data scanner.Scanner
	token rune
}

/*Create a content from string*/
func FromText(data string) Content {
	var cont Content = Content{}
	
	cont.data = analysis.ParseText(data)
	
	return cont
}

/*func fromFile(data) Content {
	var cont Content = Content{}
	
	cont.data = analysis.ParseText(data)
	
	return cont
}*/

//################################

func (c Content) NextToken() string {
	c.token = c.data.Scan()
	
	if(c.token == scanner.EOF) {
		return ""
	} else {
		return c.data.TokenText()
	}
}

