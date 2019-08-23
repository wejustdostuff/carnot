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
func Run(config *Config) error {
	// Validate target
	if ok, err := isDir(config.Target); err != nil {
		return errors.Wrap(err, "Could not read target directory")
	} else if !ok {
		return errors.New("Target is not a directory")
	}

	// Initit exiftool
	et, err := exiftool.NewExiftool()
	if err != nil {
		return errors.Wrap(err, "Could not init exiftool")
	}
	defer et.Close()

	// Walk source
	if err := filepath.Walk(config.Source, func(filename string, info os.FileInfo, err error) error {
		log := logrus.WithField("source", filename)

		// Validate error
		if err != nil {
			return errors.Wrap(err, "Returning nested error")
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Retrieve file date
		date, err := getTime(et, filename, log)
		if err != nil {
			return errors.Wrap(err, "Could not get file date")
		}

		targetName := fmt.Sprintf("%s%s", date.Format("20060102-150405"), strings.ToLower(path.Ext(filename)))
		targetDirname := fmt.Sprintf("%s/%s", config.Target, date.Format("2006/01/02"))
		targetFilename := fmt.Sprintf("%s/%s", targetDirname, targetName)

		log = log.WithField("target", targetFilename)

		if !config.Dry {
			// Create target directories
			if err := os.MkdirAll(targetDirname, os.ModePerm); err != nil {
				log.WithError(err).Error("Could not create target directory")
			}

			// Copy/Move file
			if config.Move {
				log.Info("Move")
				if err := os.Rename(filename, targetFilename); err != nil {
					log.WithError(err).Error("Could not move file")
				}
			} else {
				log.Info("Copy")
				if err := fileutils.CopyFile(filename, targetFilename); err != nil {
					log.WithError(err).Error("Could not copy file")
				}
			}
		} else {
			if config.Move {
				log.Info("Move")
			} else {
				log.Info("Copy")
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "Could not read source directory")
	}

	return nil
}

func getTime(et *exiftool.Exiftool, path string, log *logrus.Entry) (time.Time, error) {
	layout := "2006:01:02 15:04:05"
	layoutTimezone := "2006:01:02 15:04:05-07:00"

	if fileInfos := et.ExtractMetadata(path); len(fileInfos) > 0 {
		if fileInfos[0].Err != nil {
			log.WithError(fileInfos[0].Err).Debug("Exiftool infos contain errors")
		} else {
			if value, ok := fileInfos[0].Fields["DateTimeOriginal"]; ok {
				log.Debug("Using exiftool's DateTimeOriginal")
				return time.Parse(layout, value.(string))
			}
			if value, ok := fileInfos[0].Fields["FileModifyDate"]; ok {
				log.Debug("Using exiftool's FileModifyDate")
				return time.Parse(layoutTimezone, value.(string))
			}
			if value, ok := fileInfos[0].Fields["ModifyDate"]; ok {
				log.Debug("Using exiftool's ModifyDate")
				return time.Parse(layout, value.(string))
			}
		}
	}

	// Fallback solution
	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Now(), nil
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
