package errorlib

import "errors"

var InternalError = errors.New("some internal error happened")
var ConfigValueMissError = errors.New("config value missed")
