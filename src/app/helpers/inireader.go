package helpers

import (
	"os"
	"errors"
	iniread "gopkg.in/ini.v1"
	"fmt"
)

type iniReader struct {
	path string
	file * iniread.File
}


func NewINIReader (path string) (*iniReader, error) {
	k := new(iniReader)
	k.path = path

	if _, e := os.Stat(path); e != nil {
		return nil, errors.New("Unable to find ini file")
	} else {
		inifile, err := iniread.Load(path)
		if err != nil {
			return nil, errors.New("Unable to read ini file")
		} else { k.file = inifile }
	}

	return k, nil
}

func (ini * iniReader) GetValue(section string, key string) (string, error) {
	if section, err := ini.GetIni().GetSection(section); err == nil {
		if secval, errval := section.GetKey(key); errval == nil {
			return secval.String(), nil
		} else { return "", errval }
	}

	return "", fmt.Errorf("Unable to find section '%s'", section)
}

func (ini * iniReader) GetIni() (*iniread.File) {
	return ini.file
}