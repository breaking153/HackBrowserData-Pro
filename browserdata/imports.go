// Package browserdata is responsible for initializing all the necessary
// components that handle different types of browser data extraction.
// This file, imports.go, is specifically used to import various data
// handler packages to ensure their initialization logic is executed.
// These imports are crucial as they trigger the `init()` functions
// within each package, which typically handle registration of their
// specific data handlers to a central registry.
package browserdata

import (
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/bookmark"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/cookie"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/creditcard"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/download"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/extension"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/history"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/localstorage"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/password"
	_ "github.com/breaking153/HackBrowserData-Pro/browserdata/sessionstorage"
)
