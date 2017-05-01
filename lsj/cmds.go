package main

import (
	"fmt"
	"os"
	"errors"
	"io/ioutil"
	"golang.org/x/crypto/ed25519"
)

var (
	homeDir string
	lsjDir string
	pubFileName string
	prvFileName string

	KeysExist = errors.New("keys already exists")
	KeysNotExist = errors.New("keys does not exists")
)

func init() {
	homeDir = os.Getenv("HOME")
	lsjDir = fmt.Sprintf("%s/.lsj", homeDir)
	pubFileName = fmt.Sprintf("%s/lsj.pub", lsjDir)
	prvFileName = fmt.Sprintf("%s/lsj.key", lsjDir)
}

func createConfigDir() error {
	if _, err := os.Stat(lsjDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(lsjDir, 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func keysExists() bool {
	var err error

	if _, err = os.Stat(pubFileName); err == nil {
		return true
	}
	if _, err = os.Stat(prvFileName); err == nil {
		return true
	}

	return false
}

func getKeys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pubKey, err := ioutil.ReadFile(pubFileName)
	if err != nil {
		return nil, nil, err
	}

	prvKey, err := ioutil.ReadFile(prvFileName)
	if err != nil {
		return nil, nil, err
	}

	return pubKey, prvKey, nil
}

func CreateKey(force bool) error {
	if force == false {
		if keysExists() == true {
			return KeysExist
		}
	}

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

func PublishArticle(files... string) error {
	var err error

	if keysExists() == false {
		return KeysNotExist
	}

	for _, s := range files {
		if _, err = os.Stat(s); err != nil {
			return os.ErrNotExist
		}
	}

	_, prvKey, err := getKeys()
	if err != nil {
		return err
	}

	var m = make(map[string][]byte)
	for _, f := range files {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		m[f] = ed25519.Sign(prvKey, content)
	}

	return nil
}
