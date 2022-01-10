package thulium

import (
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, quark Quark)  {
	var proto *Prototype
	if v, ok := interface{}(quark).(*Prototype); ok {
		proto = v
	}
	quark.Render()
	fmt.Println(proto.Inner)
}
