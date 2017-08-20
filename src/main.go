package main

import (
	"fmt"
	"os"

	"./support"
	"github.com/mholt/archiver"
)

func main() {
	modpath := ""
	acpath := ""

	// check if the user supplied arguments
	if len(os.Args[1:]) > 1 {
		// loop over the arguments to the application
		for index, element := range os.Args {
			argumentValue := ""
			if len(os.Args) > index+1 {
				argumentValue = os.Args[index+1]
			}
			if element == "-m" {
				modpath = argumentValue
			} else if element == "-d" {
				acpath = argumentValue
				if support.Exist(acpath) != nil {
					fmt.Println("The specified Assetto Corsa path does not exist")
					os.Exit(1)
				}
			}
		}
	} else {
		fmt.Println("you must specify the mod with -m /path/to/mod")
		os.Exit(1)
	}

	acpath, err := getAssettoCorsaPath()
	if err != nil {
		panic(err)
	}

	err = archiver.Zip.Open(modpath, acpath)
	if err != nil {
		panic(err)
	}

	fmt.Println("installed " + modpath + " to " + acpath)
	os.Exit(0)
}

func getAssettoCorsaPath() (string, error) {
	acpath := "/Program Files (x86)/Steam/steamapps/common/assettocorsa"
	_, err := os.Stat(acpath)
	if err != nil {
		return "", err
	}

	return acpath, nil
}
