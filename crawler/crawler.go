package crawler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	//"github.com/rwcarlsen/goexif/exif"
	"github.com/tajtiattila/metadata"
	"gopkg.in/djherbis/times.v1"
)

// Run starts the crawler
func Run(config *Config) error {
	fmt.Println("Source: " + config.Source)
	fmt.Println("Target: " + config.Target)

	// Validate target
	if ok, err := isDir(config.Target); err != nil {
		return errors.Wrap(err, "Could not read target directory")
	} else if !ok {
		return errors.New("Target is not a directory")
	}

	if err := filepath.Walk(config.Source, walk); err != nil {
		return errors.Wrap(err, "Could not read source directory")
	}

	return nil
}

func walk(path string, info os.FileInfo, err error) error {
	// Validate error
	if err != nil {
		return errors.Wrap(err, "Returning nested error")
	}

	// Skip directories
	if info.IsDir() {
		return nil
	}

	logrus.Debug("Handling file: %s", info.Name())
	log.Println("---------------" + info.Name())

	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "Could not open file")
	}
	defer file.Close()

	contentType, err := getContentType(file)
	if err != nil {
		return errors.Wrap(err, "Could not get content type")
	}

	var date time.Time
	if t, err := getTimeMetadata(path, file); err == nil {
		date = t
	}

	log.Println("ContentType:" + contentType)

	log.Println(date.Local().String())

	return nil
}

func getTimeMetadata(path string, file *os.File) (time.Time, error) {
	if m, err := metadata.Parse(file); err == nil {
		if !m.DateTimeOriginal.IsZero() {
			log.Println("Using metadata DateTimeOriginal")
			return m.DateTimeOriginal.Time, nil
		}
		if !m.DateTimeCreated.IsZero() {
			log.Println("Using metadata DateTimeCreated")
			return m.DateTimeCreated.Time, nil
		}
	} else {
		log.Println(err.Error())
	}
	return getTimeDefault(path, file)
}

func getTimeDefault(path string, file *os.File) (time.Time, error) {
	// Try time
	if t, err := times.Stat(path); err == nil {
		if t.HasBirthTime() {
			log.Println("Using times BirthTime")
			return t.BirthTime(), nil
		}
		return t.ModTime(), nil
	}
	return time.Now(), errors.New("Could not get file time")
}

func getContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
