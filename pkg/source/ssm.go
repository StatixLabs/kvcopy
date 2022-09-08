package source

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// createClient creates a Simple Systems Manager client
func SSM(sess *session.Session, region string, name string) (string, error) {
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(false),
	})
	return *param.Parameter.Value, err
}
