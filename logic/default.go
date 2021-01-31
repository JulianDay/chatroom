package logic

import (
	"errors"
)

var (
	ErrDuplicateName = errors.New("duplicate Name")
)

var Room = NewRoom()
var LastMessage = newLastMessage(50)
var Popular = NewPopularWordMgr()
