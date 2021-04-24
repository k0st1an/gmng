package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glacier"
	cli "github.com/urfave/cli/v2"
)

// DeleteArchive ..
func DeleteArchive(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())
	_, err := glr.DeleteArchive(&glacier.DeleteArchiveInput{
		AccountId: aws.String("-"),
		ArchiveId: aws.String(c.String("archive-id")),
		VaultName: aws.String(conf.Vault),
	})

	if err != nil {
		return err
	}

	return nil
}
