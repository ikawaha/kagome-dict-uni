package data

import(
	"os"
	"time"
)


func dictUnidictDaBytes() ([]byte, error) {
	return _dictUnidictDa, nil
}

func dictUnidictDa() (*asset, error) {
	bytes, err := dictUnidictDaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.da", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
