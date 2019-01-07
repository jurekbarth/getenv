package getenv

import (
	"os"
	"testing"
)

type config struct {
	DownloadDir string `envconfig:"DOWNLOAD_DIR"`
	UnzipDir    string `envconfig:"UNZIP_DIR"`
}

func TestEmptyVar(t *testing.T) {
	var c config
	Process("GE_", &c)
	if c.DownloadDir != "" {
		t.Fail()
	}
}

func TestVar(t *testing.T) {
	os.Setenv("GE_DOWNLOAD_DIR", "downloaddir")
	var c config
	Process("GE_", &c)
	if c.DownloadDir != "downloaddir" {
		t.Fail()
	}
}

func TestVars(t *testing.T) {
	os.Setenv("GE_DOWNLOAD_DIR", "downloaddir")
	os.Setenv("GE_UNZIP_DIR", "unzipdir")
	var c config
	Process("GE_", &c)
	if c.DownloadDir != "downloaddir" {
		t.Fail()
	}
	if c.UnzipDir != "unzipdir" {
		t.Fail()
	}
}
