package data

import(
	"os"
	"time"
)


func dictUnidictDcBytes() ([]byte, error) {
	return _dictUnidictDc, nil
}

func dictUnidictDc() (*asset, error) {
	bytes, err := dictUnidictDcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.dc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
