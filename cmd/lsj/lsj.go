package main

import (
	"fmt"
	"log"
	"os"
	"github.com/hsmty/LibreScienceJournal/common"
	"github.com/hsmty/LibreScienceJournal/crypto"
)

var (
	Version = "Proof of Concept"
)

func usage() {
	fmt.Println("Usage: lsj <command> <options>")
	fmt.Println("Publish and gets scientific articles")
	fmt.Println("Available commands:")
	fmt.Println("  create-keys - Create a new key for signing")
	fmt.Println("  publish <document> [<attachments> ...] - Publish the document to the net")
	fmt.Println("  search <term1> [<term2> ...] [tag:<tag>] - Search for documents")
	fmt.Println("  fetch <uuid>")
	fmt.Println("Version: ", Version)
}

func init() {
	err := common.CreateConfigDir()
	if err != nil {
		log.Fatal("Can't create config directory", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "create-keys":
		err := CreateKeys(false)
		if err == crypto.ErrKeysExist {
			input := common.AskUserInput("Keys already exists, do you want to overwrite? [yes/No] ")
			if input == "yes" || input == "y" {
				err := CreateKeys(true)
				if err != nil {
					log.Fatal("An error ocurred while creating the keys: ", err)
				}
			}
		}
	case "publish":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		article := os.Args[2]

		var atts []string = nil
		if len(os.Args) > 3 {
			atts = os.Args[3:]
		}

		err := PublishArticle(common.Config.Server, article, atts)
		if err != nil {
			log.Fatal("An error ocurred while publishing the article: ", err)
		}
	case "search":
		fmt.Println("searching...")
	case "fetch":
		fmt.Println("fetching")
	default:
		usage()
	}
}
