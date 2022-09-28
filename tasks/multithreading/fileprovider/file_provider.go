package fileprovider

type FileProvider interface {
	GetFileName() string
	Open() error
	Read() string
	WriteString(s string) error
	Close() error
}
