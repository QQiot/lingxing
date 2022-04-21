package lingxing

type Paging struct {
	Offset     int `json:"offset,omitempty"` // 分页偏移索引（默认0）
	NextOffset int `json:"-"`                // 下一次偏移索引
	Limit      int `json:"length,omitempty"` // 分页偏移长度（默认1000）
}

func (p *Paging) SetPagingVars(offset, limit, maxLimit int) *Paging {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}
	p.Offset = offset
	p.Limit = limit
	p.NextOffset = p.Offset + p.Limit
	return p
}
