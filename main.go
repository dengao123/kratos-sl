package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/dengao123/kratos-sl/internal/change"
	"github.com/dengao123/kratos-sl/internal/project"
	"github.com/dengao123/kratos-sl/internal/proto"
	"github.com/dengao123/kratos-sl/internal/run"
	"github.com/dengao123/kratos-sl/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "kratos",
	Short:   "Kratos: An elegant toolkit for Go microservices.",
	Long:    `Kratos: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
