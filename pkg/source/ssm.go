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
	GetParameterByPathInput := ssm.GetParametersByPathInput{
		Path:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}
	if err := ssmsvc.GetParametersByPathPages(&GetParameterByPathInput, func(result *ssm.GetParametersByPathOutput, b bool) bool {
		for _, parameter := range result.Parameters {
			output[strings.ReplaceAll(*parameter.Name, name, "")] = *parameter.Value
		}
		return !b
	}); err != nil {
		return nil, err
	}
	return output, nil
}
