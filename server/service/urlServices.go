package service

import (
	"errors"
	"linkmini/config"
	"linkmini/model"
	"linkmini/utils"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/surrealdb/surrealdb.go"
)

var (
	ErrRegisterInDB  = errors.New("failed to register url data in DB")
	ErrDataNotFound  = errors.New("data not found with this hash")
	ErrQrCode        = errors.New("cannot generate the qrcode for this url")
	ErrUpdateURLData = errors.New("fail to update url data click number")
)

func CreateShortURLService(url string) (*model.URL, error) {
	URLHash := utils.GenerateHash()
	CreatedAt := time.Now()
	ExpireAt := CreatedAt.Add(45 * 24 * time.Hour)
	QrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, errors.Join(err, ErrQrCode)
	}

	URLData := model.URL{
		LongURL:   url,
		URLHash:   URLHash,
		CreatedAt: CreatedAt,
		ExpireAt:  ExpireAt,
		QrCode:    string(QrCode),
	}

	data, err := config.DB.Create("url:"+URLHash, URLData)
	if err != nil {
		return nil, errors.Join(err, ErrRegisterInDB)
	}

	createdData := new(model.URL)
	err = surrealdb.Unmarshal(data, &createdData)
	if err != nil {
		panic(err)
	}

	return createdData, nil
}

func GetLongURL(URLHash string) (*model.URL, error) {
	data, err := config.DB.Select("url:" + URLHash)
	if err != nil {
		return nil, errors.Join(err, ErrDataNotFound)
	}

	selectedURL := new(model.URL)
	err = surrealdb.Unmarshal(data, &selectedURL)
	if err != nil {
		panic(err)
	}

	clickNumUpdate := map[string]interface{}{
		"clickNum": selectedURL.ClickNum + 1,
	}
	if _, err = config.DB.Change("url:"+selectedURL.URLHash, clickNumUpdate); err != nil {
		// panic(err)
		return nil, errors.Join(err, ErrUpdateURLData)
	}

	return selectedURL, nil

}
