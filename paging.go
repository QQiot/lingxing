package lingxing

import "net/url"
import "github.com/google/go-querystring/query"

type Paging struct {
	nextOffset int `json:"-" url:"-"`           // 下一次偏移索引
	Offset     int `json:"offset" url:"offset"` // 分页偏移索引（默认0）
	Limit      int `json:"length" url:"length"` // 分页偏移长度（默认1000）
}

func (p *Paging) SetPagingVars() *Paging {
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 1000
	}
	p.nextOffset = p.Offset + p.Limit
	return p
}

// change to url.values
func toValues(i interface{}) (values url.Values) {
	values, _ = query.Values(i)
	return
}
