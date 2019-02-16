package main

import (
	"extended-design-golang/bootstrap"
)

func main() {
	app := bootstrap.GetApp()
	app.ListenAndServe()

}
