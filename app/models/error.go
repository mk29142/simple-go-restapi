package models

import "errors"

var (
    ErrNotFound = errors.New("Requested item is not found!")
)
