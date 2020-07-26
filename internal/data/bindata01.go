package data

import(
	"os"
	"time"
)


func dictUnidictAbBytes() ([]byte, error) {
	return _dictUnidictAb, nil
}

func dictUnidictAb() (*asset, error) {
	bytes, err := dictUnidictAbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ab", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
