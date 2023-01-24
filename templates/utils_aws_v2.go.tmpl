package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func (u *Utils) AutoAWSConfig(awsRegion string) (aws.Config, error) {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
}

func (u *Utils) GetAWSConfig(awsRegion, accessKeyID, secretAccessKey, sessionToken string) (aws.Config, error) {
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
				SessionToken:    sessionToken,
			},
		}),
	)
}
