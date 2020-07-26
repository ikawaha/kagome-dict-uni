package data

import(
	"os"
	"time"
)


func dictUnidictBuBytes() ([]byte, error) {
	return _dictUnidictBu, nil
}

func dictUnidictBu() (*asset, error) {
	bytes, err := dictUnidictBuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
