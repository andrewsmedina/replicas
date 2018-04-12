package main

import (
	"fmt"

	"github.com/andrewsmedina/replicas/finder"
)

func main() {
	fmt.Println("finding duplicated files")
	duplicated := finder.Find()
	fmt.Printf("%#v\n", duplicated)
}
