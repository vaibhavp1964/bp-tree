package diskmanager

// DiskManager ...
type DiskManager interface {
	Write(fileName string, offset int, data []byte) error
	Read(fileName string, offset, amount int) ([]byte, error)
}
