package main

import (
	"os"
	"flag"
	"fmt"
	"log"
	"golang.org/x/crypto/ed25519"
)

var homeDir string
var lsjDir string

func init() {
	homeDir = os.Getenv("HOME")
	lsjDir = fmt.Sprintf("%s/.lsj", homeDir)

	// XXX: Test for path exists
	os.Mkdir(lsjDir, 0700)
}

func createKey() int {
	// XXX: Test if key/pub files exists
	pubFile := fmt.Sprintf("%s/lsj.pub", lsjDir)
	prvFile := fmt.Sprintf("%s/lsj.key", lsjDir)

	pub, prv, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(pubFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(pub)
	f.Close()

	f, err = os.OpenFile(prvFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(prv)
	f.Close()

	return 0
}

func _main() int {
	createKeyFlag := flag.Bool("create-key", false,
		fmt.Sprintf("Creates priv/pub key and stores on %s/.lsj/", homeDir))

	flag.Parse()

	if *createKeyFlag == true {
		return createKey()
	}

	return 0
}

func main() {
	os.Exit(_main())
}
