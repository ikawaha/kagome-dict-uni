package data

import(
	"os"
	"time"
)


func dictUnidictCzBytes() ([]byte, error) {
	return _dictUnidictCz, nil
}

func dictUnidictCz() (*asset, error) {
	bytes, err := dictUnidictCzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
