package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var err error

	createKeyFlag := flag.Bool("create-key", false,
		fmt.Sprintf("Creates priv/pub key and stores on %s/.lsj/", homeDir))
	publishFlag := flag.String("publish", "", fmt.Sprintf("Publish article"))
	flag.Parse()

	if err = createConfigDir(); err != nil {
		log.Fatal("Couldn't open config dir: %v", err)
	}

	if *createKeyFlag == true {
		force := false

		err = CreateKey(force)
		if err != nil {
			if err == KeysExist {
				input := AskUserInput("Keys already exists, do you want to overwrite? (yes/no): ")
				if input == "yes" {
					force = true
					err = nil
				}
			}

			if force == true {
				err = CreateKey(force)
			}


		}
	} else if *publishFlag != "" {
		input := AskUserInput("You are goning to sign and send article, are you sure? (yes/no): ")
		if input == "yes" {
			err = PublishArticle(*publishFlag)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
