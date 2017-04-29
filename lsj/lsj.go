package lsj

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

		err := createKey(force)
		if err != nil {
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
