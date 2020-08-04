package uni

import (
	"archive/zip"
	"bytes"
	"fmt"
	"sort"
	"sync"

	data "github.com/ikawaha/kagome-dict-uni/internal/data"
	"github.com/ikawaha/kagome/v2/dict"
)

type FeatureIndex int

const (
	// Features are information given to a word, such as follows:
	// 公園	名詞,普通名詞,一般,*,*,*,コウエン,公園,公園,コーエン,公園,コーエン,漢,*,*,*,*
	// に	助詞,格助詞,*,*,*,*,ニ,に,に,ニ,に,ニ,和,*,*,*,*
	// 行っ	動詞,非自立可能,*,*,五段-カ行,連用形-促音便,イク,行く,行っ,イッ,行く,イク,和,*,*,*,*
	// た	助動詞,*,*,*,助動詞-タ,終止形-一般,タ,た,た,タ,た,タ,和,*,*,*,*
	// EOS
	// POSHierarchy represents part-of-speech hierarchy
	// e.g. Columns 動詞,非自立可能,*,* are POSs which hierarchy depth is 4.
	POSHierarchy = 4
	// InflectionalType represents 活用型 (e.g. 五段-カ行)
	InflectionalType FeatureIndex = 4
	// InflectionalForm represents 活用形 (e.g. 連用形-促音便)
	InflectionalForm = 5
	// Reading represents 読み (e.g. コウエン)
	Reading = 6
	// BaseForm represents 基本形 (e.g. 行く)
	BaseForm = 7
	// Pronunciation represents 発音 (e.g. コーエン)
	Pronunciation = 11
)

type systemDict struct {
	once sync.Once
	dict *dict.Dict
}

var (
	full   systemDict
	shrink systemDict
)

// Dict returns a dictionary.
func Dict() *dict.Dict {
	full.once.Do(func() {
		full.dict = loadDict(true)
		shrink.once.Do(func() {
			shrink.dict = full.dict
		})
	})
	return full.dict
}

// DictShrink returns a dictionary without content part.
// note. If an unshrinked dictionary already exists, this function returns it.
func DictShrink() *dict.Dict {
	shrink.once.Do(func() {
		shrink.dict = loadDict(false)
	})
	return shrink.dict
}

func loadDict(full bool) (d *dict.Dict) {
	pieces := data.AssetNames()
	sort.Strings(pieces)

	rs := make([]SizeReaderAt, 0, len(pieces))
	for _, v := range pieces {
		b, err := data.Asset(v)
		if err != nil {
			panic(fmt.Errorf("assert error, %q, %v", v, err))
		}
		rs = append(rs, bytes.NewReader(b))
	}
	r := NewMultiSizeReaderAt(rs...)
	zr, err := zip.NewReader(r, r.Size())
	if err != nil {
		panic(err)
	}
	d, err = dict.Load(zr, full)
	if err != nil {
		panic(err)
	}
	return d
}
