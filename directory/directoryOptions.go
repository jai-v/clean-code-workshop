package directoryOptions

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"path"
)

type Duplicates struct {
	Hashes, DuplicateSlice map[string]string
	DupeSize               int64
}

func NewDuplicates() *Duplicates {
	return &Duplicates{
		Hashes:         map[string]string{},
		DuplicateSlice: map[string]string{},
		DupeSize:       0,
	}
}

func (d *Duplicates) TraverseDir(directory string) error {
	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fullpath := (path.Join(directory, entry.Name()))

		if !entry.Mode().IsDir() && !entry.Mode().IsRegular() {
			continue
		}
		if entry.IsDir() {
			d.TraverseDir(fullpath)
			continue
		}
		err = d.checkHash(fullpath, entry.Size())
		if err != nil {
			return err
		}
	}
	return nil
}

func getHash(fullpath string) (string, error) {
	file, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	if _, err := hash.Write(file); err != nil {
		return "", err
	}
	hashSum := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashSum)
	return hashString, nil
}

func (d *Duplicates) checkHash(fullpath string, fileSize int64) error {
	hashString, err := getHash(fullpath)
	if err != nil {
		return err
	}
	if hashEntry, ok := d.Hashes[hashString]; ok {
		d.DuplicateSlice[hashEntry] = fullpath
		d.DupeSize += fileSize
	} else {
		d.Hashes[hashString] = fullpath
	}
	return nil
}
