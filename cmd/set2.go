package cmd

import (
	"fmt"
	"github.com/juggler434/crypto/padding"
	"github.com/spf13/cobra"
)

var input string
var blockLength int

func init() {
	rootCmd.AddCommand(set2Command)
	set2Command.AddCommand(set2Challenge9)

	set2Challenge9.Flags().StringVarP(&input, "input", "", "", "Input to Pad")
	set2Challenge9.Flags().IntVarP(&blockLength, "length", "", 16, "desired block length")

}

var set2Command = &cobra.Command{
	Use:   "set2",
	Short: "Command for running solutions for set 2 of cryptopals",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please run with a problem")
	},
}

var set2Challenge9 = &cobra.Command{
	Use:   "challenge9",
	Short: "runs PKCS7 Padding",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		res := padding.PKCS7([]byte(input), blockLength)
		fmt.Printf("%+q\n", res)
	},
}
