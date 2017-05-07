package main

import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/base64"
	"net/http"
	"github.com/hsmty/LibreScienceJournal/crypto"
)

func CreateKeys(force bool) error {
	if force == false {
		if crypto.KeysExists() == true {
			return crypto.ErrKeysExist
		}
	}

	return crypto.CreateKeys()
}

func PublishArticle(server string, article string, atts []string) error {
	var err error

	if crypto.KeysExists() == false {
		return crypto.ErrNoKeys
	}

	pubKey, err := crypto.GetPubKey()
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
	articleSing, err := crypto.SignContent(articleContent)
	if err != nil {
		return err
	}

	var attsSing = make(map[string][][]byte)
	for _, f := range atts {
		att, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		attsSing[f] = append(attsSing[f], att)
		attSign, err := crypto.SignContent(att)
		if err != nil {
			return err
		}
		attsSing[f] = append(attsSing[f], attSign)
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
