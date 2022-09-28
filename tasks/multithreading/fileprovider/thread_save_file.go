package fileprovider

import (
	"bufio"
	"os"
	"sync"
)

type ThreadSaveFile struct {
	sync.Mutex

	file     *os.File
	fileName string
}

func NewThreadSaveFile(fileName string) *ThreadSaveFile {
	return &ThreadSaveFile{fileName: fileName}
}

func (tsf *ThreadSaveFile) GetFileName() string {
	return tsf.fileName
}

func (tsf *ThreadSaveFile) Open() error {
	tsf.Lock()
	f, err := os.OpenFile(tsf.fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		tsf.Unlock()
		return err
	}

	tsf.file = f
	return nil
}

func (tsf *ThreadSaveFile) Read() string {
	r := bufio.NewReader(tsf.file)
	s := bufio.NewScanner(r)
	s.Scan()

	return s.Text()
}

func (tsf *ThreadSaveFile) WriteString(s string) error {
	_, err := tsf.file.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}

func (tsf *ThreadSaveFile) Close() error {
	defer tsf.Unlock()
	err := tsf.file.Close()
	if err != nil {
		return err
	}
	return nil
}
