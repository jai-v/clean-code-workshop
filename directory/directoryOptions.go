package directoryOptions

import (
	"crypto/sha1"
	"fmt"
	"io/fs"
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
		file, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return err
		}
		err = d.hashFile(file, fullpath, entry)
		if err != nil {
			return err
		}

	}
	return nil
}

func (d *Duplicates) hashFile(file []byte, fullpath string, entry fs.FileInfo) error {
	hash := sha1.New()
	if _, err := hash.Write(file); err != nil {
		return err
	}
	hashSum := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashSum)
	if hashEntry, ok := d.Hashes[hashString]; ok {
		d.DuplicateSlice[hashEntry] = fullpath
		d.DupeSize += entry.Size()
	} else {
		d.Hashes[hashString] = fullpath
	}
	return nil
}
