package conf

import (
	"errors"
	"fmt"

	"github.com/go-akka/configuration"
)

var conf *configuration.Config

func Open(fileNames ...string) (err error) {
	for _, f := range fileNames {
		if err = openFile(f); err == nil {
			return
		}
	}

	return errors.New("Didn't open or found Config File")
}

func openFile(fileName string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	fmt.Println(fileName)
	conf = configuration.LoadConfig(fileName)
	return
}
