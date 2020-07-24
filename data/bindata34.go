package data

import(
	"os"
	"time"
)


func dictUnidictBiBytes() ([]byte, error) {
	return _dictUnidictBi, nil
}

func dictUnidictBi() (*asset, error) {
	bytes, err := dictUnidictBiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
