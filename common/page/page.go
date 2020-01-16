package page

import (
	"errors"
	"fmt"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
	OrderAsc        = "asc"
	OrderDesc       = "desc"
)

type Page struct {
	PageSize    int           `json:"pageSize"`    //每页显示的条数
	RecordCount int           `json:"recordCount"` //记录数
	CurrentPage int           `json:"currentPage"` //当前页
	PageNum     int           `json:"pageNum"`     //页数
	InferList   []interface{} `json:"inferList"`   //记录信息
}

func (p *Page) getPageCount() int {
	num := p.RecordCount / p.PageSize
	mod := p.RecordCount % p.PageSize
	if mod != 0 {
		num++
	}
	if p.RecordCount == 0 {
		return p.RecordCount
	}
	return num
}

func (p *Page) calcPageNum() {
	p.PageNum = p.getPageCount()
}

// 初始化分页
func (p *Page) SetPageInfo(currentPage, pageSize, recordCount int, inferList []interface{}) {
	p.RecordCount = recordCount
	p.CurrentPage = currentPage
	p.PageSize = pageSize
	p.PageNum = p.getPageCount()
	p.InferList = inferList
}

// 分页接口
// 提供分页的公共接口, 只需要传入已经排好序的列表, 排序方式, 页码, 分页大小
func PagingList(inferList []interface{}, order string, pageIndex, pageSize int) (Page, error) {
	paging := Page{}
	if inferList == nil || len(inferList) == 0 {
		paging.SetPageInfo(pageIndex, pageSize, 0, make([]interface{}, 0))
		return paging, nil
	}

	var fromIndex, toIndex int
	inferListLen := len(inferList)
	switch order {
	case OrderAsc:
		fromIndex = (pageIndex - 1) * pageSize
		toIndex = fromIndex + pageSize
	case OrderDesc:
		toIndex = inferListLen - ((pageIndex - 1) * pageSize)
		fromIndex = toIndex - pageSize
	default:
		return paging, errors.New(fmt.Sprint("Invalid order"))
	}
	if fromIndex < 0 {
		fromIndex = 0
	} else if fromIndex > inferListLen {
		fromIndex = inferListLen
	}

	if toIndex < 0 {
		toIndex = 0
	} else if toIndex > inferListLen {
		toIndex = inferListLen
	}

	// 截取列表
	inferListNew := inferList[fromIndex:toIndex]
	// 如果是倒叙, 反转列表
	if order == OrderDesc {
		reverseList(inferListNew)
	}
	paging.SetPageInfo(pageIndex, pageSize, inferListLen, inferListNew)
	return paging, nil
}

// 反转接口
func reverseList(revList []interface{}) {
	if revList == nil || len(revList) == 0 {
		return
	}
	for i, j := 0, len(revList)-1; i < j; i, j = i+1, j-1 {
		revList[i], revList[j] = revList[j], revList[i]
	}
}
