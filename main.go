package main

import (
	"clean-code-workshop/constants"
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync/atomic"
)

func traverseDir(hashes, duplicates map[string]string, dupeSize *int64, entries []os.FileInfo, directory string) {
	for _, entry := range entries {
		fullpath := (path.Join(directory, entry.Name()))

		if !entry.Mode().IsDir() && !entry.Mode().IsRegular() {
			continue
		}

		if entry.IsDir() {
			dirFiles, err := ioutil.ReadDir(fullpath)
			if err != nil {
				panic(err)
			}
			traverseDir(hashes, duplicates, dupeSize, dirFiles, fullpath)
			continue
		}
		file, err := ioutil.ReadFile(fullpath)
		if err != nil {
			panic(err)
		}
		hash := sha1.New()
		if _, err := hash.Write(file); err != nil {
			panic(err)
		}
		hashSum := hash.Sum(nil)
		hashString := fmt.Sprintf("%x", hashSum)
		if hashEntry, ok := hashes[hashString]; ok {
			duplicates[hashEntry] = fullpath
			atomic.AddInt64(dupeSize, entry.Size())
		} else {
			hashes[hashString] = fullpath
		}
	}
}

func sizeConversion(sizeInBytes int64, conversionSize int64) float64 {
	return float64(sizeInBytes) / float64(conversionSize)
}

func toFloatString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 2)
}

func toReadableSize(sizeInBytes int64) string {
	if sizeInBytes > constants.OneTeraByte {
		sizeInTB := sizeConversion(sizeInBytes, constants.OneTeraByte)
		return toFloatString(sizeInTB) + " TB"
	}
	if sizeInBytes > constants.OneGigaByte {
		sizeInTB := sizeConversion(sizeInBytes, constants.OneGigaByte)
		return toFloatString(sizeInTB) + " GB"
	}
	if sizeInBytes > constants.OneMegaByte {
		sizeInTB := sizeConversion(sizeInBytes, constants.OneMegaByte)
		return toFloatString(sizeInTB) + " MB"
	}

	if sizeInBytes > constants.OneKiloByte {
		sizeInTB := sizeConversion(sizeInBytes, constants.OneKiloByte)
		return toFloatString(sizeInTB) + " KB"
	}
	return toFloatString(float64(sizeInBytes)) + " B"
}

func main() {
	var err error
	dir := flag.String("path", "", "the path to traverse searching for duplicates")
	flag.Parse()

	if *dir == "" {
		*dir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	hashes := map[string]string{}
	duplicates := map[string]string{}
	var dupeSize int64

	entries, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	traverseDir(hashes, duplicates, &dupeSize, entries, *dir)

	fmt.Println("DUPLICATES")

	fmt.Println("TOTAL FILES:", len(hashes))
	fmt.Println("DUPLICATES:", len(duplicates))
	fmt.Println("TOTAL DUPLICATE SIZE:", toReadableSize(dupeSize))
}

// running into problems of not being able to open directories inside .app folders
