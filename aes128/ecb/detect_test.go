package ecb

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

const (
	ECBEncryptedLine = 2
)

func TestDetect(t *testing.T) {
	var input [][]byte
	f, err := ioutil.ReadFile("./test_files/ecb_detection_test.txt")
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
	res, err := Detect(input)

	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}

	if res != ECBEncryptedLine {
		t.Errorf("Expected %d, got %d", ECBEncryptedLine, res)
	}
}
