package config

import (
	"testing"
	"strings"
)

func TestGlobalDir(t *testing.T) {
	dir := GlobalDir()
	if !strings.HasSuffix(dir, ".nanobox") {
		t.Errorf("missing nanobox suffix")
	}
}

func TestLocalDir(t *testing.T) {
	dir := LocalDir()
	if !strings.Contains(dir, "nanobox/util/config") {
		t.Errorf("missing nanobox suffix")
	}
}

func TestLocalDirName(t *testing.T) {
	if LocalDirName() != "config" {
		t.Errorf("local dir name mismatch")
	}
}

func TestBinDir(t *testing.T) {
	dir := BinDir()
	if !strings.HasSuffix(dir, ".nanobox/bin") {
		t.Errorf("bin dir failure")
	}
}