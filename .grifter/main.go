
package main

import _ "github.com/nosnibor89/awww-muscle/grifts"
import "os"
import "log"
import "github.com/markbates/grift/grift"
import "path/filepath"

func main() {
	grift.CommandName = "grift"
	if err := os.Chdir(filepath.Dir("D:/Development/Go/src/github.com/nosnibor89/awww-muscle/grifts")); err != nil {
	  log.Fatal(err)
	}
	err := grift.Exec(os.Args[1:], false)
	if err != nil {
		log.Fatal(err)
	}
}