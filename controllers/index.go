package controllers

import (
	"fmt"

	"github.com/dogntao/rymz/routers"
)

type IndexController struct {
	str string
}

func init() {
	routers.Router.ConMap["Index"] = &IndexController{}
}

func (ind *IndexController) Index() {
	fmt.Println("index/index")
}
