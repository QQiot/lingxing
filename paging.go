package lingxing

type Paging struct {
	Offset     int `json:"offset"` // 分页偏移索引（默认0）
	NextOffset int `json:"-"`      // 下一次偏移索引
	Limit      int `json:"length"` // 分页偏移长度（默认1000）
}

func (p *Paging) SetPagingVars() *Paging {
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.Limit <= 0 {
		p.Limit = 1000
	}
	p.NextOffset = p.Offset + p.Limit
	return p
}
