package crawler

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/hacdias/fileutils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Run starts the crawler
func Run(source, target string, dry, move, force bool) error {
	logrus.WithFields(logrus.Fields{
		"dry":    dry,
		"move":   move,
		"force":  force,
		"source": source,
		"target": target,
	}).Info("crawling directory")

	// Validate source
	if ok, err := isDir(source); err != nil {
		return errors.Wrap(err, "could not read source directory")
	} else if !ok {
		return errors.New("source is not a directory")
	}

	// Validate target
	if ok, err := isDir(target); err != nil {
		return errors.Wrap(err, "could not read target directory")
	} else if !ok {
		return errors.New("target is not a directory")
	}

	// Initit exiftool
	et, err := exiftool.NewExiftool()
	if err != nil {
		return errors.Wrap(err, "could not init exiftool")
	}
	defer et.Close()

	// Walk source
	if err := filepath.Walk(source, func(filename string, info os.FileInfo, err error) error {
		log := logrus.WithField("src", filename)

		// Validate error
		if err != nil {
			return errors.Wrap(err, "returning nested error")
		}

		// Skip directories
		if info.IsDir() {
			log.Debug("skipping directory")
			return nil
		}

		// Skip dot files
		if strings.HasPrefix(info.Name(), ".") {
			log.Debug("skipping dot file")
			return nil
		}

		// Retrieve file date
		date, err := getTime(et, filename)
		if err != nil {
			log.WithError(err).Error("could not get file time")
			return nil
		}

		targetName := fmt.Sprintf("%s%s", date.Format("20060102-150405"), strings.ToLower(path.Ext(filename)))
		targetDirname := fmt.Sprintf("%s/%s", target, date.Format("2006/01/02"))
		targetFilename := fmt.Sprintf("%s/%s", targetDirname, targetName)

		log = log.WithField("dest", targetFilename)

		// Validate target
		if _, err := os.Stat(targetFilename); !os.IsNotExist(err) && !force {
			log.Warn("destination already exists, skipping file")
			return nil
		}

		// Copy/Move file
		if !dry {
			// Create target directories
			if err := os.MkdirAll(targetDirname, os.ModePerm); err != nil {
				log.WithError(err).Error("could not create target directory")
				return nil
			}

			if move {
				if err := os.Rename(filename, targetFilename); err != nil {
					log.WithError(err).Error("could not move file")
					return nil
				}
			} else {
				if err := fileutils.CopyFile(filename, targetFilename); err != nil {
					log.WithError(err).Error("could not copy file")
					return nil
				}
			}
		} else {
			if move {
				log.Info("moving file")
			} else {
				log.Info("copying file")
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "could not read source directory")
	}

	return nil
}

func getTime(et *exiftool.Exiftool, path string) (time.Time, error) {
	layout := "2006:01:02 15:04:05"
	layoutTimezone := "2006:01:02 15:04:05-07:00"

	if fileInfos := et.ExtractMetadata(path); len(fileInfos) > 0 && fileInfos[0].Err == nil {
		if value, ok := fileInfos[0].Fields["DateTimeOriginal"]; ok {
			return time.Parse(layout, value.(string))
		}
		if value, ok := fileInfos[0].Fields["FileModifyDate"]; ok {
			return time.Parse(layoutTimezone, value.(string))
		}
		if value, ok := fileInfos[0].Fields["ModifyDate"]; ok {
			return time.Parse(layout, value.(string))
		}
	}

	// Fallback solution
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Now(), err
	}

	return fileInfo.ModTime(), nil
}

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
