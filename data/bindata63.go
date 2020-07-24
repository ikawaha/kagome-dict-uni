package data

import(
	"os"
	"time"
)


func dictUnidictClBytes() ([]byte, error) {
	return _dictUnidictCl, nil
}

func dictUnidictCl() (*asset, error) {
	bytes, err := dictUnidictClBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
