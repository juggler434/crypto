package cmd

import (
	"fmt"
	"os"

	cryptopals "github.com/juggler434/crypto/set1"

	"github.com/spf13/cobra"
)

var hexString string

func init() {
	set1Challenge1.Flags().StringVarP(&hexString, "hex", "", "", "Hex string to be base64 encoded")

	rootCmd.AddCommand(set1Command)
	set1Command.AddCommand(set1Challenge1)
}

var set1Command = &cobra.Command{
	Use:   "set1",
	Short: "Command for running solutions for set 1 of cryptopals",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please run with a problem")
	},
}

var set1Challenge1 = &cobra.Command{
	Use:   "challenge1",
	Short: "runs hex to base64 encoding",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := cryptopals.HexToBase64([]byte(hexString))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}
