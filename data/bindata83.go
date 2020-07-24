package data

import(
	"os"
	"time"
)


func dictUnidictDfBytes() ([]byte, error) {
	return _dictUnidictDf, nil
}

func dictUnidictDf() (*asset, error) {
	bytes, err := dictUnidictDfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.df", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595602641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
