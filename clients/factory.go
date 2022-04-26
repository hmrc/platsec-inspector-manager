package factory

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type UserInput struct {
	AwsAccount string
	Usernname string
	Region string
	Profile string
	Action string
	FilterName string
	FilterType string
	ComparissionOperator string
	MfaToken string
	UserConfig aws.Config
	UserContext context.Context
}

// SetDefaultConfig loads up the credentials from the .aws folder
func (u *UserInput) SetDefaultConfig() {
	u.UserConfig, _ = config.LoadDefaultConfig(u.UserContext,
		config.WithSharedConfigProfile(u.Profile), config.WithRegion(u.Region))
}