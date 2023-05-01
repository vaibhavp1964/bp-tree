package diskmanager

import (
	"log"
	"os"
)

// Impl ...
type Impl struct {
	log *log.Logger
}

// NewDiskManagerImpl ...
func NewDiskManagerImpl(log *log.Logger) *Impl {
	return &Impl{
		log: log,
	}
}

func (dmi *Impl) Read(fileName string, offset, amount int) ([]byte, error) {
	// open file in read mode and defer closing the file
	f, err := os.Open(fileName)
	if err != nil {
		dmi.log.Fatalln("error encountered while opening file:", fileName, "in read-only mode:", err)
		return nil, err
	}
	defer f.Close()

	// read and return data
	b := make([]byte, amount)
	n, err := f.ReadAt(b, int64(offset))
	if err != nil {
		dmi.log.Fatalln("error encountered while reading data from file:", fileName, ":", err)
		return nil, err
	}
	return b[:n], nil
}

func (dmi *Impl) Write(fileName string, offset int, data []byte) error {
	// open file in write only mode and defer closing the file
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		dmi.log.Fatalln("error encountered while opening file:", fileName, "in write-only mode:", err)
		return err
	}
	defer f.Close()

	// write the data to the file
	_, err = f.WriteAt(data, int64(offset))
	if err != nil {
		dmi.log.Fatalln("error encountered while writing data to file:", fileName, ":", err)
	}
	return err
}
