package models

// TemplateData holds the data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	// These are just messages
	Flash   string
	Warning string
	Error   string
}
