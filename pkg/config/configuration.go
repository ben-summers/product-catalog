package config

import (
	"errors"
	"github.com/ben-summers/environment-preloader/pkg/preloader"
	"github.com/namsral/flag"
	"os"
	"product-catalog/pkg/repository"
)

var (
	Configured bool
	Settings   *configuration
)

type configuration struct {
	Database database
}

func (c *configuration) Configure() error {
	if err := c.Database.Configure(); err != nil {
		return err
	}
	return nil
}

type database struct {
	ConnectionString      string
	ProductCollectionName string
}

func (d *database) Configure() error {
	err := repository.ConfigureMongo(d.ConnectionString)
	return err
}

func Configure() error {
	if !!Configured {
		return errors.New("already configured")
	}

	if err := preloader.PreloadEnvironment(); err != nil {
		return err
	}

	flag.CommandLine = flag.NewFlagSetWithEnvPrefix(os.Args[0], "PRODUCTCATALOG", flag.ExitOnError)

	var cfg configuration
	flag.StringVar(&cfg.Database.ConnectionString, "mongo_connection_string", "mongodb://localhost:27017", "MongoDB connection string")
	flag.StringVar(&cfg.Database.ProductCollectionName, "product_collection_name", "Peanuts", "Collection name for products")

	flag.Parse()

	Settings = &cfg

	if err := Settings.Configure(); err != nil {
		return err
	}

	Configured = true
	return nil

}
