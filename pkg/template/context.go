package template

type Context interface {
	Get(key string) interface{}
	Put(key string, value interface{})
}

type HtmlContext struct {
	values map[string]interface{}
}

func NewHtmlContext() *HtmlContext {
	return &HtmlContext{values: make(map[string]interface{})}
}

func (h *HtmlContext) Get(key string) interface{} {
	return h.values[key]
}

func (h *HtmlContext) Put(key string, value interface{}) {
	h.values[key] = value
}
