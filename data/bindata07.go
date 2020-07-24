package data

import(
	"os"
	"time"
)


func dictUnidictAhBytes() ([]byte, error) {
	return _dictUnidictAh, nil
}

func dictUnidictAh() (*asset, error) {
	bytes, err := dictUnidictAhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ah", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
