package filesystem

import "errors"

var (
	ErrValueNoExist = errors.New("Value no exist.")
	ErrPathNotDirectory = errors.New("The path is not directory.")
)
