package data

import(
	"os"
	"time"
)


func dictUnidictDdBytes() ([]byte, error) {
	return _dictUnidictDd, nil
}

func dictUnidictDd() (*asset, error) {
	bytes, err := dictUnidictDdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.dd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
