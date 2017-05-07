package crypto

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"golang.org/x/crypto/ed25519"
	"github.com/hsmty/LibreScienceJournal/common"
)

var (
	pubFileName string
	prvFileName string

	pubKey ed25519.PublicKey
	prvKey ed25519.PrivateKey

	ErrKeysExist = errors.New("The key pair already exists")
	ErrNoKeys    = errors.New("The key pair does not exists")
)

func init() {
	pubFileName = fmt.Sprintf("%s/lsj.pub", common.LsjDir)
	prvFileName = fmt.Sprintf("%s/lsj.key", common.LsjDir)

	pubKey = nil
	prvKey = nil
}

func KeysExists() bool {
	var err error

	if _, err = os.Stat(pubFileName); err == nil {
		return true
	}
	if _, err = os.Stat(prvFileName); err == nil {
		return true
	}

	return false
}

func GetPubKey() (ed25519.PublicKey, error) {
	var err error = nil

	if pubKey == nil {
		pubKey, err = ioutil.ReadFile(pubFileName)
	}

	return pubKey, err
}

func GetPrvKey() (ed25519.PrivateKey, error) {
	var err error = nil

	if prvKey == nil {
		prvKey, err = ioutil.ReadFile(prvFileName)
	}

	return prvKey, err
}

func GetKeys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pubKey, err := GetPubKey()
	if err != nil {
		return nil, nil, err
	}

	prvKey, err := GetPrvKey()
	if err != nil {
		return nil, nil, err
	}

	return pubKey, prvKey, nil
}

func CreateKeys() error {
	pub, prv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}

	pubFile, err := os.OpenFile(pubFileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer pubFile.Close()
	if _, err := pubFile.Write(pub); err != nil {
		return err
	}

	privFile, err := os.OpenFile(prvFileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer privFile.Close()

	if _, err := privFile.Write(prv); err != nil {
		return err
	}

	return nil
}

func SignContent(content []byte) ([]byte, error) {
	prvKey, err := GetPrvKey()
	if err != nil {
		return nil, err
	}

	return ed25519.Sign(prvKey, content), nil
}

func SignFile(file string) ([]byte, []byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}
	contentSing, err := SignContent(content)
	if err != nil {
		return nil, nil, err
	}

	return content, contentSing, nil
}

func SignFiles(files []string) (map[string][][]byte, error) {
	var fileSings = make(map[string][][]byte)

	for _, f := range files {
		content, contentSing, err := SignFile(f)
		if err != nil {
			return nil, err
		}

		fileSings[f] = append(fileSings[f], content)
		fileSings[f] = append(fileSings[f], contentSing)
	}

	return fileSings, nil
}
