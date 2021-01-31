package logic

import (
	"chatroom/dfa"
)

func FilterSensitive(content string) string {
	dfa := dfa.NewDFAUtil()
	if dfa == nil {
		return content
	}
	content, _ = dfa.Cover(content, '*')
	return content
}
