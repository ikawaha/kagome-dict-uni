package data

import(
	"os"
	"time"
)


func dictUnidictAqBytes() ([]byte, error) {
	return _dictUnidictAq, nil
}

func dictUnidictAq() (*asset, error) {
	bytes, err := dictUnidictAqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.aq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
