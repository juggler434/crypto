package xor

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetHammingDistance(t *testing.T) {
	s1 := []byte("this is a test")
	s2 := []byte("wokka wokka!!!")
	expected := 37

	r := getHammingDistance(s1, s2)
	if r != expected {
		t.Errorf("Expected: %d, got: %d", expected, r)
	}
}

func TestDecrypt(t *testing.T) {
	t.Run("with valid input", func(t *testing.T) {
		f, err := ioutil.ReadFile(BreakXorTestFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ueb := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
		_, err = base64.StdEncoding.Decode(ueb, f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ueb = bytes.Trim(ueb, "\x00")

		r, err := Decrypt(ueb)
		if err != nil {
			t.Errorf("Expected err to be nil, got: %s", err)
		}
		if string(r) != BreakXorOutput {
			t.Errorf("Expected: %s, got: %s", BreakXorOutput, r)
		}
	})
}

const BreakXorOutput = `Ah, look at all the lonely people
Ah, look at all the lonely people

Eleanor Rigby picks up the rice in the church where a wedding has been
Lives in a dream
Waits at the window, wearing the face that she keeps in a jar by the door
Who is it for?

All the lonely people
Where do they all come from?
All the lonely people
Where do they all belong?

Father McKenzie writing the words of a sermon that no one will hear
No one comes near
Look at him working, darning his socks in the night when there's nobody there
What does he care?

All the lonely people
Where do they all come from?
All the lonely people
Where do they all belong?

Ah, look at all the lonely people
Ah, look at all the lonely people

Eleanor Rigby died in the church and was buried along with her name
Nobody came
Father McKenzie wiping the dirt from his hands as he walks from the grave
No one was saved

All the lonely people
(Ah, look at all the lonely people)
Where do they all come from?
All the lonely people
(Ah, look at all the lonely people)
Where do they all belong?`

const BreakXorKey = "JUDE" //This isn't used, I included it for reference
const BreakXorTestFile = "./test_files/xor_decrypt_test.txt"