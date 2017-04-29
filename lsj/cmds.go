package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/ed25519"
)

var homeDir string
var lsjDir string
var pubFileName string
var prvFileName string

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

func keysExists() error {
	if _, err := os.Stat(pubFileName); err == nil {
		return os.ErrExist
	}
	if _, err := os.Stat(prvFileName); err == nil {
		return os.ErrExist
	}

	return nil
}

func CreateKey(force bool) error {
	if force == false {
		if err := keysExists(); err != nil {
			return err
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
