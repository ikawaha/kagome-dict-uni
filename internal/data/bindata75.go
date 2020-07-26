package data

import(
	"os"
	"time"
)


func dictUnidictCxBytes() ([]byte, error) {
	return _dictUnidictCx, nil
}

func dictUnidictCx() (*asset, error) {
	bytes, err := dictUnidictCxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
