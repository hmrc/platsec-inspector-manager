package factory

import (
	"fmt"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type UserInput struct {
	AwsAccount string
	Username string
	Region string
	Profile string
	Action string
	FilterName string
	FilterType string
	ComparissionOperator string
	MfaToken string
	SessionDuration int32
	UserConfig aws.Config
	UserContext context.Context
	ServiceCredentials ReturnedCredentials
	TemporaryCredentials Credentials
}

// ReturnedCredentials holds the client credentials for working with service clients
type ReturnedCredentials struct {
	AccessKeyId       string
	SecretAccessKeyId string
	SessionToken      string
	AssumedRole       string
}

// Credentials holds the session credentials used for client creation
type Credentials struct {
	AccessKeyId       string
	SecretAccessKeyId string
	SessionToken      string
}

// SetDefaultConfig loads up the credentials from the .aws folder
func (u *UserInput) SetDefaultConfig() {
	u.UserConfig, _ = config.LoadDefaultConfig(u.UserContext,
		config.WithSharedConfigProfile(u.Profile), config.WithRegion(u.Region))
}

//
func (u *UserInput) GetSerialNumber() *string {
	serialNumber := fmt.Sprintf("arn:aws:iam::%s:mfa/%s", u.AwsAccount, u.Username)
	return &serialNumber
}

// NewSTSClient returns an STS client generated from default config
func NewSTSClientFactory() func(cfg aws.Config) *sts.Client {
	return func(cfg aws.Config) *sts.Client {
		return sts.NewFromConfig(cfg)
	}
}
