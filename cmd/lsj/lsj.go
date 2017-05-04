package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const defaultServer string = "127.0.0.1:8080"

func main() {
	var err error

	serverFlag := flag.String("server", defaultServer, fmt.Sprintf(
		"Server to publish"))
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
			files := strings.Fields(*publishFlag)
			article := files[0]
			atts :=files[1:]
			err = PublishArticle(*serverFlag, article, atts)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
