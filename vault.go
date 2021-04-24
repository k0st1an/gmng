package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glacier"
	cli "github.com/urfave/cli/v2"
)

// ListVaults ..
func ListVaults(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())

	result, err := glr.ListVaults(&glacier.ListVaultsInput{AccountId: aws.String("-")})
	if err != nil {
		return err
	}

	return printJSON(result.VaultList)
}
