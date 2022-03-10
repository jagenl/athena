package http

import "athena/pkg/config"

var (
	b *Base
	Host string
)

type Bl struct {

}
func init()  {
	Host =config.GetEnv("bl","https://t.gongfudou.com")
}

func (bl *Bl)Get(url string,params interface{}, headers interface{}) (response string) {
	//是否有加密的参数 需要添加到参数中的 header中的
	b.Params = params
	//panic: assignment to entry in nil map
	b.Header = headers
	return b.Get(Host+url)
}