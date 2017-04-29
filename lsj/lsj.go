package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func createKey(force bool) error {
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

func main() {
	createKeyFlag := flag.Bool("create-key", false,
		fmt.Sprintf("Creates priv/pub key and stores on %s/.lsj/", homeDir))

	flag.Parse()

	if err := createConfigDir(); err != nil {
		log.Fatal("Couldn't open config dir: %v", err)
	}

	if *createKeyFlag == true {
		err := createKey(false)
		if err != nil {
			force := false

			if os.IsExist(err) {
				fmt.Print("Keys already exists, do you want to overwrite? (yes/no): ")
				var input string
				fmt.Scanln(&input)

				if strings.TrimRight(input, "\n") == "yes" {
					force = true
					err = nil
				}
			}

			if force == true {
				err = createKey(force)
			}

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
