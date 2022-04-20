package primality

import (
	"log"

	"github.com/rufusclark/prime/tool"
)

var cache []int

func init() {
	var err error
	cache, err = tool.LoadCache("prime.cache", 19)
	if err != nil {
		log.Fatal(err)
	}
}
