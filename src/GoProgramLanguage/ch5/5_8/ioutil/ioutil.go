package ioutil

import "os"

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}

func ReadAll(f *os.File) ([]byte, error) {
	return nil, nil
}
