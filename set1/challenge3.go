package cryptopals

//SingleXorCiphter takes an encoded message and decodes it into a human readable string
func SingleXorCipher(codedMessage []byte) []byte, error{
	//Input: Byte Slice (Hex encoded, xor'd)
	//Output: Byte Slice (Human readable)

	// Steps:
	// 1. Decode slice from hex
	// 2. XOR the decoded slice against a character
	// 3. Loop throught the result, scoring based on character frequency
	// 4. If score is higher than previous best score, save result
	// 5. Return
}