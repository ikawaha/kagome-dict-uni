package data

import(
	"os"
	"time"
)


func dictUnidictAdBytes() ([]byte, error) {
	return _dictUnidictAd, nil
}

func dictUnidictAd() (*asset, error) {
	bytes, err := dictUnidictAdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ad", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
