package main

import (
	"clean-code-workshop/constants"
	directoryOptions "clean-code-workshop/directory"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func sizeConversion(sizeInBytes int64, conversionSize int64) float64 {
	return float64(sizeInBytes) / float64(conversionSize)
}

func toFloatString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

func toReadableSize(sizeInBytes int64) string {
	switch {
	case sizeInBytes > constants.OneTeraByte:
		sizeInTB := sizeConversion(sizeInBytes, constants.OneTeraByte)
		return toFloatString(sizeInTB) + " TB"

	case sizeInBytes > constants.OneGigaByte:
		sizeInTB := sizeConversion(sizeInBytes, constants.OneGigaByte)
		return toFloatString(sizeInTB) + " GB"

	case sizeInBytes > constants.OneMegaByte:
		sizeInTB := sizeConversion(sizeInBytes, constants.OneMegaByte)
		return toFloatString(sizeInTB) + " MB"

	case sizeInBytes > constants.OneKiloByte:
		sizeInTB := sizeConversion(sizeInBytes, constants.OneKiloByte)
		return toFloatString(sizeInTB) + " KB"
	default:
		return toFloatString(float64(sizeInBytes)) + " B"
	}
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

	d := directoryOptions.NewDuplicates()

	err = d.TraverseDir(*dir)
	if err != nil {
		panic(err)
	}

	fmt.Println("DUPLICATES")

	fmt.Println("TOTAL FILES:", len(d.Hashes))
	fmt.Println("DUPLICATES:", len(d.DuplicateSlice))
	fmt.Println("TOTAL DUPLICATE SIZE:", toReadableSize(d.DupeSize))
}

// running into problems of not being able to open directories inside .app folders
