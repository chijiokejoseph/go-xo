package files

import "errors"

var ErrNotExist = errors.New("file does not exist")
var ErrMarshal = errors.New("json marshaling to convert Record object to json failed")
var ErrUnMarshal = errors.New("json unmarshalling for target file failed")
var ErrNoWrite = errors.New("failed to write data to file")
var ErrNoRead = errors.New("failed to read data from file")
var ErrNoCreate = errors.New("failed to open file with create flags")