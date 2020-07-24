package data

import(
	"os"
	"time"
)


func dictUnidictCgBytes() ([]byte, error) {
	return _dictUnidictCg, nil
}

func dictUnidictCg() (*asset, error) {
	bytes, err := dictUnidictCgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
