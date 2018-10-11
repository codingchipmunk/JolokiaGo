package responseValues

type ResponseValue struct {
	Backend Backend `json:"backend"`
	Id      string  `json:"id"`
}

type Backend struct {
	Sse  SseInfo  `json:"sse"`
	Pull PullInfo `json:"pull"`
}

type PullInfo struct {
	MaxEntries int    `json:"maxEntries"`
	Store      string `json:"store"`
}

type SseInfo struct {
	ContentType string `json:"backChannel.contentType"`
	Encoding    string `json:"backChannel.encoding"`
}
