package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	app := cli.NewApp()
	app.Version = "0.0.2"
	app.Name = "gmng"
	app.Before = initial
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "conf",
			Aliases: []string{"c"},
			Value:   os.Getenv("HOME") + "/.gmng.yml",
			Usage:   "path to config file",
		},
		&cli.StringFlag{
			Name:    "vault",
			Aliases: []string{"V"},
			Usage:   "vault name",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "vualts",
			Aliases: []string{"v"},
			Action:  ListVaults,
			Usage:   "list all vaults",
		},
		{
			Name:    "job",
			Aliases: []string{"j"},
			Usage:   "working with jobs",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  ListJobs,
					Usage:   "list all jobs",
				},
				{
					Name:    "inventory",
					Aliases: []string{"i"},
					Action:  InitiateJobInventoryRetrieval,
					Usage:   "initiate inventory-retrieval",
				},
				{
					Name:    "list-archives",
					Aliases: []string{"L"},
					Action:  ListInventoryRetrieval,
					Usage:   "list all archives (inventory-retrieval) in JSON",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "job-id",
							Aliases:  []string{"j", "id"},
							Usage:    "job ID",
							Required: true,
						},
					},
				},
				{
					Name:    "get-archive",
					Aliases: []string{"g"},
					Action:  InitiateArchiveRetrieval,
					Usage:   "initiate archive-retrieval",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "archive-id",
							Aliases:  []string{"a", "id"},
							Usage:    "archive ID",
							Required: true,
						},
						&cli.StringFlag{
							Name:    "descripotion",
							Aliases: []string{"d"},
							Usage:   "descripotion of job",
						},
					},
				},
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete the archive",
			Action:  DeleteArchive,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "archive-id",
					Aliases:  []string{"a", "id"},
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func initial(c *cli.Context) error {
	oFileStat, err := os.Stat(c.String("conf"))
	if err != nil {
		log.Fatalln(err)
	}

	if oFileStat.Mode().String() != "-rw-------" {
		log.Fatalf("set up 0600 permissions for config file %s", c.String("conf"))
	}

	readConfig(c.String("conf"))

	if c.String("vault") != "" {
		conf.Vault = c.String("vault")
	}

	return nil
}
