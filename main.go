/*
Copyright © 2024 Bearise Babra github.com/Beadko
*/
package main

import (
	"github.com/Beadko/mywinebook/cmd"
	"github.com/Beadko/mywinebook/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
