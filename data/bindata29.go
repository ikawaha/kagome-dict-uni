package data

import(
	"os"
	"time"
)


func dictUnidictBdBytes() ([]byte, error) {
	return _dictUnidictBd, nil
}

func dictUnidictBd() (*asset, error) {
	bytes, err := dictUnidictBdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
