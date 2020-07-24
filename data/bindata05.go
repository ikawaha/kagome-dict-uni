package data

import(
	"os"
	"time"
)


func dictUnidictAfBytes() ([]byte, error) {
	return _dictUnidictAf, nil
}

func dictUnidictAf() (*asset, error) {
	bytes, err := dictUnidictAfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.af", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
