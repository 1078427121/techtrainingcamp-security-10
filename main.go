package main

import (
	"fmt"

	"techtrainingcamp-security-10/internal/resource"
	"techtrainingcamp-security-10/internal/route"
)

func main() {
	server, err := resource.NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}
	router, err := route.NewRoute(server.Service)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
	server.Close()
	return
}
