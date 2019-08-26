package crawler

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/hacdias/fileutils"
	"github.com/sirupsen/logrus"
	"gopkg.in/djherbis/times.v1"
)

// File struct
type File struct {
	Path         string
	Info         os.FileInfo
	Date         time.Time
	DateField    string
	ExifMetaData map[string]interface{}
}

// SetDate sets the internal date
func (f *File) SetDate(fields []string) {
	var dateValue string
	var dateField string
	layout := "2006:01:02 15:04:05"
	layoutTZ := "2006:01:02 15:04:05-07:00"

	// Try exif meta data
	for _, field := range fields {
		if value, ok := f.ExifMetaData[field]; ok {
			dateField = field
			dateValue = value.(string)
			break
		}
	}

	if dateValue != "" {
		var dateLayout string
		if strings.Contains(dateValue, "-") || strings.Contains(dateValue, "+") {
			dateLayout = layoutTZ
		} else {
			dateLayout = layout
		}
		if date, err := time.Parse(dateLayout, dateValue); err != nil {
			logrus.WithError(err).Error("could not parse time")
		} else {
			f.Date = date
			f.DateField = dateField
		}
	}

	// Try birth time
	if f.Date.IsZero() {
		if t, err := times.Stat(f.Path); err == nil && t.HasBirthTime() {
			f.Date = t.BirthTime()
			f.DateField = "BirthTime"
		}
	}

	if f.Date.IsZero() {
		f.Date = f.Info.ModTime()
		f.DateField = "ModTime"
	}
}

// Exists returns true if destination file already exists
func (f *File) Exists(dirname string) bool {
	if _, err := os.Stat(f.GetPath(dirname)); !os.IsNotExist(err) {
		return true
	}
	return false
}

// GetPath returns the files destination path
func (f *File) GetPath(dirname string) string {
	return fmt.Sprintf("%s/%s/%s", dirname, f.Date.Format("2006/01/02"), f.GetFilename())
}

// GetFilename returns the files destination name
func (f *File) GetFilename() string {
	format := "20060102-150405"
	return fmt.Sprintf("%s%s", f.Date.Format(format), strings.ToLower(path.Ext(f.Path)))
}

// Move file to destination directory
func (f *File) Move(dirname string) error {
	newPath := f.GetPath(dirname)

	if err := os.MkdirAll(path.Dir(newPath), os.ModePerm); err != nil {
		return err
	}

	return os.Rename(f.Path, newPath)
}

// Copy file to destination directory
func (f *File) Copy(dirname string) error {
	newPath := f.GetPath(dirname)

	if err := os.MkdirAll(path.Dir(newPath), os.ModePerm); err != nil {
		return err
	}

	return fileutils.CopyFile(f.Path, newPath)
}
