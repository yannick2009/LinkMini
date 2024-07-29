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
	CreatedAt := time.Now().UTC()
	ExpireAt := CreatedAt.Add(45 * 24 * time.Hour).UTC()
	QrCode, err := qrcode.Encode(URLHash, qrcode.Medium, 256)
	if err != nil {
		return nil, errors.Join(err, ErrQrCode)
	}

	URLData := model.URL{
		LongURL:   url,
		URLHash:   URLHash,
		CreatedAt: CreatedAt,
		ExpireAt:  ExpireAt,
		QrCode:    QrCode,
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

	return selectedURL, nil

}

func GetLongURLAndUpdate(URLHash string) (*model.URL, error) {
	data, err := config.DB.Select("url:" + URLHash)
	if err != nil {
		return nil, errors.Join(err, ErrDataNotFound)
	}

	selectedURL := new(model.URL)
	err = surrealdb.Unmarshal(data, &selectedURL)
	if err != nil {
		panic(err)
	}

	copySelectedURL := selectedURL
	copySelectedURL.ClickNum++
	config.DB.Update("url:"+selectedURL.URLHash, copySelectedURL)

	return selectedURL, nil
}

func DeleteURLExpired(URLHash string) error {
	_, err := config.DB.Delete("url:" + URLHash)
	return err
}
