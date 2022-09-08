package main

import (
	"clifig/pkg/auth"
	"clifig/pkg/source"
	"clifig/pkg/target"
	"flag"
	"fmt"
)

func main() {
	region := flag.String("region", "us-east-1", "Set the AWS region.")
	profile := flag.String("profile", "default", "Set the AWS SSO profile to use.")
	keyPointer := flag.String("ssm-parameter-path", "", "set the SSM parameter store key path. It must being and end with a '/'")
	flag.Parse()
	key := string(*keyPointer)
	if key[:1] != "/" {
		panic("'ssm-parameter-path' must begin with an '/'")
	}
	if key[len(key)-1:] != "/" {
		panic("'ssm-parameter-path' must end with an '/'")
	}

	sess, err := auth.AWS(*region, *profile)
	if err != nil {
		fmt.Println(err)
	}

	param, err := source.SSM(sess, *region, key)
	if err != nil {
		fmt.Println(err)
	}

	output := target.Env(param)
	fmt.Println(output)
}
