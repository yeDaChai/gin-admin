package schema

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int  `json:"total"`
	Current  uint `json:"current"`
	PageSize uint `json:"pageSize"`
}

// HTTPNewItem HTTP响应创建成功后的记录ID
type HTTPNewItem struct {
	RecordID string `json:"record_id"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageIndex uint // 页索引
	PageSize  uint // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int // 总数据条数
}
