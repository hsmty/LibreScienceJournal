package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	homeDir     string
	lsjDir      string
	pubFileName string
	prvFileName string

	ErrKeysExist = errors.New("The key pair already exists")
	ErrNoKeys    = errors.New("The key pair does not exists")
)

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

func keysExists() bool {
	var err error

	if _, err = os.Stat(pubFileName); err == nil {
		return true
	}
	if _, err = os.Stat(prvFileName); err == nil {
		return true
	}

	return false
}

func getKeys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pubKey, err := ioutil.ReadFile(pubFileName)
	if err != nil {
		return nil, nil, err
	}

	prvKey, err := ioutil.ReadFile(prvFileName)
	if err != nil {
		return nil, nil, err
	}

	return pubKey, prvKey, nil
}

func CreateKey(force bool) error {
	if force == false {
		if keysExists() == true {
			return ErrKeysExist
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

func PublishArticle(server string, article string, atts []string) error {
	var err error

	if keysExists() == false {
		return ErrNoKeys
	}
	pubKey, prvKey, err := getKeys()
	if err != nil {
		return err
	}

	if _, err = os.Stat(article); err != nil {
		return os.ErrNotExist
	}
	for _, s := range atts {
		if _, err = os.Stat(s); err != nil {
			return os.ErrNotExist
		}
	}

	articleContent, err := ioutil.ReadFile(article)
	if err != nil {
		return err
	}
	articleSing := ed25519.Sign(prvKey, articleContent)

	var attsSing = make(map[string][][]byte)
	for _, f := range atts {
		att, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		attsSing[f] = append(attsSing[f], att)
		attsSing[f] = append(attsSing[f], ed25519.Sign(prvKey, att))
	}

	var serverUrl string
	if server[0:6] == "http://" {
		serverUrl = fmt.Sprintf("%s/articles", server)
	} else {
		serverUrl = fmt.Sprintf("http://%s/articles", server)
	}

	req, err := http.NewRequest("POST", serverUrl, bytes.NewReader(articleContent))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("AuthorKey", base64.StdEncoding.EncodeToString(pubKey))
	req.Header.Set("Signature", base64.StdEncoding.EncodeToString(articleSing))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	uuid, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	serverUrl = fmt.Sprintf("%s/%s/attachments", serverUrl, string(uuid))
	for _, v := range attsSing {
		req, err := http.NewRequest("POST", serverUrl, bytes.NewReader(v[0]))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "text/html")
		req.Header.Set("AuthorKey", base64.StdEncoding.EncodeToString(pubKey))
		req.Header.Set("Signature", base64.StdEncoding.EncodeToString(v[1]))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
	}

	return nil
}
