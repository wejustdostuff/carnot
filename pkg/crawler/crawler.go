package crawler

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// GetFiles ...
func GetFiles(source string) ([]*File, error) {
	files := []*File{}
	log := logrus.WithFields(logrus.Fields{
		"pkg":    "crawler",
		"source": source,
	})

	// Validate source
	if ok, err := isDir(source); err != nil {
		return files, errors.Wrap(err, "could not read source directory")
	} else if !ok {
		return files, errors.New("source is not a directory")
	}

	// Initit exiftool
	et, err := exiftool.NewExiftool()
	if err != nil {
		return files, errors.Wrap(err, "could not init exiftool")
	}
	defer et.Close()

	// Walk source
	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		log := log.WithField("path", path)

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

		file := &File{
			Path:         path,
			Info:         info,
			ExifMetaData: map[string]interface{}{},
		}

		// Retrieve files
		if fileInfos := et.ExtractMetadata(path); len(fileInfos) > 0 && fileInfos[0].Err == nil {
			for _, fileInfo := range fileInfos {
				for key, value := range fileInfo.Fields {
					file.ExifMetaData[key] = value
				}
			}
		}

		// Set file's date
		file.SetDate([]string{"DateTimeOriginal"})

		files = append(files, file)

		return nil
	}); err != nil {
		return files, errors.Wrap(err, "could not read source directory")
	}

	return files, nil
}

// GetMetadata ...
func GetMetadata(files []*File) map[string]bool {
	metadata := map[string]int{}

	for _, file := range files {
		for key := range file.ExifMetaData {
			metadata[key] = metadata[key] + 1
		}
	}

	values := map[string]bool{}
	for key, value := range metadata {
		values[key] = value == len(files)
	}

	return values
}

// isDir helper
func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
