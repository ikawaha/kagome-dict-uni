package data

import(
	"os"
	"time"
)


func dictUnidictBoBytes() ([]byte, error) {
	return _dictUnidictBo, nil
}

func dictUnidictBo() (*asset, error) {
	bytes, err := dictUnidictBoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595765650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
