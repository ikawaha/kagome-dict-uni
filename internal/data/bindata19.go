package data

import(
	"os"
	"time"
)


func dictUnidictAtBytes() ([]byte, error) {
	return _dictUnidictAt, nil
}

func dictUnidictAt() (*asset, error) {
	bytes, err := dictUnidictAtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.at", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
