package main

import "strings"

//PathSeparator パスのスラッシュ
const PathSeparator = "/"

//Path Path型文字列解析
type Path struct {
	Path string
	ID   string
}

//NewPath パスの分割
func NewPath(p string) *Path {
	var id string
	p = strings.Trim(p, PathSeparator)
	s := strings.Split(p, PathSeparator)
	if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], PathSeparator)
	}
	return &Path{Path: p, ID: id}
}

//HasID IDを付ける
func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
