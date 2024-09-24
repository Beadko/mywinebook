/*
Copyright © 2024 Bearise Babra github.com/Beadko
*/
package main

import (
	"github.com/Beadko/mywinebook/cmd"
	"github.com/Beadko/mywinebook/db"
)

func main() {
	if err := db.OpenDatabase(); err != nil {
		panic(err.Error())
	}
	cmd.Execute()
}
