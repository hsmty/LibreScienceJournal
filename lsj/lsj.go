package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"log"
	"os"
)

var homeDir string
var lsjDir string

func init() {
	homeDir = os.Getenv("HOME")
	lsjDir = fmt.Sprintf("%s/.lsj", homeDir)
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

func createKey() error {
	pubFile := fmt.Sprintf("%s/lsj.pub", lsjDir)
	prvFile := fmt.Sprintf("%s/lsj.key", lsjDir)

	pub, prv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(pubFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	f.Write(pub)
	f.Close()

	f, err = os.OpenFile(prvFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	f.Write(prv)
	f.Close()

	return nil
}

func main() {
	createKeyFlag := flag.Bool("create-key", false,
		fmt.Sprintf("Creates priv/pub key and stores on %s/.lsj/", homeDir))

	flag.Parse()

	if err := createConfigDir(); err != nil {
		log.Fatal("Couldn't open config dir: %v", err)
	}

	if *createKeyFlag == true {
		err := createKey()
		if err != nil {
			log.Fatal(err)
		}
	}
}
