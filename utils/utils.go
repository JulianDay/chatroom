package utils

import "regexp"

//分词
func SplitWords(str string) (ret []string) {
	compile, err := regexp.Compile(`\S+`)
	if err != nil {
		panic(err)
	}

	result := compile.FindAllStringSubmatch(str, -1)
	for _, v := range result {
		ret = append(ret, v...)
	}
	return ret
}
