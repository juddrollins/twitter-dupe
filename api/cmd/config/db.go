package config

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// ENV variables
type (
	Env      string
	Database struct {
		TableName string
		*dynamodb.Client
	}
	Config struct {
		Env Env
		Aws aws.Config
		Ddb Database
	}
)

const (
	// Envs
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"

	// Env variables
	EnvEnvVar      = "ENV"
	DDBTableEnvVar = "DDB_TABLE_NAME"
)

type Entry struct {
	PK   string `json:"PK"`
	SK   string `json:"SK"`
	Data string `json:"Data"`
}

type EntryService interface {
	CreateRecord(m Entry) (Entry, error)
	GetRecord(PK string, SK string) (Entry, error)
}

func New() Config {
	env := Env(os.Getenv(EnvEnvVar))
	awsCfg, err := getAwsConfig(env)
	if err != nil {
		panic("failed to create config due to error: " + err.Error())
	}

	return Config{
		Env: env,
		Aws: awsCfg,
		Ddb: getDdbClient(awsCfg),
	}
}

func getDdbClient(cfg aws.Config) Database {
	tableName := os.Getenv(DDBTableEnvVar)

	return Database{
		TableName: tableName,
		Client:    dynamodb.NewFromConfig(cfg),
	}
}

func getAwsConfig(env Env) (aws.Config, error) {
	switch env {
	case Local:
		return config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("default"))
	case Dev:
		return config.LoadDefaultConfig(context.Background())
	case Prod:
		return config.LoadDefaultConfig(context.Background())
	default:
		return aws.Config{}, errors.New(`"ENV" variable was not set, valid options are: [local, dev, prod]`)
	}
}
