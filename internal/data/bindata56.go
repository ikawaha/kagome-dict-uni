package data

import(
	"os"
	"time"
)


func dictUnidictCeBytes() ([]byte, error) {
	return _dictUnidictCe, nil
}

func dictUnidictCe() (*asset, error) {
	bytes, err := dictUnidictCeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ce", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
