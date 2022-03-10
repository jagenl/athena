package http

import (
	httplib "net/http"
)

type Base struct {
	Params interface{}
	Header interface{}
}


func (h * Base)Get(url string,headers ...interface{}) (response string) {
	if h.Params !=nil{
		//for k,v :=range h.Header{
		//	httplib.Header{}.Add(k,v)
		//	fmt.Println("key:",k,",value:",v)
		//}
		//
	}
	response ="ddd"
	return
}

func (h * Base)Post(url string,headers ...interface{}) (response string) {
	httplib.Header{}.Add("token","xxxx")
	response ="ddd"
	return
}