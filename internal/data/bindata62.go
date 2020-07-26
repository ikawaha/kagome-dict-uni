package data

import(
	"os"
	"time"
)


func dictUnidictCkBytes() ([]byte, error) {
	return _dictUnidictCk, nil
}

func dictUnidictCk() (*asset, error) {
	bytes, err := dictUnidictCkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ck", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
