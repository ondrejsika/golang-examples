package main

import (
	"github.com/sikalabs/gobble/pkg/config"
	"github.com/sikalabs/gobble/pkg/play"
	"github.com/sikalabs/gobble/pkg/run"
	"github.com/sikalabs/gobble/pkg/task"
	"github.com/sikalabs/gobble/pkg/task/lib/print"
	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "gobble_with_cobra",
	Short: "Gobble binary with Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		runGobble(FlagName)
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}

func main() {
	Cmd.Execute()
}

func runGobble(name string) {
	c := config.Config{
		Meta: config.ConfigMeta{
			SchemaVersion: 3,
		},
		Global: config.GlobalConfig{
			NoStrictHostKeyChecking: true,
			Vars: map[string]interface{}{
				"global": "global",
			},
		},
		Hosts: map[string][]config.ConfigHost{
			"all": {
				{
					SSHTarget: "localhost",
					Vars: map[string]interface{}{
						"local": "local",
					},
				},
			},
		},
		Plays: []play.Play{
			{
				Name:  "Hello World",
				Hosts: []string{"all"},
				Tasks: []task.Task{
					{
						Name: "Hello World",
						Print: print.TaskPrint{
							Template: "Hello " + name + " from embedded Gobble -- {{.Vars.global}} {{.Vars.local}}",
						},
					},
				},
			},
		},
	}

	run.Run(c, false, false, []string{}, []string{})
}
