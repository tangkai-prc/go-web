package dp

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

// 正则匹配
type Pattern struct {
	matched bool   // 是否匹配的全局值
	pattern []rune // 正则串
}

func NewPattern(expr string) *Pattern {
	return &Pattern{pattern: []rune(expr)}
}

func (p *Pattern) Match(text string) bool { // 文本串及长度
	p.matched = false
	p.rmatch(0, 0, []rune(text))
	return p.matched
}

/*
只有 * 和 ? 两种通配符
* 表示任意多(>=0)个字符
? 表示0个或1任意字符

使用回溯的方式实现

ti 表示待匹配字符的指针
pj 表示正则表达式指针位置
*/

func (p *Pattern) rmatch(ti, pj int, text []rune) {
	if p.matched { // 已经匹配
		return
	}
	if pj == len(p.pattern) { // 正则串到结束
		if ti == len(text) {
			// 文本串也结束 找到匹配
			p.matched = true
		}
		return
	}
	// -----递归出口分割线-----
	switch {
	case p.pattern[pj] == '*': // 匹配多个字符
		for k := 0; k < len(text)-ti; k++ {
			p.rmatch(ti+k, pj+1, text)
		}
	case p.pattern[pj] == '?': //匹配0个或1个
		p.rmatch(ti, pj+1, text)
		p.rmatch(ti+1, pj+1, text)
	case ti < len(text) && p.pattern[pj] == text[ti]: // 纯字符匹配才能继续,也属于剪枝操作
		p.rmatch(ti+1, pj+1, text)
	}
}

// 使用动态规划的方式怎么写?

func TestMatch(t *testing.T) {
	pattern := NewPattern("c*htm?")
	assert.Equal(t, pattern.Match("chtml"), true)
	assert.Equal(t, pattern.Match("cxxxxhtml"), true)
	assert.Equal(t, pattern.Match("cxxxxhtmlc"), false)

}
