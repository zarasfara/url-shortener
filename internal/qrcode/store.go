package qrcode

// Store is an interface for storing QR codes.
type Store interface {
	Save(alias, urlToSave string) (string, error)
	Get(alias string) (string, error)
	Delete(alias string) error
}
