package qr

//const qrPath = "tmp"
//
//func GenerateQRCode(url, name string) error {
//	absPath, err := filepath.Abs(qrPath)
//	if err != nil {
//		return err
//	}
//
//	// create dir if not exists
//	err = os.MkdirAll(absPath, os.ModePerm)
//	if err != nil {
//		return err
//	}
//
//	filePath := filepath.Join(absPath, fmt.Sprintf("%s.png", name))
//
//	err = qrcode.WriteFile(url, qrcode.Medium, 256, filePath)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
