package target

import (
	"fmt"
	"unicode/utf8"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func SSM(sess *session.Session, region string, values map[string]string, prefix string) error {
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	parameters := ParseParameter(values, prefix)
	for _, parameter := range parameters {
		results, err := PutParameter(ssmsvc, &parameter.Name, &parameter.Value, &parameter.ParamType)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(*results.Version)
	}
	return nil
}

func PutParameter(svc ssmiface.SSMAPI, name *string, value *string, paramType *string) (*ssm.PutParameterOutput, error) {
	results, err := svc.PutParameter(&ssm.PutParameterInput{
		Name:  name,
		Value: value,
		Type:  paramType,
	})

	return results, err
}

type ParameterStoreInput struct {
	Name      string
	Value     string
	ParamType string
}

func ParseParameter(values map[string]string, prefix string) []ParameterStoreInput {
	var output []ParameterStoreInput
	for key, value := range values {
		_paramType := "String"
		if key[0:1] == "*" {
			_, i := utf8.DecodeRuneInString(key)
			key = key[i:]
			_paramType = "SecureString"
		}
		_name := prefix + key
		output = append(output, ParameterStoreInput{Name: _name, Value: value, ParamType: _paramType})
	}
	return output
}
