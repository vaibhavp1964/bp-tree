package fileops

import (
	"io/fs"
	"os"
)

// FileManager controls the interactions with the underlying file
type FileManager struct {
	file     *os.File
	pageSize int
	numPages int
}

// NewFileManager returns a FileManager object taking in a filepath and a pageSize parameter
func NewFileManager(filePath string, pageSize int) (*FileManager, error) {
	// create or open file
	f, err := createOrOpenFile(filePath)

	if err != nil {
		return nil, err
	}

	fileSize, err := getFileSize(f)
	if err != nil {
		return nil, err
	}

	numPages := fileSize / pageSize

	return &FileManager{
		file:     f,
		pageSize: pageSize,
		numPages: numPages,
	}, nil
}

func createOrOpenFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)

	if err != nil && os.IsNotExist(err) {
		// create file
		f, err := os.Create(filePath)

		if err != nil {
			return nil, err
		}
		return f, nil
	}

	// open file with read write permissions and return the file
	f, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func getFileMetadata(f *os.File) (fs.FileInfo, error) {
	return f.Stat()
}

func getFileSize(f *os.File) (int, error) {
	metadata, err := getFileMetadata(f)

	if err != nil {
		return -1, err
	}

	return int(metadata.Size()), nil
}

// Read bytes from the file corresponding to a particular page
func (f *FileManager) Read(pageNum int) ([]byte, error) {
	offset := pageNum * f.pageSize
	_, err := f.file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, f.pageSize)
	_, err = f.file.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

// Write an array of bytes to a particular page
func (f *FileManager) Write(data []byte, pageNum int) error {
	offset := pageNum * f.pageSize
	_, err := f.file.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	_, err = f.file.Write(data)
	return err
}
