package oracle

import (
	"bytes"
	"crypto/rand"
	"github.com/juggler434/crypto/aes128/cbc"
	"github.com/juggler434/crypto/encoding/base64"
	"io/ioutil"
	"testing"
)

const cbcTestFile = "./test_files/cbc_oracle_test.txt"

func readTestFile(t *testing.T) [][]byte {
	f, err := ioutil.ReadFile(cbcTestFile)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	l := bytes.Split(f, []byte("\n"))
	return l
}

func TestNewCBCOracle(t *testing.T) {
	l := readTestFile(t)

	r, err := NewCBCOracle(l)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	v := r.(*CBCOracle)
	if v.Key == nil {
		t.Error("expected a key to be initialized, got nil")
	}

	if v.Plaintext == nil || len(v.Plaintext) == 0 {
		t.Error("expected plaintext lines, got nil")
	}
}

func TestCBCOracle_Encrypt(t *testing.T) {
	l := readTestFile(t)
	r, err := NewCBCOracle(l)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	o := r.(*CBCOracle)

	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	res, err := r.Encrypt(iv)
	if err != nil {
		t.Errorf("unexptected error: %s", err)
	}

	v, _ := base64.Decode(res)

	riv := v[:16]
	if !bytes.Equal(iv, riv) {
		t.Errorf("expected first 16 bytes to equal: %s, got %s", iv, riv)
	}

	e := v[17:]

	_, err = cbc.Decrypt(e, o.Key, iv)
	if err != nil {
		t.Errorf("expected nil error, got: %s", err)
	}
}
