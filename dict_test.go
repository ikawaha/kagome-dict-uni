package uni

import (
	"reflect"
	"testing"

	"github.com/ikawaha/kagome/v2/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
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

func TestLoadShrink(t *testing.T) {
	d, err := dict.LoadShrink(testDictPath)
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

func TestContentsMeta(t *testing.T) {
	d := Dict()
	if want, got := 0, d.ContentsMeta[dict.POSStartIndex]; want != int(got) {
		t.Errorf("pos start index: want %d, got %d", want, got)
	}
	if want, got := 6, d.ContentsMeta[dict.POSEndIndex]; want != int(got) {
		t.Errorf("pos end index: want %d, got %d", want, got)
	}
	if want, got := 7, d.ContentsMeta[dict.BaseFormIndex]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 6, d.ContentsMeta[dict.ReadingIndex]; want != int(got) {
		t.Errorf("reading index: want %d, got %d", want, got)
	}
	if want, got := 9, d.ContentsMeta[dict.PronunciationIndex]; want != int(got) {
		t.Errorf("pronunciation index: want %d, got %d", want, got)
	}
}

func TestBaseForm(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	tokens := tnz.Tokenize("寿司を食べたい")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].BaseForm(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//食べたい->食べる
	if got, ok := tokens[3].BaseForm(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "食べる"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestReading(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Reading(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//公園->コウエン
	if got, ok := tokens[1].Reading(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "コウエン"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestPronunciation(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Pronunciation(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//公園->コウエン
	if got, ok := tokens[1].Pronunciation(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "コーエン"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestPOS(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	tokens := tnz.Tokenize("公園に行く")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got := tokens[0].POS(); len(got) > 0 {
		t.Errorf("want empty, got %+v", got)
	}
	//行く->
	if want, got := []string{"動詞", "非自立可能", "*", "*", "五段-カ行", "終止形-一般"}, tokens[3].POS(); !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
