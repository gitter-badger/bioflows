package helpers

import "os"


type FileDetails struct {
	Base string
	FileName string
	Scheme string
	Local bool
}

func WriteOrAppend(filename string , data []byte , perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
