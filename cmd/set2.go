package cmd

import (
	"fmt"
	"github.com/juggler434/crypto/aes128"
	"github.com/juggler434/crypto/aes128/cbc"
	"github.com/juggler434/crypto/encoding/base64"
	"github.com/juggler434/crypto/padding"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var input string
var blockLength int
var file string
var initializationVector string

func init() {
	rootCmd.AddCommand(set2Command)
	set2Command.AddCommand(set2Challenge9)
	set2Command.AddCommand(set2Challenge10)
	set2Command.AddCommand(set2Challenge11)

	set2Challenge9.Flags().StringVarP(&input, "input", "", "", "Input to Pad")
	set2Challenge9.Flags().IntVarP(&blockLength, "length", "", 16, "desired block length")
	set2Challenge10.Flags().StringVarP(&file, "file", "", "", "Path to file")
	set2Challenge10.Flags().StringVarP(&key, "key", "", "", "key to encrypt/decrypt with")
	set2Challenge10.Flags().StringVarP(&initializationVector, "iv", "", "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "initialization vector to start for encryption")
	set2Challenge11.Flags().StringVarP(&file, "file", "", "", "file to encrypt and detect")

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

var set2Challenge10 = &cobra.Command{
	Use:   "challenge10",
	Short: "runs cbc encryption or decryption",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("must be run with either encrypt or decrypt")
			os.Exit(1)
		}

		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("failed to read file contents: %s", err)
		}

		switch action := args[0]; action {
		case "encrypt":
			res, err := cbc.Encrypt(input, []byte(key), []byte(initializationVector))
			if err != nil {
				fmt.Printf("failed to encrypt input: %s", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", base64.Encode(res))
		case "decrypt":
			b64input, err := base64.Decode(input)
			if err != nil {
				fmt.Printf("failed to decode base 64 input: %s", err)
				os.Exit(1)
			}
			res, err := cbc.Decrypt(b64input, []byte(key), []byte(initializationVector))
			if err != nil {
				fmt.Printf("failed to decrypt input: %s", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", res)
		}
	},
}

var set2Challenge11 = &cobra.Command{
	Use:   "challenge11",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("failed to read file: %s", err)
			os.Exit(1)
		}

		encryptedText, encryptionPattern, err := aes128.RandomEncrypt(input)
		if err != nil {
			fmt.Printf("failed to encrypt file: %s", err)
		}

		fmt.Printf("Encrypted file using: %d\n", encryptionPattern)
		dp := aes128.DetectMode(encryptedText)
		fmt.Printf("DetectMode detected: %d\n", dp)
	},
}
