package data

import(
	"os"
	"time"
)


func dictUnidictAkBytes() ([]byte, error) {
	return _dictUnidictAk, nil
}

func dictUnidictAk() (*asset, error) {
	bytes, err := dictUnidictAkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ak", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
