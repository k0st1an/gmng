package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func newGlacierSession() *session.Session {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigDisable,
			Config: aws.Config{
				// LogLevel: aws.LogLevel(aws.LogDebugWithEventStreamBody),
				// LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
				// LogLevel: aws.LogLevel(aws.LogDebugWithSigning),
				Region: aws.String(conf.Region),
				Credentials: credentials.NewStaticCredentials(
					conf.AccessKeyID,
					conf.SecretAcccessKey,
					"",
				),
			},
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	return sess
}
