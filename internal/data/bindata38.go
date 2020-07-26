package data

import(
	"os"
	"time"
)


func dictUnidictBmBytes() ([]byte, error) {
	return _dictUnidictBm, nil
}

func dictUnidictBm() (*asset, error) {
	bytes, err := dictUnidictBmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
