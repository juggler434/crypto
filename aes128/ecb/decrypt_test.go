package ecb

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name           string
		file           string
		key            []byte
		expectedReturn []byte
		shouldError    bool
	}{
		{
			name:           "base case",
			file:           "./test_files/aes_decrypt_test.txt",
			key:            []byte("YELLOW SUBMARINE"),
			expectedReturn: []byte(ExpectedAESReturn),
			shouldError:    false,
		},
		{
			name:           "invalid key",
			file:           "./test_files/aes_decrypt_test.txt",
			key:            []byte("YELLO"),
			expectedReturn: nil,
			shouldError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			f, err := ioutil.ReadFile(test.file)
			if err != nil {
				t.Errorf("error should be nil, got: %s", err)
			}
			ueb := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
			_, err = base64.StdEncoding.Decode(ueb, f)
			if err != nil {
				t.Errorf("error should be nil, got: %s", err)
			}
			res, err := Decrypt(ueb, test.key)
			if test.shouldError {
				if err == nil {
					t.Errorf("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			//for i, b := range test.expectedReturn {
			//	if res[i] != b {
			//		fmt.Printf("The res char is: %d. the expected char is :%d, the index is: %d", b, res[i], i)
			//		break
			//	}
			//}

			if !bytes.Equal(res, test.expectedReturn) {
				t.Errorf("Expected: %+q, got: %+q", ExpectedAESReturn, res)
			}
		})
	}

	//t.Run("with valid input", func(t *testing.T) {
	//
	//	f, err := ioutil.ReadFile("./test_files/aes_decrypt_test.txt")
	//	if err != nil {
	//		t.Errorf("error should be nil, got: %s", err)
	//	}
	//	ueb := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
	//	_, err = base64.StdEncoding.Decode(ueb, f)
	//	if err != nil {
	//		t.Errorf("error should be nil, got: %s", err)
	//	}
	//	res, err := Decrypt(ueb, []byte("YELLOW SUBMARINE"))
	//	if err != nil {
	//		t.Errorf("Expected: error to be nil, got: %s", err)
	//	}
	//
	//	for i, b := range []byte(ExpectedAESReturn) {
	//		if res[i] != b {
	//			fmt.Printf("The res char is: %d. the expected char is :%d, the index is: %d", b, res[i], i)
	//			break
	//		}
	//	}
	//
	//	if string(res) != ExpectedAESReturn {
	//		t.Errorf("Expected: %s, got: %s", ExpectedAESReturn, res)
	//	}
	//})
}

const ExpectedAESReturn = `I'm back and I'm ringin' the bell 
A rockin' on the mike while the fly girls yell 
In ecstasy in the back of me 
Well that's my DJ Deshay cuttin' all them Z's 
Hittin' hard and the girlies goin' crazy 
Vanilla's on the mike, man I'm not lazy. 

I'm lettin' my drug kick in 
It controls my mouth and I begin 
To just let it flow, let my concepts go 
My posse's to the side yellin', Go Vanilla Go! 

Smooth 'cause that's the way I will be 
And if you don't give a damn, then 
Why you starin' at me 
So get off 'cause I control the stage 
There's no dissin' allowed 
I'm in my own phase 
The girlies sa y they love me and that is ok 
And I can dance better than any kid n' play 

Stage 2 -- Yea the one ya' wanna listen to 
It's off my head so let the beat play through 
So I can funk it up and make it sound good 
1-2-3 Yo -- Knock on some wood 
For good luck, I like my rhymes atrocious 
Supercalafragilisticexpialidocious 
I'm an effect and that you can bet 
I can take a fly girl and make her wet. 

I'm like Samson -- Samson to Delilah 
There's no denyin', You can try to hang 
But you'll keep tryin' to get my style 
Over and over, practice makes perfect 
But not if you're a loafer. 

You'll get nowhere, no place, no time, no girls 
Soon -- Oh my God, homebody, you probably eat 
Spaghetti with a spoon! Come on and say it! 

VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino 
Intoxicating so you stagger like a wino 
So punks stop trying and girl stop cryin' 
Vanilla Ice is sellin' and you people are buyin' 
'Cause why the freaks are jockin' like Crazy Glue 
Movin' and groovin' trying to sing along 
All through the ghetto groovin' this here song 
Now you're amazed by the VIP posse. 

Steppin' so hard like a German Nazi 
Startled by the bases hittin' ground 
There's no trippin' on mine, I'm just gettin' down 
Sparkamatic, I'm hangin' tight like a fanatic 
You trapped me once and I thought that 
You might have it 
So step down and lend me your ear 
'89 in my time! You, '90 is my year. 

You're weakenin' fast, YO! and I can tell it 
Your body's gettin' hot, so, so I can smell it 
So don't be mad and don't be sad 
'Cause the lyrics belong to ICE, You can call me Dad 
You're pitchin' a fit, so step back and endure 
Let the witch doctor, Ice, do the dance to cure 
So come up close and don't be square 
You wanna battle me -- Anytime, anywhere 

You thought that I was weak, Boy, you're dead wrong 
So come on, everybody and sing this song 

Say -- Play that funky music Say, go white boy, go white boy go 
play that funky music Go white boy, go white boy, go 
Lay down and boogie and play that funky music till you die. 

Play that funky music Come on, Come on, let me hear 
Play that funky music white boy you say it, say it 
Play that funky music A little louder now 
Play that funky music, white boy Come on, Come on, Come on 
Play that funky music 
`
