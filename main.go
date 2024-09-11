package main

import (
	"proxyserver/cmd"
)

func main() {
	cmd.Execute()
}

/*
func main() {
	srv := server.Create("5135")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
*/
