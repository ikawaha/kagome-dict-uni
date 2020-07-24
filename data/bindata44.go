package data

import(
	"os"
	"time"
)


func dictUnidictBsBytes() ([]byte, error) {
	return _dictUnidictBs, nil
}

func dictUnidictBs() (*asset, error) {
	bytes, err := dictUnidictBsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
