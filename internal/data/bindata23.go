package data

import(
	"os"
	"time"
)


func dictUnidictAxBytes() ([]byte, error) {
	return _dictUnidictAx, nil
}

func dictUnidictAx() (*asset, error) {
	bytes, err := dictUnidictAxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ax", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
