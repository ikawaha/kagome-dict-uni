package data

import(
	"os"
	"time"
)


func dictUnidictBqBytes() ([]byte, error) {
	return _dictUnidictBq, nil
}

func dictUnidictBq() (*asset, error) {
	bytes, err := dictUnidictBqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
