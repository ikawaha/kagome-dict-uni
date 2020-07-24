package data

import(
	"os"
	"time"
)


func dictUnidictAvBytes() ([]byte, error) {
	return _dictUnidictAv, nil
}

func dictUnidictAv() (*asset, error) {
	bytes, err := dictUnidictAvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.av", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
