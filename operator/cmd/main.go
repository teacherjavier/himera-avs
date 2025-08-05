package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/imua-xyz/imua-avs/core/config"
	"github.com/imua-xyz/imua-avs/operator"
	"github.com/imua-xyz/imua-avs/types"

	sdkutils "github.com/imua-xyz/imua-avs-sdk/utils"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Config path is required")
	}

	config, err := types.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}

	op, err := operator.NewOperatorFromConfig(config)
	if err != nil {
		log.Fatalf("Cannot create operator: %s", err)
	}

	log.Println("HIMERA Operator starting...")
	err = op.Start(context.Background())
	if err != nil {
		log.Fatalf("Operator stopped with error: %s", err)
	}
}
