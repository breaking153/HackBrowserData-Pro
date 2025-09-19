package extension

import (
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"

	"github.com/breaking153/HackBrowserData-Pro/extractor"
	"github.com/breaking153/HackBrowserData-Pro/types"
	"github.com/breaking153/HackBrowserData-Pro/utils/fileutil"
)

func init() {
	extractor.RegisterExtractor(types.ChromiumExtension, func() extractor.Extractor {
		return new(ChromiumExtension)
	})
	extractor.RegisterExtractor(types.FirefoxExtension, func() extractor.Extractor {
		return new(FirefoxExtension)
	})
}

type ChromiumExtension []*Extension

type Extension struct {
	ID          string
	URL         string
	Enabled     bool
	Name        string
	Description string
	Version     string
	HomepageURL string
}

func (c *ChromiumExtension) Extract(_ []byte) error {
	extensionFile, err := fileutil.ReadFile(types.ChromiumExtension.TempFilename())
	if err != nil {
		return err
	}
	defer os.Remove(types.ChromiumExtension.TempFilename())

	result, err := parseChromiumExtensions(extensionFile)
	if err != nil {
		return err
	}
	*c = result
	return nil
}

func parseChromiumExtensions(content string) ([]*Extension, error) {
	settingKeys := []string{
		"settings.extensions",
		"settings.settings",
		"extensions.settings",
	}
	var settings gjson.Result
	for _, key := range settingKeys {
		settings = gjson.Parse(content).Get(key)
		if settings.Exists() {
			break
		}
	}
	if !settings.Exists() {
		return nil, fmt.Errorf("cannot find extensions in settings")
	}
	var c []*Extension

	settings.ForEach(func(id, ext gjson.Result) bool {
		location := ext.Get("location")
		if !location.Exists() {
			return true
		}
		switch location.Int() {
		case 5, 10: // https://source.chromium.org/chromium/chromium/src/+/main:extensions/common/mojom/manifest.mojom
			return true
		}
		// https://source.chromium.org/chromium/chromium/src/+/main:extensions/browser/disable_reason.h
		enabled := !ext.Get("disable_reasons").Exists()
		b := ext.Get("manifest")
		if !b.Exists() {
			c = append(c, &Extension{
				ID:      id.String(),
				Enabled: enabled,
				Name:    ext.Get("path").String(),
			})
			return true
		}
		c = append(c, &Extension{
			ID:          id.String(),
			URL:         getChromiumExtURL(id.String(), b.Get("update_url").String()),
			Enabled:     enabled,
			Name:        b.Get("name").String(),
			Description: b.Get("description").String(),
			Version:     b.Get("version").String(),
			HomepageURL: b.Get("homepage_url").String(),
		})
		return true
	})

	return c, nil
}

func getChromiumExtURL(id, updateURL string) string {
	if strings.HasSuffix(updateURL, "clients2.google.com/service/update2/crx") {
		return "https://chrome.google.com/webstore/detail/" + id
	} else if strings.HasSuffix(updateURL, "edge.microsoft.com/extensionwebstorebase/v1/crx") {
		return "https://microsoftedge.microsoft.com/addons/detail/" + id
	}
	return ""
}

func (c *ChromiumExtension) Name() string {
	return "Extension"
}

func (c *ChromiumExtension) Len() int {
	return len(*c)
}

type FirefoxExtension []*Extension

var lang = language.Und

func (f *FirefoxExtension) Extract(_ []byte) error {
	s, err := fileutil.ReadFile(types.FirefoxExtension.TempFilename())
	if err != nil {
		return err
	}
	_ = os.Remove(types.FirefoxExtension.TempFilename())
	j := gjson.Parse(s)
	for _, v := range j.Get("addons").Array() {
		// https://searchfox.org/mozilla-central/source/toolkit/mozapps/extensions/internal/XPIDatabase.jsm#157
		if v.Get("location").String() != "app-profile" {
			continue
		}

		if lang != language.Und {
			locale := findFirefoxLocale(v.Get("locales").Array(), lang)
			*f = append(*f, &Extension{
				ID:          v.Get("id").String(),
				Enabled:     v.Get("active").Bool(),
				Name:        locale.Get("name").String(),
				Description: locale.Get("description").String(),
				Version:     v.Get("version").String(),
				HomepageURL: locale.Get("homepageURL").String(),
			})
			continue
		}

		*f = append(*f, &Extension{
			ID:          v.Get("id").String(),
			Enabled:     v.Get("active").Bool(),
			Name:        v.Get("defaultLocale.name").String(),
			Description: v.Get("defaultLocale.description").String(),
			Version:     v.Get("version").String(),
			HomepageURL: v.Get("defaultLocale.homepageURL").String(),
		})
	}
	return nil
}

func findFirefoxLocale(locales []gjson.Result, targetLang language.Tag) gjson.Result {
	tags := make([]language.Tag, 0, len(locales))
	indices := make([]int, 0, len(locales))
	for i, locale := range locales {
		for _, tagStr := range locale.Get("locales").Array() {
			tag, _ := language.Parse(tagStr.String())
			if tag == language.Und {
				continue
			}
			tags = append(tags, tag)
			indices = append(indices, i)
		}
	}
	_, tagIndex, _ := language.NewMatcher(tags).Match(targetLang)
	return locales[indices[tagIndex]]
}

func (f *FirefoxExtension) Name() string {
	return "Extension"
}

func (f *FirefoxExtension) Len() int {
	return len(*f)
}
