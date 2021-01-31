package logic

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type PopularWord struct {
	Word  string
	Count int
}
type PopularWordList []PopularWord

func (p PopularWordList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PopularWordList) Len() int           { return len(p) }
func (p PopularWordList) Less(i, j int) bool { return p[i].Count > p[j].Count }
func (p PopularWordList) String() (ret string) {
	var b strings.Builder
	for i, w := range p {
		b.WriteString(fmt.Sprintf("[%d] %v:%d\n", (i + 1), w.Word, w.Count))
	}
	return b.String()
}

type wordsCount struct {
	count map[string]int //{word,count}
}

func (wc *wordsCount) addWords(words []string) {
	for _, word := range words {
		wc.count[word]++
	}
}

type popularWordMgr struct {
	mu            sync.Mutex
	word          map[int64]*wordsCount //{time,wordsCount}
	accessSeconds int64
}

func NewPopularWordMgr() *popularWordMgr {
	return &popularWordMgr{
		word: map[int64]*wordsCount{},
	}
}
func (pw *popularWordMgr) AddWords(words []string) {
	if len(words) == 0 {
		return
	}
	pw.mu.Lock()
	defer pw.mu.Unlock()
	pw.expireWords()
	seconds := time.Now().Unix()
	if wc, ok := pw.word[seconds]; ok {
		wc.addWords(words)
	} else {
		wc := &wordsCount{
			count: make(map[string]int),
		}
		wc.addWords(words)
		pw.word[seconds] = wc
	}
}

func (pw *popularWordMgr) expireWords() {
	seconds := time.Now().Unix()
	if pw.accessSeconds == seconds {
		return
	}
	pw.accessSeconds = seconds
	//每秒触发一次
	expireSeconds := seconds - 5
	for key, _ := range pw.word {
		if key < expireSeconds {
			delete(pw.word, key)
		}
	}
}
func (pw *popularWordMgr) GetPopularWords() PopularWordList {
	pw.mu.Lock()
	pw.expireWords()
	allWordCount := make(map[string]int)
	for _, value := range pw.word {
		for word, count := range value.count {
			allWordCount[word] += count
		}
	}
	pw.mu.Unlock()
	p := make(PopularWordList, len(allWordCount))
	i := 0
	for k, v := range allWordCount {
		p[i] = PopularWord{Word: k, Count: v}
		i++
	}
	sort.Sort(p)
	return p
}
