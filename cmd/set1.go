package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/juggler434/crypto/aes128/ecb"
	"github.com/juggler434/crypto/encoding/base64"
	"github.com/juggler434/crypto/encoding/hex"
	"github.com/juggler434/crypto/xor"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var hexString string
var hexString2 string
var fileName string
var plainText string
var key string

func init() {
	set1Challenge1.Flags().StringVarP(&hexString, "hex", "", "", "Hex string to convert to base64")
	rootCmd.AddCommand(set1Command)
	set1Command.AddCommand(set1Challenge1)
	set1Command.AddCommand(set1Challenge2)
	set1Command.AddCommand(set1Challenge3)
	set1Command.AddCommand(set1Challenge4)
	set1Command.AddCommand(set1Challenge5)
	set1Command.AddCommand(set1Challenge6)
	set1Command.AddCommand(set1Challenge7)
	set1Command.AddCommand(set1Challenge8)

	set1Challenge2.Flags().StringVarP(&hexString, "hex1", "", "", "Hex string to compare")
	set1Challenge2.Flags().StringVarP(&hexString2, "hex2", "", "", "Hex string to compare to")
	set1Challenge3.Flags().StringVarP(&hexString, "hex", "", "", "Hex string to run cipher on")
	set1Challenge4.Flags().StringVarP(&fileName, "file", "", "", "Text file location")
	set1Challenge5.Flags().StringVarP(&plainText, "input", "", "", "Text to encrpyt")
	set1Challenge5.Flags().StringVarP(&key, "key", "", "", "Key for encrypting text")
	set1Challenge6.Flags().StringVarP(&fileName, "file", "", "", "File to be decrypted")
	set1Challenge7.Flags().StringVarP(&fileName, "file", "", "", "File to be decrypted")
	set1Challenge7.Flags().StringVarP(&key, "key", "", "", "Key for encrypted text")
	set1Challenge8.Flags().StringVarP(&fileName, "file", "", "", "File to detect encryption in")

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
		ret, err := hex.ToBase64([]byte(hexString))
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
		input1, err := hex.Decode([]byte(hexString))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input2, err := hex.Decode([]byte(hexString2))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ret, err := xor.Fixed(input1, input2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := hex.Encode(ret)
		fmt.Printf("%s\n", res)
	},
}

var set1Challenge3 = &cobra.Command{
	Use:   "challenge3",
	Short: "performs a single byte xor cipher on hex encoded string",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		inp, err := hex.Decode([]byte(hexString))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ret, _ := xor.SingleCharDecode([]byte(inp))
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge4 = &cobra.Command{
	Use:   "challenge4",
	Short: "finds which string has been single xor encoded in a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var input [][]byte
		f, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lns := strings.Split(string(f), "\n")
		for _, ln := range lns {
			dhl, err := hex.Decode([]byte(ln))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			input = append(input, dhl)
		}
		ret := xor.DetectSingleCharEncryption(input)
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge5 = &cobra.Command{
	Use:   "challenge5",
	Short: "Repeting Xor encrypts a string using the given key.  Returns a hex encoded string",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(plainText)
		ret := xor.EncryptWithRepeatingKey([]byte(plainText), []byte(key))

		ret = hex.Encode(ret)
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge6 = &cobra.Command{
	Use:   "challenge6",
	Short: "Decrypt a Base64 encoded repeating XOR encoded file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ueb, err := base64.Decode(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ueb = bytes.Trim(ueb, "\x00")
		ret := xor.Decrypt(ueb)
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge7 = &cobra.Command{
	Use:   "challenge7",
	Short: "Decrypt and AES ECB 128 encrypted file that has been base64 encoded",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ueb, err := base64.Decode(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ret, err := ecb.Decrypt(ueb, []byte(key))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}

var set1Challenge8 = &cobra.Command{
	Use:   "challenge8",
	Short: "Find AES encrypted line in a file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var input [][]byte
		f, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(bytes.NewReader(f))
		for scanner.Scan() {
			hdl, err := hex.Decode(scanner.Bytes())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			input = append(input, hdl)
		}
		res := ecb.Detect(input)
		fmt.Printf("%d", res)
	},
}
