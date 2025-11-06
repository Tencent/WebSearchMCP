package prosearch

// CloudAPIRequest 云api搜索单条
type CloudAPIRequest struct {
	SecretID  string `json:"-"`
	SecretKey string `json:"-"`
	Query     string `json:"Query"`
	Mode      int    `json:"Mode"`
	Cnt       int    `json:"Cnt,omitempty"`
	Site      string `json:"Site,omitempty"`
	FromTime  int64  `json:"FromTime,omitempty"`
	ToTime    int64  `json:"ToTime,omitempty"`
	Industry  string `json:"Industry,omitempty"`
	FromMcp   bool   `json:"FromMcp"`
}

// Rsp 搜索结果
type Rsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ResponseHeader struct {
			VrRequestId string `json:"vr_request_id"`
			RequestId   string `json:"request_id"`
		} `json:"response_header"`
		ResponseData struct {
			Query string `json:"query"`
			Docs  []struct {
				Type    int      `json:"type"`
				Passage string   `json:"passage"`
				Score   float64  `json:"score"`
				Date    string   `json:"date"`
				Title   string   `json:"title"`
				Url     string   `json:"url"`
				Site    string   `json:"site"`
				Images  []string `json:"images"`
			} `json:"docs"`
		} `json:"response_data"`
	} `json:"data"`
}
