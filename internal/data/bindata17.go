package data

import(
	"os"
	"time"
)


func dictUnidictArBytes() ([]byte, error) {
	return _dictUnidictAr, nil
}

func dictUnidictAr() (*asset, error) {
	bytes, err := dictUnidictArBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ar", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765649, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
