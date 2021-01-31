package dfa

import (
	"strings"
	"sync"
)

type MATCHTYPE int

const (
	ALL MATCHTYPE = iota
	SINGLE
	INVALID_WORDS           = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,！,。,；,：,’,{,},【,】,[,],、"
	SENSITIVE_CHILDRED_SIZE = 256
)

var InvalidWords = make(map[string]struct{})
var Util *DFAUtil

type sensitiveNode struct {
	isEnd    bool
	children map[rune]*sensitiveNode
}

func newSensitiveNode() *sensitiveNode {
	return &sensitiveNode{
		children: make(map[rune]*sensitiveNode, SENSITIVE_CHILDRED_SIZE),
	}
}

type DFAUtil struct {
	root *sensitiveNode
	mu   sync.Mutex
}

func NewDFAUtil() *DFAUtil {
	return Util
}

// 初始化屏蔽字树
func Init(words []string) {
	invalidArr := strings.Split(INVALID_WORDS, ",")
	for _, v := range invalidArr {
		InvalidWords[v] = struct{}{}
	}

	dfaUtil := &DFAUtil{
		root: newSensitiveNode(),
	}
	for _, word := range words {
		sensitiveRune := []rune(word)
		if len(sensitiveRune) > 1 {
			dfaUtil.AddWord(sensitiveRune)
		}
	}

	Util = dfaUtil
}

func (dfaUtil *DFAUtil) AddWord(word []rune) {
	if dfaUtil.root == nil {
		return
	}
	dfaUtil.mu.Lock()
	defer dfaUtil.mu.Unlock()

	currNode := dfaUtil.root
	for _, single := range word {
		if targetNode, exist := currNode.children[single]; !exist {
			targetNode = newSensitiveNode()
			currNode.children[single] = targetNode
			currNode = targetNode
		} else {
			currNode = targetNode
		}
	}

	currNode.isEnd = true
}

func (dfaUtil *DFAUtil) Contains(sentence string) bool {
	var flag = false
	var matchFlag = 0
	sentenceRune := []rune(sentence)
	currNode := dfaUtil.root
	length := len(sentenceRune)

	for i := 0; i < length; i++ {
		if _, exist := InvalidWords[string(sentenceRune[i])]; exist {
			continue
		}

		if targetNode, exist := currNode.children[sentenceRune[i]]; exist {
			matchFlag++
			currNode = targetNode
			if currNode.isEnd {
				flag = true
				break
			}
		} else {
			currNode = dfaUtil.root
		}
	}
	if matchFlag < 2 || !flag {
		return false
	}
	return true
}

func (dfaUtil *DFAUtil) SearchSensitive(sentence string, matchType MATCHTYPE) (matchIndexList []*matchIndex) {
	sentenceRune := []rune(sentence)
	currNode := dfaUtil.root
	tag, start := -1, -1
	length := len(sentenceRune)

	for i := 0; i < length; i++ {
		if _, exist := InvalidWords[string(sentenceRune[i])]; exist {
			continue
		}

		if targetNode, exist := currNode.children[sentenceRune[i]]; exist {
			tag++
			if tag == 0 {
				start = i
			}

			currNode = targetNode
			if currNode.isEnd {
				matchIndexList = append(matchIndexList, newMatchIndex(start, i))
				if matchType == SINGLE {
					return matchIndexList
				}

				//重新回到树的顶部,找下一个敏感词
				currNode = dfaUtil.root
				tag, start = -1, -1
			}
		} else {
			if start != -1 {
				i = start
			}

			currNode = dfaUtil.root
			tag, start = -1, -1
		}
	}

	return matchIndexList
}

func (dfaUtil *DFAUtil) Cover(sentence string, mask rune) (string, bool) {
	matchIndexList := dfaUtil.SearchSensitive(sentence, ALL)
	if len(matchIndexList) == 0 {
		return sentence, false
	}

	sentenceRune := []rune(sentence)
	for _, matchIndexStruct := range matchIndexList {
		for i := matchIndexStruct.start; i <= matchIndexStruct.end; i++ {
			sentenceRune[i] = mask
		}
	}

	return string(sentenceRune), true
}

type matchIndex struct {
	start int
	end   int
}

func newMatchIndex(start, end int) *matchIndex {
	return &matchIndex{
		start: start,
		end:   end,
	}
}
