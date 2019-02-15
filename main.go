package main

import (
	"extended-design-golang/app/boostrap"
)

func main() {
	app := boostrap.GetApp()
	app.ListenAndServe()

}
