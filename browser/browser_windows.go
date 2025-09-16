//go:build windows

package browser

import (
	"github.com/moond4rk/hackbrowserdata/types"
)

var (
	chromiumList = map[string]struct {
		name        string
		profilePath []string
		storage     string
		dataTypes   []types.DataType
	}{
		"chrome": {
			name:        chromeName,
			profilePath: chromeUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"edge": {
			name:        edgeName,
			profilePath: edgeProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"chromium": {
			name:        chromiumName,
			profilePath: chromiumUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"chrome-beta": {
			name:        chromeBetaName,
			profilePath: chromeBetaUserDataPath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"opera": {
			name:        operaName,
			profilePath: operaProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"opera-gx": {
			name:        operaGXName,
			profilePath: operaGXProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"vivaldi": {
			name:        vivaldiName,
			profilePath: vivaldiProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"coccoc": {
			name:        coccocName,
			profilePath: coccocProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"brave": {
			name:        braveName,
			profilePath: braveProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"yandex": {
			name:        yandexName,
			profilePath: yandexProfilePath,
			dataTypes:   types.DefaultYandexTypes,
		},
		"360Speed": {
			name:        speed360Name,
			profilePath: speed360ProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"360se": {
			name:        brwoser360seName,
			profilePath: brwoser360seProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"qq": {
			name:        qqBrowserName,
			profilePath: qqBrowserProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"dc": {
			name:        dcBrowserName,
			profilePath: dcBrowserProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
		"sogou": {
			name:        sogouName,
			profilePath: sogouProfilePath,
			dataTypes:   types.DefaultChromiumTypes,
		},
	}
	firefoxList = map[string]struct {
		name        string
		storage     string
		profilePath string
		dataTypes   []types.DataType
	}{
		"firefox": {
			name:        firefoxName,
			profilePath: firefoxProfilePath,
			dataTypes:   types.DefaultFirefoxTypes,
		},
	}
)

var (
	chromeUserDataPath     = []string{homeDir + "/AppData/Local/Google/Chrome/User Data/Default/"}
	chromeBetaUserDataPath = []string{homeDir + "/AppData/Local/Google/Chrome Beta/User Data/Default/"}
	chromiumUserDataPath   = []string{homeDir + "/AppData/Local/Chromium/User Data/Default/"}
	edgeProfilePath        = []string{homeDir + "/AppData/Local/Microsoft/Edge/User Data/Default/"}
	braveProfilePath       = []string{homeDir + "/AppData/Local/BraveSoftware/Brave-Browser/User Data/Default/"}
	speed360ProfilePath    = []string{homeDir + "/AppData/Local/360chrome/Chrome/User Data/Default/",
		homeDir + "/AppData/Local/360chromeX/Chrome/User Data/Default/"}
	brwoser360seProfilePath = []string{homeDir + "/AppData/Roaming/360se6/User Data/Default/"}
	qqBrowserProfilePath    = []string{homeDir + "/AppData/Local/Tencent/QQBrowser/User Data/Default/"}
	operaProfilePath        = []string{homeDir + "/AppData/Roaming/Opera Software/Opera Stable/"}
	operaGXProfilePath      = []string{homeDir + "/AppData/Roaming/Opera Software/Opera GX Stable/"}
	vivaldiProfilePath      = []string{homeDir + "/AppData/Local/Vivaldi/User Data/Default/"}
	coccocProfilePath       = []string{homeDir + "/AppData/Local/CocCoc/Browser/User Data/Default/"}
	yandexProfilePath       = []string{homeDir + "/AppData/Local/Yandex/YandexBrowser/User Data/Default/"}
	dcBrowserProfilePath    = []string{homeDir + "/AppData/Local/DCBrowser/User Data/Default/"}
	sogouProfilePath        = []string{homeDir + "/AppData/Roaming/SogouExplorer/Webkit/Default/"}

	firefoxProfilePath = homeDir + "/AppData/Roaming/Mozilla/Firefox/Profiles/"
)
