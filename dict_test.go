package uni

import (
	"testing"

	"github.com/ikawaha/kagome/v2/dict"
)

const (
	UniDictEntrySize = 756466
	testDictPath     = "testdata/uni.dict"
)


func TestDictShrink(t *testing.T) {
	d := DictShrink()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLoadShurnk(t *testing.T) {
	d, err := dict.LoadShurink(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNew(t *testing.T) {
	d := Dict()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLoadDictFile(t *testing.T) {
	d, err := dict.LoadDictFile(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSingleton(t *testing.T) {
	a := Dict()
	b := Dict()
	if a != b {
		t.Errorf("got %p and %p, expected singleton", a, b)
	}
}
