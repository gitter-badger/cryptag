// Steve Phillips / elimisteve
// 2015.11.14

package cryptag

import (
	"os"
	"path"
	"runtime"
)

var (
	// TrustedBasePath is where private, non-shared data (e.g., crypto
	// keys) can be safely stored.
	TrustedBasePath = path.Join(os.Getenv("HOME"), ".cryptag")

	// BackendPath is the path to the directory where backend config
	// files are stored (e.g., "/home/myusername/.cryptag/backends").
	// This can be overridden with the BACKEND_PATH
	// environment variable (rarely useful; this exists so that all
	// CrypTag backend configs could be on a USB drive).
	BackendPath = path.Join(os.Getenv("HOME"), ".cryptag", "backends")

	// DefaultLocalDataPath is the default directory where local
	// backends will be told to store their data.  This directory will
	// contain 'rows' and 'tags' subdirectories.
	DefaultLocalDataPath = path.Join(os.Getenv("HOME"), ".cryptag")

	LocalDataPath = DefaultLocalDataPath
)

func init() {
	if p := os.Getenv("BACKEND_PATH"); p != "" {
		BackendPath = p
	}

	// Change LocalDataPath if apparently on Android
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm" {
		androidStorage := "/storage/sdcard0"
		if _, err := os.Stat(androidStorage); os.IsNotExist(err) {
			LocalDataPath = androidStorage + "/.cryptag"
		}
	}

	// Change LocalDataPath if on Sandstorm (useful for servers)
	if os.Getenv("SANDSTORM") == "1" {
		LocalDataPath = "/var"
	}
}
