package data

import(
	"os"
	"time"
)


func dictUnidictBgBytes() ([]byte, error) {
	return _dictUnidictBg, nil
}

func dictUnidictBg() (*asset, error) {
	bytes, err := dictUnidictBgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
