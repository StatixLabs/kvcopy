package main

import (
	"clifig/pkg/auth"
	"clifig/pkg/source"
	"clifig/pkg/target"

	"flag"
	"fmt"
)

func main() {
	input := flag.String("input", "", "*REQUIRED* select where to get config values from.")
	output := flag.String("output", "", "*REQUIRED* select where the config values should go.")

	region := flag.String("region", "", "Set the AWS region.")
	profile := flag.String("profile", "", "Set the AWS SSO profile to use.")
	prefix := flag.String("prefix", "", "The prefix to use for SSM.")
	filePath := flag.String("filepath", "", "If using `file`, the path you want to read or write to.")

	flag.Parse()
	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	var parameters map[string]string
	switch checkInput := input; *checkInput {
	case "file":
		if *filePath == "" {
			fmt.Println("You need to set `filepath` flag when using `file` input.")
			return
		}
		parameters = source.File(*filePath)
	case "ssm":
		if *region == "" || *profile == "" || *prefix == "" {
			fmt.Print("You must define region, profile and prefix when using SSM as an output.")
			return
		}
		if string(*prefix)[:1] != "/" || string(*prefix)[len(*prefix)-1:] != "/" {
			fmt.Print("'prefix' must begin and end with a '/'.")
			return
		}
		sess, err := auth.AWS(*region, *profile)
		if err != nil {
			fmt.Println(err)
		}
		parameters, err = source.SSM(sess, *region, *prefix)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("%s.\n", "You must select an input for this to work...")
		return
	}

	switch checkOutput := output; *checkOutput {
	case "env":
		fmt.Println(target.Env(parameters))
	case "ssm":
		if *region == "" || *profile == "" || *prefix == "" {
			fmt.Print("You must define region, profile and prefix when using SSM as an output.")
			return
		}
		if string(*prefix)[:1] != "/" || string(*prefix)[len(*prefix)-1:] != "/" {
			fmt.Print("'prefix' must begin and end with a '/'.")
			return
		}
		sess, err := auth.AWS(*region, *profile)
		if err != nil {
			fmt.Println(err)
		}
		err = target.SSM(sess, *region, parameters, *prefix)
		if err != nil {
			fmt.Println(err)
		}
		return
	default:
		fmt.Printf("%s.\n", "You must select an output for this to work...")
		return
	}
}
