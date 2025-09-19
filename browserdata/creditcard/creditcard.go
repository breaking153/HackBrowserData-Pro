package creditcard

import (
	"database/sql"
	"os"

	"github.com/breaking153/HackBrowserData-Pro/crypto"
	"github.com/breaking153/HackBrowserData-Pro/extractor"
	"github.com/breaking153/HackBrowserData-Pro/log"
	"github.com/breaking153/HackBrowserData-Pro/types"
	// import sqlite3 driver
	_ "modernc.org/sqlite"
)

func init() {
	extractor.RegisterExtractor(types.ChromiumCreditCard, func() extractor.Extractor {
		return new(ChromiumCreditCard)
	})
	extractor.RegisterExtractor(types.YandexCreditCard, func() extractor.Extractor {
		return new(YandexCreditCard)
	})
}

type ChromiumCreditCard []Card

type Card struct {
	GUID            string
	Name            string
	ExpirationYear  string
	ExpirationMonth string
	CardNumber      string
	Address         string
	NickName        string
}

const (
	queryChromiumCredit = `SELECT guid, name_on_card, expiration_month, expiration_year, card_number_encrypted, billing_address_id, nickname FROM credit_cards`
)

func (c *ChromiumCreditCard) Extract(masterKey []byte) error {
	db, err := sql.Open("sqlite", types.ChromiumCreditCard.TempFilename())
	if err != nil {
		return err
	}
	defer os.Remove(types.ChromiumCreditCard.TempFilename())
	defer db.Close()

	rows, err := db.Query(queryChromiumCredit)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			name, month, year, guid, address, nickname string
			value, encryptValue                        []byte
		)
		if err := rows.Scan(&guid, &name, &month, &year, &encryptValue, &address, &nickname); err != nil {
			log.Debugf("scan chromium credit Card error: %v", err)
		}
		ccInfo := Card{
			GUID:            guid,
			Name:            name,
			ExpirationMonth: month,
			ExpirationYear:  year,
			Address:         address,
			NickName:        nickname,
		}
		if len(encryptValue) > 0 {
			if len(masterKey) == 0 {
				value, err = crypto.DecryptWithDPAPI(encryptValue)
			} else {
				value, err = crypto.DecryptWithChromium(masterKey, encryptValue)
			}
			if err != nil {
				log.Debugf("decrypt chromium credit Card error: %v", err)
			}
		}

		ccInfo.CardNumber = string(value)
		*c = append(*c, ccInfo)
	}
	return nil
}

func (c *ChromiumCreditCard) Name() string {
	return "creditcard"
}

func (c *ChromiumCreditCard) Len() int {
	return len(*c)
}

type YandexCreditCard []Card

func (c *YandexCreditCard) Extract(masterKey []byte) error {
	db, err := sql.Open("sqlite", types.YandexCreditCard.TempFilename())
	if err != nil {
		return err
	}
	defer os.Remove(types.YandexCreditCard.TempFilename())
	defer db.Close()
	rows, err := db.Query(queryChromiumCredit)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			name, month, year, guid, address, nickname string
			value, encryptValue                        []byte
		)
		if err := rows.Scan(&guid, &name, &month, &year, &encryptValue, &address, &nickname); err != nil {
			log.Debugf("scan chromium credit Card error: %v", err)
		}
		ccInfo := Card{
			GUID:            guid,
			Name:            name,
			ExpirationMonth: month,
			ExpirationYear:  year,
			Address:         address,
			NickName:        nickname,
		}
		if len(encryptValue) > 0 {
			if len(masterKey) == 0 {
				value, err = crypto.DecryptWithDPAPI(encryptValue)
			} else {
				value, err = crypto.DecryptWithChromium(masterKey, encryptValue)
			}
			if err != nil {
				log.Debugf("decrypt chromium credit Card error: %v", err)
			}
		}
		ccInfo.CardNumber = string(value)
		*c = append(*c, ccInfo)
	}
	return nil
}

func (c *YandexCreditCard) Name() string {
	return "creditcard"
}

func (c *YandexCreditCard) Len() int {
	return len(*c)
}
