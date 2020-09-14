// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(
		http.Dir("../../ui"), vfsgen.Options{
			PackageName:  "assets",
			BuildTags:    "!dev",
			VariableName: "Assets",
		})

	if err != nil {
		log.Fatalln(err)
	}
}
