package controllers

import (
	"github.com/dogntao/rymz/routers"
)

type ProductController struct {
	str string
}

func init() {
	routers.Router.ConMap["product"] = &ProductController{}
}
