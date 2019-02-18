package controllers

import (
	"github.com/dogntao/rymz/routers"
)

type IndexController struct {
	str string
}

func init() {
	routers.Router.ConMap["index"] = &IndexController{}
}
