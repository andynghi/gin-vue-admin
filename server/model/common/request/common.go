package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // Page number
	PageSize int    `json:"pageSize" form:"pageSize"` // Size of each page
	Keyword  string `json:"keyword" form:"keyword"`   //Keyword
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // Primary key ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // Role ID
}

type Empty struct{}
