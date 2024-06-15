// cmd/hello.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints Hello, world!",
	Long:  `This command prints Hello, world! to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
