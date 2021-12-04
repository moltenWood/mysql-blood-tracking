package controllers

// 现目前仅针对 sql的查询语句
import (
	"dbScheduleAnalysis/conf"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	TableName VariableName
	FieldName VariableName // 可能为 “”
}

// FindTargetPositions 在arrayStr中查找所有targetStr所在的下表列表
// arrayStr: 需要被查找的范围
// targetStr: 需要查找的字符串
// 返回：下表
func FindTargetPositions(arrayStr string, targetStr string) []int {
	return []int{}
}

// ParseSqlSentence 将sqlSentence解析成 [struct(关键字，关键字后面的内容),...]
func ParseSqlSentence(sqlSentence string) []KeyWordAndContent {
	var elements = strings.Split(sqlSentence, " ")
	var keyWordStructs = getKeyWordsStructs(elements)
	return keyWordStructs
}

type KeyWordAndContent struct {
	keyword string
	content []string
}

func getKeyWordsStructs(elements []string) []KeyWordAndContent {
	var keyWordStructs []KeyWordAndContent
	var currenKeyWord KeyWordAndContent
	for i := range elements {
		var isSqlKeyWord = conf.IsSqlKeyWord(elements[i])
		if isSqlKeyWord {
			if currenKeyWord.keyword != "" {
				keyWordStructs = append(keyWordStructs, currenKeyWord)
				currenKeyWord = KeyWordAndContent{}
			}
			currenKeyWord.keyword = elements[i]
		} else {
			currenKeyWord.content = append(currenKeyWord.content, elements[i])
		}
	}
	keyWordStructs = append(keyWordStructs, currenKeyWord)
	return keyWordStructs
}

type VariableName struct {
	Content string
	Type    conf.TargetObjectType
}

type RenameInfo struct {
}

// RenameMap 重命名map
// key: 变量名
// value: 数据库中的实际表名或者字段名
type RenameMap map[VariableName]string

//// ArrayTree 域的层级关系
//type ArrayTree struct {
//	ArrayName  string  // 这个括号域 可能被重命名为ArrayName 否则为 “”
//	KeyPart    KeyPart // 这个不为零值 时，ArrayName 和 ChildArray一定是零值，反之这个为零值， 没有对内容做进一步解析
//	ChildArray []ArrayTree
//}

var InnerBracketId int // 第几个内部括号组，全局id

type RenameThis struct {
	innerBracketId     int
	keyAndContentSlice []KeyWordAndContent
}

type Collector struct {
	renameThisSLice []RenameThis
}

// GetArrayTree 根据括号的包裹情况，分割不同层的域
func GetArrayTree(sqlSentence string, innerBracketId *int, collector *Collector) {
	var bracketElements = GetBracketArray(sqlSentence)
	fmt.Printf("bracketElements: %v \n", bracketElements)
	for i := 0; i < len(sqlSentence); i++ {
		if sqlSentence[i] == '(' {
			var bracketsInfo = GetTheOtherBracketIndex(bracketElements, i)
			// 待进一步处理的括号内部内容
			var innerSqlSentence = sqlSentence[i+1 : bracketsInfo["rightBracketIndex"]]
			//todo
			GetArrayTree(innerSqlSentence, innerBracketId, collector)
			// 去除这部分括号内部内容
			sqlSentence = replaceBracketArray(sqlSentence, i, bracketsInfo["rightBracketIndex"], *innerBracketId)
			ShowIndex(sqlSentence)
			bracketElements = updateBracketElements(bracketElements, i, bracketsInfo["rightBracketIndex"], *innerBracketId)
			fmt.Printf("-->bracketElements: %v \n", bracketElements)

			var keyAndContentSlice = ParseSqlSentence(sqlSentence)
			var renameThis = RenameThis{*innerBracketId, keyAndContentSlice}
			collector.renameThisSLice = append(collector.renameThisSLice, renameThis)
			*innerBracketId++
			i++ // 刚好在被去除后的右括号上
		}
	}
}

func ShowIndex(sql string) {
	for i, v := range []byte(sql) {
		fmt.Printf("[%v]->[%v] ", i, string(v))
	}
	fmt.Println("")
}

func updateBracketElements(bracketElements []BracketElement, leftBracketIndex int, rightBracketIndex int, innerBracketId int) []BracketElement {
	var oldLength = rightBracketIndex - leftBracketIndex + 1
	var newLength = len(strconv.Itoa(innerBracketId)) + 2
	var itIsTime = false
	for i := range bracketElements {
		if !itIsTime {
			if bracketElements[i].leftBracketIndex == leftBracketIndex {
				bracketElements[i].rightBracketIndex += newLength - oldLength
				itIsTime = true
			}
		} else {
			bracketElements[i].leftBracketIndex += newLength - oldLength
			bracketElements[i].rightBracketIndex += newLength - oldLength
		}
	}
	return bracketElements
}

func ShowCollector(collector *Collector) {
	for i := range collector.renameThisSLice {
		fmt.Printf("%v\n", collector.renameThisSLice[i])
	}
}

func showBrackets(sqlSentence string, brackets []BracketElement) {
	var byteSqlSentence = []byte(sqlSentence)

	for i := range brackets {
		fmt.Println(string(byteSqlSentence[brackets[i].leftBracketIndex : brackets[i].rightBracketIndex+1]))
	}
}

type BracketElement struct {
	leftBracketIndex  int
	rightBracketIndex int
}

func GetBracketArray(sqlSentence string) []BracketElement {
	var brackets []BracketElement
	for i := len(sqlSentence) - 1; i >= 0; i-- {
		if sqlSentence[i] == '(' {
			var oneBracketElement BracketElement
			oneBracketElement.leftBracketIndex = i
			for j := i; j < len(sqlSentence); j++ {
				if sqlSentence[j] == ')' && !rightBracketHasBeenUsed(brackets, j) {
					oneBracketElement.rightBracketIndex = j
					break
				}
			}
			brackets = append(brackets, oneBracketElement)
		}
	}
	var reversedBrackets = make([]BracketElement, 0, len(brackets))
	for i := len(brackets) - 1; i >= 0; i-- {
		reversedBrackets = append(reversedBrackets, brackets[i])
	}
	return reversedBrackets
}

func rightBracketHasBeenUsed(brackets []BracketElement, index int) bool {
	var HasBeenUsed = false
	for i := range brackets {
		if brackets[i].rightBracketIndex == index {
			HasBeenUsed = true
			break
		}
	}
	return HasBeenUsed
}

func GetTheOtherBracketIndex(elements []BracketElement, index int) map[string]int {
	for i := range elements {
		if elements[i].leftBracketIndex == index {
			return map[string]int{"leftBracketIndex": index, "rightBracketIndex": elements[i].rightBracketIndex}
		} else if elements[i].rightBracketIndex == index {
			return map[string]int{"leftBracketIndex": elements[i].leftBracketIndex, "rightBracketIndex": index}
		}
	}
	return map[string]int{"leftBracketIndex": -1, "rightBracketIndex": -1}
}

// index: 第几个内部括号组
func replaceBracketArray(sqlSentence string, leftBracketIndex int, rightBracketIndex int, index int) string {
	return sqlSentence[:leftBracketIndex] + "(" + strconv.Itoa(index) + ")" + sqlSentence[rightBracketIndex+1:]
}
