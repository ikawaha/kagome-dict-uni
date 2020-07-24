package data

import(
	"os"
	"time"
)


func dictUnidictCiBytes() ([]byte, error) {
	return _dictUnidictCi, nil
}

func dictUnidictCi() (*asset, error) {
	bytes, err := dictUnidictCiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ci", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
