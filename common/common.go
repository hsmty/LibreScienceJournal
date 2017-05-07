package common

import (
	"os"
	"fmt"
)

var (
	HomeDir     string
	LsjDir      string
)

func init() {
	HomeDir = os.Getenv("HOME")
	LsjDir = fmt.Sprintf("%s/.lsj", HomeDir)
}

func CreateConfigDir() error {
	if _, err := os.Stat(LsjDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(LsjDir, 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
