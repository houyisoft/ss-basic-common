package constant

import (
	"github.com/kataras/iris"
)

type Pager struct {
	Rows      []map[string]string `json:"rows"`       //记录结果集
	Page      int                 `json:"page"`       //当前页数
	Pages int                     `json:"pages"` //总计页数
	Total int                     `json:"total"` //总计记录数
	PageSize  int                 `json:"pageSize"`  //每页记录数
}


//  bootstraptable 分页参数
type Pagination struct {
	PageNum  int   `json:"pageNum"`  //当前看的是第几页
	PageSize    int   `json:"pageSize"`    //每页显示多少条数据
	Total int64 `json:"total"` //总条数
	Pages   int64 `json:"pages"`   //总页数

	// 用于分页设置的参数
	Start int `json:"_"`
	Limit int `json:"_"`

	SortName  string `json:"_"` //用于指定的排序
	SortOrder string `json:"_"` // desc或asc

	// 时间范围
	StartDate string `json:"_"`
	EndDate   string `json:"_"`

	Keyword string `json:"keyword"` // 关键字  用于模糊查询

	//Uid string // 公用的特殊参数
}

func NewPagination(ctx iris.Context) (*Pagination, error) {
	pageNum, _ := ctx.URLParamInt("pageNum")
	pageSize, _ := ctx.URLParamInt("pageSize")
	sortName := ctx.URLParam("sortName")
	sortOrder := ctx.URLParam("sortOrder")
	keyword := ctx.URLParam("keyword")
	//if err1 != nil || err2 != nil {
	//	return nil, errors.New("请求的分页参数解析错误")
	//}

	page := Pagination{
		PageNum: pageNum,
		PageSize:   pageSize,
		SortName:   sortName,
		SortOrder:  sortOrder,
		Keyword:    keyword,
	}
	page.PageSetting()
	return &page, nil
}

// 设置分页参数
func (p *Pagination) PageSetting() {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}

	p.Start = (p.PageNum - 1) * p.PageSize
	p.Limit = p.PageSize
}

type PageList struct {
	Page  *Pagination `json:"page"`
	Datas interface{} `json:"datas"`
}

func PageSetVal(page *Pagination, totalNum int64, data interface{}) (pageList *PageList) {
	page.Total = totalNum
	if totalNum == 0 {
		page.Pages = 0
	} else {
		page.Pages = (totalNum-1)/int64(page.PageSize) + 1
	}
	pageList = &PageList{
		Page:  page,
		Datas: data,
	}
	return pageList
}
