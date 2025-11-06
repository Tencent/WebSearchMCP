package constants

// ToolName  mcp server tool name
type ToolName string

const (
	// ProSearch 增强搜索
	ProSearch ToolName = "prosearch"
)

// String 返回tool name的string值
func (tn ToolName) String() string {
	return string(tn)
}

const (
	ProSearchProxyURL = "wsa.tencentcloudapi.com"
	Timeout           = 6000
	ToolDesc          = "针对用户主动搜索的问题（query），基于问答搜索的海量数据 & 权威内容 & 智能化答案抽取 生成具体的答案内容"
	InputSchema       = `{
  "properties": {
    "query": {
      "description": "查询语",
      "type": "string"
    },
    "mode": {
      "description": "0：自然检索结果，1：VR卡结果，2：混合结果（VR+自然检索）、（VR卡：指较为可信来源的数据，如天气、金价数据）",
      "type": "integer"
    },
    "site": {
      "description": "指定网址搜索, 需要查询某个特定网址的内容时，可传入该参数，仅支持单域名筛选，例如：qq.com",
      "type": "string"
    },
    "from_time": {
      "description": "起始时间, 当query具有强时效性可使用该参数, 搭配to_time参数进行使用，获取指定的时间段的搜索结果。例如：2025-01-21 15:57:01, 需要和to_time参数搭配使用，单用会报错",
      "type": "string"
    },
    "end_time": {
      "description": "结束时间，例如：2025-01-21 16:57:03，需要和from_time参数搭配使用，单用会报错",
      "type": "string"
    }
  },
  "type": "object",
  "required": [
    "query"
  ]
}`
)
