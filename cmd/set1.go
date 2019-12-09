package cmd

import (
	"fmt"
	"os"

	cryptopals "github.com/juggler434/crypto/set1"

	"github.com/spf13/cobra"
)

var hexString string
var hexString2 string
var fileName string

func init() {
	set1Challenge1.Flags().StringVarP(&hexString, "hex", "", "", "Hex string to convert to base64")
	rootCmd.AddCommand(set1Command)
	set1Command.AddCommand(set1Challenge1)
	set1Command.AddCommand(set1Challenge2)
	set1Command.AddCommand(set1Challenge3)
	set1Command.AddCommand(set1Challenge4)

	set1Challenge2.Flags().StringVarP(&hexString, "hex1", "", "", "Hex string to compare")
	set1Challenge2.Flags().StringVarP(&hexString2, "hex2", "", "", "Hex string to compare to")
	set1Challenge3.Flags().StringVarP(&hexString, "hex", "", "", "Hex string to run cipher on")
	set1Challenge4.Flags().StringVarP(&fileName, "file", "", "", "Text file location")

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

var set1Challenge2 = &cobra.Command{
	Use:   "challenge2",
	Short: "performs fixed xor comparison on two hex encoded strings",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := cryptopals.FixedXor([]byte(hexString), []byte(hexString2))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge3 = &cobra.Command{
	Use:   "challenge3",
	Short: "performs a single byte xor cipher on hex encoded string",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ret, _, err := cryptopals.SingleXorCipher([]byte(hexString))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge4 = &cobra.Command{
	Use: "challenge4",
	Short: "finds which string has been single xor encoded in a file",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := cryptopals.FindXorCipherString(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}
