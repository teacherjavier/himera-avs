package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/imua-xyz/imua-avs-sdk/utils"
	
    // --- RUTA DE IMPORTACIÓN CORREGIDA ---
	"github.com/imua-xyz/imua-avs/avs"
	"github.com/imua-xyz/imua-avs/types"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	watchDir := flag.String("watch-dir", "./tasks/new", "Directory to watch for new tasks")
	processedDir := flag.String("processed-dir", "./tasks/processed", "Directory to move processed tasks to")
	interval := flag.Int("interval", 10, "Interval in seconds to check for new tasks")
	flag.Parse()

	if *configPath == "" { log.Fatal("Config path is required") }

	os.MkdirAll(*watchDir, os.ModePerm)
	os.MkdirAll(*processedDir, os.ModePerm)

	nodeConfig := types.NodeConfig{}
	if err := sdkutils.ReadYamlConfig(*configPath, &nodeConfig); err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}

	avsService, err := avs.NewAvsService(nodeConfig)
	if err != nil {
		log.Fatalf("Cannot create AVS service: %s", err)
	}

	err = avsService.Start(context.Background(), *watchDir, *processedDir, time.Duration(*interval)*time.Second)
	if err != nil {
		log.Fatalf("AVS service stopped with error: %s", err)
	}
}
