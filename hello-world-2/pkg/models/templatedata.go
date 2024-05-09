package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //interface used when you don't know what type it is
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
