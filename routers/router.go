package routers

import (
	"github.com/dogntao/rymz/dtgo"
)

var Router *dtgo.MainControler

func init() {
	Router = &dtgo.MainControler{}
	Router.ConMap = make(map[string]interface{})
	Router.Params = make(map[string]string)
}
