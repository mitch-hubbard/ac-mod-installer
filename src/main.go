package main

import (
	"fmt"
	"github.com/mholt/archiver"
	"os"
)

func main() {
	modpath := ""

	// check if the user supplied arguments
	if len(os.Args[1:]) > 1 {
		// loop over the arguments to the application
		for index,element := range os.Args[1:] {
			if element == "-m"{
				modpath = os.Args[index+2]
			}
		}
	} else{
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

func getAssettoCorsaPath() (string,error) {
	acpath := "/Program Files (x86)/Steam/steamapps/common/assettocorsa"
	_, err  := os.Stat(acpath)
	if err != nil {
		return "", err
	}

	return acpath, nil
}
