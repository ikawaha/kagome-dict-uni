package data

import(
	"os"
	"time"
)


func dictUnidictAeBytes() ([]byte, error) {
	return _dictUnidictAe, nil
}

func dictUnidictAe() (*asset, error) {
	bytes, err := dictUnidictAeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ae", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
