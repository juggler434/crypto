package xor

import (
	"fmt"
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)


const TestFile = "./test_files/xor_test.txt"
const TestString = "test case"
func TestDetectSingleCharEncryption(t *testing.T) {
	t.Run("with valid params", func(t *testing.T) {
		var input [][]byte
		f, err := ioutil.ReadFile(TestFile)
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
		res := DetectSingleCharEncryption(input)

		if string(res) != TestString {
			t.Errorf("Exptected: %s, got: %s", TestString, res)
		}
	})
}