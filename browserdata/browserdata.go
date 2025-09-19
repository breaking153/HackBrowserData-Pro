package browserdata

import (
	"fmt"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/bookmark"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/cookie"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/creditcard"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/download"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/extension"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/history"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/localstorage"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/password"
	"github.com/breaking153/HackBrowserData-Pro/browserdata/sessionstorage"
	"github.com/breaking153/HackBrowserData-Pro/extractor"
	"github.com/breaking153/HackBrowserData-Pro/log"
	"github.com/breaking153/HackBrowserData-Pro/types"
	"github.com/breaking153/HackBrowserData-Pro/utils/fileutil"
)

type BrowserData struct {
	Extractors map[types.DataType]extractor.Extractor
}
type BrowserOutput struct {
	Cookies        []cookie.Cookie
	BookMarks      []bookmark.Bookmark
	SessionStorage []sessionstorage.Session
	LocalStorage   []localstorage.Storage
	Passwords      []password.LoginData
	History        []history.History
	Extensions     []*extension.Extension
	CreditCards    []creditcard.Card
	Downloads      []download.Download
}

func (b *BrowserOutput) PrinttoStr() {
	if b.Cookies != nil {
		log.Warnf("browser data type cookies, data count %d", len(b.Cookies))
	}
	if b.BookMarks != nil {
		log.Warnf("browser data type bookmarks, data count %d", len(b.BookMarks))
	}
	if b.SessionStorage != nil {
		log.Warnf("browser data type session storage, data count %d", len(b.SessionStorage))
	}
	if b.LocalStorage != nil {
		log.Warnf("browser data type local storage, data count %d", len(b.LocalStorage))
	}
	if b.Passwords != nil {
		log.Warnf("browser data type passwords, data count %d", len(b.Passwords))
	}
	if b.History != nil {
		log.Warnf("browser data type history, data count %d", len(b.History))
	}
	if b.Extensions != nil {
		log.Warnf("browser data type extensions, data count %d", len(b.Extensions))
	}
	if b.CreditCards != nil {
		log.Warnf("browser data type credit cards, data count %d", len(b.CreditCards))
	}
	if b.Downloads != nil {
		log.Warnf("browser data type downloads, data count %d", len(b.Downloads))
	}
}

// ... existing code ...
func (d *BrowserData) ToBrowserOutput() BrowserOutput {
	bout := BrowserOutput{}
	// 处理Cookies
	if cookieExtractor, ok := d.Extractors[types.ChromiumCookie].(*cookie.ChromiumCookie); ok {
		if bout.Cookies == nil {
			bout.Cookies = *cookieExtractor
		} else {
			// 为什么会已经有了？
			fmt.Print("found conflicting cookies, append to existing cookies")
			bout.Cookies = append(bout.Cookies, *cookieExtractor...)
		}
	}
	// 处理Bookmarks
	if bookmarkExtractor, ok := d.Extractors[types.ChromiumBookmark].(*bookmark.ChromiumBookmark); ok {
		if bout.BookMarks == nil {
			bout.BookMarks = *bookmarkExtractor
		} else {
			fmt.Print("found conflicting bookmarks, append to existing bookmarks")
			bout.BookMarks = append(bout.BookMarks, *bookmarkExtractor...)
		}
	}
	// 处理SessionStorage
	if sessionExtractor, ok := d.Extractors[types.ChromiumSessionStorage].(*sessionstorage.ChromiumSessionStorage); ok {
		if bout.SessionStorage == nil {
			bout.SessionStorage = *sessionExtractor
		} else {
			fmt.Print("found conflicting session storage, append to existing session storage")
			bout.SessionStorage = append(bout.SessionStorage, *sessionExtractor...)
		}
	}
	// 处理Passwords
	if passwordExtractor, ok := d.Extractors[types.ChromiumPassword].(*password.ChromiumPassword); ok {
		if bout.Passwords == nil {
			bout.Passwords = *passwordExtractor
		} else {
			fmt.Print("found conflicting passwords, append to existing passwords")
			bout.Passwords = append(bout.Passwords, *passwordExtractor...)
		}
	}
	// 处理History
	if historyExtractor, ok := d.Extractors[types.ChromiumHistory].(*history.ChromiumHistory); ok {
		if bout.History == nil {
			bout.History = *historyExtractor
		} else {
			fmt.Print("found conflicting history, append to existing history")
			bout.History = append(bout.History, *historyExtractor...)
		}
	}
	// 处理Extensions
	if extensionExtractor, ok := d.Extractors[types.ChromiumExtension].(*extension.ChromiumExtension); ok {
		if bout.Extensions == nil {
			bout.Extensions = *extensionExtractor
		} else {
			fmt.Print("found conflicting extensions, append to existing extensions")
			bout.Extensions = append(bout.Extensions, *extensionExtractor...)
		}
	}
	// chrome-download
	if downloadExtractor, ok := d.Extractors[types.ChromiumDownload].(*download.ChromiumDownload); ok {
		if bout.Downloads == nil {
			bout.Downloads = *downloadExtractor
		} else {
			fmt.Print("found conflicting chrome downloads, append to existing chrome downloads")
			bout.Downloads = append(bout.Downloads, *downloadExtractor...)
		}
	}
	// chrome-localstorage
	if localStorageExtractor, ok := d.Extractors[types.ChromiumLocalStorage].(*localstorage.ChromiumLocalStorage); ok {
		if bout.LocalStorage == nil {
			bout.LocalStorage = *localStorageExtractor
		} else {
			fmt.Print("found conflicting chrome local storage, append to existing chrome local storage")
			bout.LocalStorage = append(bout.LocalStorage, *localStorageExtractor...)
		}
	}
	// chrome-creditcard
	if creditCardExtractor, ok := d.Extractors[types.ChromiumCreditCard].(*creditcard.ChromiumCreditCard); ok {
		if bout.CreditCards == nil {
			bout.CreditCards = *creditCardExtractor
		} else {
			fmt.Print("found conflicting credit cards, append to existing credit cards")
			bout.CreditCards = append(bout.CreditCards, *creditCardExtractor...)
		}
	}
	// firefox-bookmark
	if bookmarkExtractor, ok := d.Extractors[types.FirefoxBookmark].(*bookmark.FirefoxBookmark); ok {
		if bout.BookMarks == nil {
			bout.BookMarks = *bookmarkExtractor
		} else {
			fmt.Print("found conflicting firefox bookmarks, append to existing firefox bookmarks")
			bout.BookMarks = append(bout.BookMarks, *bookmarkExtractor...)
		}
	}
	// firefox-history
	if historyExtractor, ok := d.Extractors[types.FirefoxHistory].(*history.FirefoxHistory); ok {
		if bout.History == nil {
			bout.History = *historyExtractor
		} else {
			fmt.Print("found conflicting firefox history, append to existing firefox history")
			bout.History = append(bout.History, *historyExtractor...)
		}
	}
	// firefox-extension
	if extensionExtractor, ok := d.Extractors[types.FirefoxExtension].(*extension.FirefoxExtension); ok {
		if bout.Extensions == nil {
			bout.Extensions = *extensionExtractor
		} else {
			fmt.Print("found conflicting firefox extensions, append to existing firefox extensions")
			bout.Extensions = append(bout.Extensions, *extensionExtractor...)
		}
	}
	// firefox-download
	if downloadExtractor, ok := d.Extractors[types.FirefoxDownload].(*download.FirefoxDownload); ok {
		if bout.Downloads == nil {
			bout.Downloads = *downloadExtractor
		} else {
			fmt.Print("found conflicting firefox downloads, append to existing firefox downloads")
			bout.Downloads = append(bout.Downloads, *downloadExtractor...)
		}
	}
	// firefox-localstorage
	if localStorageExtractor, ok := d.Extractors[types.FirefoxLocalStorage].(*localstorage.FirefoxLocalStorage); ok {
		if bout.LocalStorage == nil {
			bout.LocalStorage = *localStorageExtractor
		} else {
			fmt.Print("found conflicting firefox local storage, append to existing firefox local storage")
			bout.LocalStorage = append(bout.LocalStorage, *localStorageExtractor...)
		}
	}
	// firefox-password
	if passwordExtractor, ok := d.Extractors[types.FirefoxPassword].(*password.FirefoxPassword); ok {
		if bout.Passwords == nil {
			bout.Passwords = *passwordExtractor
		} else {
			fmt.Print("found conflicting firefox passwords, append to existing firefox passwords")
			bout.Passwords = append(bout.Passwords, *passwordExtractor...)
		}
	}
	// firefox-session
	if sessionExtractor, ok := d.Extractors[types.FirefoxSessionStorage].(*sessionstorage.FirefoxSessionStorage); ok {
		if bout.SessionStorage == nil {
			bout.SessionStorage = *sessionExtractor
		} else {
			fmt.Print("found conflicting firefox session, append to existing firefox session")
			bout.SessionStorage = append(bout.SessionStorage, *sessionExtractor...)
		}
	}
	// yandex-creditcard
	if creditCardExtractor, ok := d.Extractors[types.YandexCreditCard].(*creditcard.YandexCreditCard); ok {
		if bout.CreditCards == nil {
			bout.CreditCards = *creditCardExtractor
		} else {
			fmt.Print("found conflicting yandex credit cards, append to existing yandex credit cards")
			bout.CreditCards = append(bout.CreditCards, *creditCardExtractor...)
		}
	}
	return bout
} // ... existing code ...

func New(items []types.DataType) *BrowserData {
	bd := &BrowserData{
		Extractors: make(map[types.DataType]extractor.Extractor),
	}
	bd.addExtractors(items)
	return bd
}

func (d *BrowserData) Recovery(masterKey []byte) error {
	// masterkey hex编码为string
	//hexString := hex.EncodeToString(masterKey)
	//println(hexString)
	for _, source := range d.Extractors {
		if err := source.Extract(masterKey); err != nil {
			log.Debugf("parse %s error: %v", source.Name(), err)
			continue
		}
	}
	return nil
}

func (d *BrowserData) Output(dir, browserName, flag string) {
	output := newOutPutter(flag)

	for _, source := range d.Extractors {
		if source.Len() == 0 {
			// if the length of the export data is 0, then it is not necessary to output
			continue
		}
		filename := fileutil.Filename(browserName, source.Name(), output.Ext())

		f, err := output.CreateFile(dir, filename)
		if err != nil {
			log.Debugf("create file %s error: %v", filename, err)
			continue
		}
		if err := output.Write(source, f); err != nil {
			log.Debugf("write to file %s error: %v", filename, err)
			continue
		}
		if err := f.Close(); err != nil {
			log.Debugf("close file %s error: %v", filename, err)
			continue
		}
		log.Warnf("export success: %s", filename)
	}
}

func (d *BrowserData) addExtractors(items []types.DataType) {
	for _, itemType := range items {
		if source := extractor.CreateExtractor(itemType); source != nil {
			d.Extractors[itemType] = source
		} else {
			log.Debugf("source not found: %s", itemType)
		}
	}
}
