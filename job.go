package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glacier"
	cli "github.com/urfave/cli/v2"
)

// ListJobs ..
func ListJobs(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())

	result, err := glr.ListJobs(&glacier.ListJobsInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(conf.Vault),
	})
	if err != nil {
		return err
	}

	return printJSON(result.JobList)
}

// InitiateJobInventoryRetrieval ..
func InitiateJobInventoryRetrieval(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())

	result, err := glr.InitiateJob(&glacier.InitiateJobInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(conf.Vault),
		JobParameters: &glacier.JobParameters{
			Type: aws.String("inventory-retrieval"),
		},
	})
	if err != nil {
		return err
	}

	return printJSON(result)
}

// ListInventoryRetrieval ..
func ListInventoryRetrieval(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())

	result, err := glr.GetJobOutput(&glacier.GetJobOutputInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(conf.Vault),
		JobId:     aws.String(c.String("job-id")),
	})
	if err != nil {
		return err
	}
	defer result.Body.Close()

	var bodyBytes []byte

	if *result.Status == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(result.Body)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("HTTP status != 200")
	}

	fmt.Println(string(bodyBytes))
	return nil
}

// InitiateArchiveRetrieval ..
func InitiateArchiveRetrieval(c *cli.Context) error {
	glr := glacier.New(newGlacierSession())
	result, err := glr.InitiateJob(&glacier.InitiateJobInput{
		AccountId: aws.String("-"),
		VaultName: aws.String(conf.Vault),
		JobParameters: &glacier.JobParameters{
			Type:        aws.String("archive-retrieval"),
			ArchiveId:   aws.String(c.String("archive-id")),
			Description: aws.String(c.String("descripotion")),
		},
	})
	if err != nil {
		return err
	}

	return printJSON(result)
}
