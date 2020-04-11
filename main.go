package main

import "giao/vendors"
import _ "giao/routes"

func main()  {
	vendors.App.Start(":8080")
}
