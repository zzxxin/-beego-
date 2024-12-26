package utils

import (
	"fmt"
	"html/template"
	"strings"
)

// Pagination 分页结构体
type Pagination struct {
	TotalCount  int // 总记录数
	PageSize    int // 每页显示数量
	CurrentPage int // 当前页码
	TotalPages  int // 总页数
}

// NewPagination 用于创建分页对象
func NewPagination(totalCount, pageSize, currentPage int) *Pagination {
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}
	if currentPage <= 0 {
		currentPage = 1 // 默认当前页为第1页
	}
	totalPages := (totalCount + pageSize - 1) / pageSize

	return &Pagination{
		TotalCount:  totalCount,
		PageSize:    pageSize,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}
}

// PageLinks 生成分页链接
func (p *Pagination) PageLinks() template.HTML {
	if p.TotalPages <= 1 {
		return template.HTML("") // 如果只有一页，不显示分页
	}

	var links []string

	// 生成前一页链接
	if p.CurrentPage > 1 {
		prevPage := p.CurrentPage - 1
		links = append(links, fmt.Sprintf(`<li><a href="?page=%d" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>`, prevPage))
	} else {
		// 禁用 "上一页" 链接
		links = append(links, `<li class="disabled"><span aria-hidden="true">&laquo;</span></li>`)
	}

	// 生成页码链接
	for i := 1; i <= p.TotalPages; i++ {
		if i == p.CurrentPage {
			links = append(links, fmt.Sprintf(`<li class="active"><span>%d</span></li>`, i))
		} else {
			links = append(links, fmt.Sprintf(`<li><a href="?page=%d">%d</a></li>`, i, i))
		}
	}

	// 生成下一页链接
	if p.CurrentPage < p.TotalPages {
		nextPage := p.CurrentPage + 1
		links = append(links, fmt.Sprintf(`<li><a href="?page=%d" aria-label="Next"><span aria-hidden="true">&raquo;</span></a></li>`, nextPage))
	} else {
		// 禁用 "下一页" 链接
		links = append(links, `<li class="disabled"><span aria-hidden="true">&raquo;</span></li>`)
	}

	// 使用 strings.Join 将 links 切片中的元素拼接为一个字符串
	html := strings.Join(links, "")

	// 返回 HTML 类型，避免转义
	return template.HTML(fmt.Sprintf(`<ul class="pagination">%s</ul>`, html))
}
