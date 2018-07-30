package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func hashBytes(input []byte) string {
	h := sha1.New()
	h.Write(input)

	hex := fmt.Sprintf("%x", h.Sum(nil))
	return hex
}

func download(url string, key string, checksum string) error {
	bsClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "beatsaver-downloader")

	res, getErr := bsClient.Do(req)
	if getErr != nil {
		return err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return err
	}

	hash := hashBytes(body)
	if hash != checksum {
		return errors.New("invalid checksum")
	}

	r, err := zip.NewReader(bytes.NewReader(body), res.ContentLength)
	if err != nil {
		return err
	}

	dest := filepath.Join(".", dir, key)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)
		os.MkdirAll(dest, os.ModeDir)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
