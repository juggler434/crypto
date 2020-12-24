package cmd

import (
	"fmt"
	"github.com/juggler434/crypto/account"
	"github.com/juggler434/crypto/aes128"
	"github.com/juggler434/crypto/aes128/cbc"
	"github.com/juggler434/crypto/aes128/oracle"
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
	set2Command.AddCommand(set2Challenge12)
	set2Command.AddCommand(set2Challenge13)
	set2Command.AddCommand(set2Challenge14)
	set2Command.AddCommand(set2Challenge16)

	set2Challenge9.Flags().StringVarP(&input, "input", "", "", "Input to Pad")
	set2Challenge9.Flags().IntVarP(&blockLength, "length", "", 16, "desired block length")
	set2Challenge10.Flags().StringVarP(&file, "file", "", "", "Path to file")
	set2Challenge10.Flags().StringVarP(&key, "key", "", "", "key to encrypt/decrypt with")
	set2Challenge10.Flags().StringVarP(&initializationVector, "iv", "", "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", "initialization vector to start for encryption")
	set2Challenge11.Flags().StringVarP(&file, "file", "", "", "file to encrypt and detect")
	set2Challenge12.Flags().StringVarP(&input, "input", "", "", "Input to run oracle on")
	set2Challenge13.Flags().StringVarP(&input, "input", "", "", "email trying to access admin panel")
	set2Challenge14.Flags().StringVarP(&input, "input", "", "", "bit that is going to be encrypted (or if you're in the know, will get you a secret")
	set2Challenge16.Flags().StringVarP(&input, "input", "", "", "post a comment (or hack our very secure system)")
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
	Short: "runs Pad Padding",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		res := pkcs7.Pad([]byte(input), blockLength)
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

		ro := oracle.NewRandom()
		encryptedText, encryptionPattern, err := ro.Encrypt(input)
		if err != nil {
			fmt.Printf("failed to encrypt file: %s", err)
		}

		fmt.Printf("Encrypted file using: %d\n", encryptionPattern)
		dp := aes128.DetectMode(encryptedText)
		fmt.Printf("DetectMode detected: %d\n", dp)
	},
}

var set2Challenge12 = &cobra.Command{
	Use:   "challenge12",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		di, err := base64.Decode([]byte(input))
		if err != nil {
			fmt.Printf("failed to base64 decode input: %s", err)
			os.Exit(1)
		}

		server := oracle.NewECBOracle(di)
		ret, err := aes128.BreakECBSimple(server)

		if err != nil {
			fmt.Printf("failed to break ECS encryption: %s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", ret)
	},
}

var set2Challenge13 = &cobra.Command{
	Use:   "challenge13",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking email to see if allowed access to admin panel")
		us := account.NewUserService()
		var ret []byte
		var err error
		if input == "hack the planet" {
			ret = account.Attack(us)
		} else {
			ret, err = us.GetUser([]byte(input))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		allowed, err := us.CheckAdminPermission(ret)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !allowed {
			fmt.Println("Access denied")
		} else {
			fmt.Println("Here are all the things!")
		}

	},
}

var set2Challenge14 = &cobra.Command{
	Use:   "challenge14",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		encrypter := oracle.NewAdvancedECBOracle([]byte("SUPER SECRET API KEY"))
		var ret []byte
		var err error
		if input == "hack the planet" {
			ret, err = aes128.BreakECBAdvanced(encrypter)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			eb, err := encrypter.Encrypt([]byte(input))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ret = base64.Encode(eb)
		}

		fmt.Printf("%s\n", ret)
	},
}

var set2Challenge16 = &cobra.Command{
	Use:   "challenge16",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Posting comment")
		bfe := oracle.NewCommentEncrypter()
		var aa bool
		var ret []byte
		var err error
		if input == "hack the planet" {
			aa, err = aes128.BitflipAttack(bfe)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			res, err := bfe.Encrypt([]byte(input))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			aa, err = bfe.CheckIsAdmin(res)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			ret = base64.Encode(res)
		}
		if aa {
			fmt.Println("Welcome Administrator, let's do some damage")
		} else {
			fmt.Println("Thank you for your comment user")
			fmt.Printf("%s\n", ret)
		}

	},
}
