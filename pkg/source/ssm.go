package source

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// createClient creates a Simple Systems Manager client
func SSM(sess *session.Session, region string, name string) (map[string]string, error) {
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	output := make(map[string]string)
	param, err := ssmsvc.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return output, err
	}
	for _, parameter := range param.Parameters {
		output[strings.ReplaceAll(*parameter.Name, name, "")] = *parameter.Value
	}
	return output, nil
}
