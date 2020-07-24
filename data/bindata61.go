package data

import(
	"os"
	"time"
)


func dictUnidictCjBytes() ([]byte, error) {
	return _dictUnidictCj, nil
}

func dictUnidictCj() (*asset, error) {
	bytes, err := dictUnidictCjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
