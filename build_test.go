package magelib

import (
	"log"
	"os"
	"testing"
)

const (
	binaryName = "sample"
)

var (
	kinds = []*BuildPair{
		NewBuildPair("linux", "386"),
		NewBuildPair("linux", "amd64"),
		NewBuildPair("linux", "arm"),
		NewBuildPair("darwin", "386"),
		NewBuildPair("darwin", "amd64"),
		NewBuildPair("darwin", "arm"),
	}
)

func TestBuild(t *testing.T) {
	if err := Build(binaryName, kinds); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(binaryName); err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(binaryName); err != nil {
		log.Fatal(err)
	}
}
