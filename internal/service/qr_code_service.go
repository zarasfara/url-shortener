package service

import "github.com/zarasfara/url-shortener/internal/qrcode"

// TODO: Подумать об интерфейсе
type QRCodeService struct {
	store qrcode.Store
}

func NewQRCodeService(store qrcode.Store) *QRCodeService {
	return &QRCodeService{store: store}
}

func (qr *QRCodeService) Save(alias, urlToSave string) (string, error) {
	return qr.store.Save(alias, urlToSave)
}

func (qr *QRCodeService) Get(alias string) (string, error) {
	return qr.store.Get(alias)
}
