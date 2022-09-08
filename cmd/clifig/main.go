package main

import (
	"clifig/pkg/auth"
	"clifig/pkg/source"
	"fmt"
	"os"
)

func main() {
	region := os.Args[1]
	profile := os.Args[2]
	key := os.Args[3]
	sess, err := auth.AWS(region, profile)
	if err != nil {
		fmt.Println(err)
	}
	param, err := source.SSM(sess, region, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(param)
}
