package main

import (
	"github.com/iyaoo/go-IMChat/admin/router"
)

func main() {
	r := router.UserRouter()
	r.Run()
}
