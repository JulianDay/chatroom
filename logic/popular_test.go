package logic

import (
	"testing"
	"time"
)

func TestPopularWords(t *testing.T) {
	popular := NewPopularWordMgr()
	popular.AddWords([]string{"a", "b", "c"})
	popular.AddWords([]string{"a", "b"})
	popular.AddWords([]string{"a"})
	words := popular.GetPopularWords()
	t.Log(words.String())
	time.Sleep(time.Second)
	popular.AddWords([]string{"a", "b"})
	words = popular.GetPopularWords()
	t.Log(words)
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		popular.AddWords([]string{"c"})
		words = popular.GetPopularWords()
		t.Log(words)
		time.Sleep(time.Second)
	}
}
