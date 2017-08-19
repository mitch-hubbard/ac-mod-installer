package support

import "os"

func Exist(filepath string) error {
	_, err := os.Stat(filepath)
	return err
}