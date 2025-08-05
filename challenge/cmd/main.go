package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/urfave/cli"

	"github.com/imua-xyz/imua-avs/challenge"
	"github.com/imua-xyz/imua-avs-sdk/utils"
	"github.com/imua-xyz/imua-avs/types"
)

func main() {
	app := cli.NewApp()
	app.Name = "himera-challenger"
	app.Usage = "HIMERA AVS Challenger Service"
	app.Description = "A service that monitors and resolves challenges for the HIMERA AVS."
	
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Path to the configuration file",
			Value: "config.yaml",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Starts the challenger as a long-running service",
			Action: func(c *cli.Context) error {
				nodeConfig, err := types.ReadConfig(c.GlobalString("config"))
				if err != nil { return err }
				challenger, err := challenge.NewChallengerFromConfig(nodeConfig)
				if err != nil { return err }
				log.Println("Challenger starting in service mode...")
				return challenger.Start(context.Background())
			},
		},
		{
			Name:  "resolve",
			Usage: "Manually resolve a challenge for a single task",
			Flags: []cli.Flag{
				cli.Uint64Flag{Name: "task-id", Usage: "The ID of the task to challenge", Required: true},
			},
			Action: func(c *cli.Context) error {
				taskID := c.Uint64("task-id")
				nodeConfig, err := types.ReadConfig(c.GlobalString("config"))
				if err != nil { return err }
				challenger, err := challenge.NewChallengerFromConfig(nodeConfig)
				if err != nil { return err }

                taskInfo, err := challenger.avsReader.GetTaskInfo(&bind.CallOpts{}, challenger.avsAddr.String(), taskID)
				if err != nil { return fmt.Errorf("cannot get task info for task %d: %w", taskID, err) }
                
                himeraTaskDefId, ok := challenger.taskDefHashMap[taskInfo.Hash]
                if !ok { return fmt.Errorf("could not find task definition for hash %s", taskInfo.Hash.Hex()) }
                
				log.Printf("Manually resolving challenge for task ID: %d\n", taskID)
				return challenger.ResolveChallenge(context.Background(), taskID, himeraTaskDefId, taskInfo.TaskInput)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Application failed: %s", err)
	}
}
