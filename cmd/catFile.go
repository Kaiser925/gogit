// Package cmd /*
package cmd

import (
	"fmt"

	"github.com/Kaiser925/gogit/internal/pkg/object"
	"github.com/Kaiser925/gogit/internal/pkg/output"
	"github.com/Kaiser925/gogit/internal/pkg/repository"
	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type and size information for repository objects",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := args[0]
		sha := args[1]
		repo, err := repository.New(".")
		if err != nil {
			output.Fatal(err.Error())
		}
		p, err := repo.Open(sha)
		if err != nil {
			output.Fatal(err.Error())
		}
		o, err := object.Parse(t, p)
		if err != nil {
			output.Fatal(err.Error())
		}
		fmt.Print(o.String())
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}
