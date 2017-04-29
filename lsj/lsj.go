package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	createKeyFlag := flag.Bool("create-key", false,
		fmt.Sprintf("Creates priv/pub key and stores on %s/.lsj/", homeDir))
	flag.Parse()

	if err := createConfigDir(); err != nil {
		log.Fatal("Couldn't open config dir: %v", err)
	}

	if *createKeyFlag == true {
		force := false

		err := CreateKey(force)
		if err != nil {
			if os.IsExist(err) {
				input := AskUserInput("Keys already exists, do you want to overwrite? (yes/no): ")
				if input == "yes" {
					force = true
					err = nil
				}
			}

			if force == true {
				err = CreateKey(force)
			}

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
