package auth

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

// getSession creates a session
func AWS(region, profile string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	})
	return sess, err
}
