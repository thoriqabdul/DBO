package request

type BasePaginateRequest struct {
	Search string `json:"search,omitempty;query:search"`
	Page   int    `json:"page,omitempty;query:page"`
	Limit  int    `json:"limit,omitempty;query:limit"`
}

func (query *BasePaginateRequest) GetOffset() int {
	return (query.Page - 1) * query.Limit
}
