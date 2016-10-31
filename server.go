package main

import (
	"github.com/go-martini/martini"
	_ "net/http/pprof"

	"code.oa.com/sean/api-doc/router"
)

func getMartini() *martini.ClassicMartini {
	m := martini.Classic()
	return m
}

func main() {
	m := getMartini()

	router.InitRouter(m)

	m.RunOnAddr(":3500")
}
