package customErrors

import "errors"

var ErrNoMatches = errors.New("no matches found")
var ErrNoMatch = errors.New("no match found")
