package models

type ProjectDetails struct {
	Name string `json:"name"`
	PackageManager  string `json:"packageManager"`
}

type TemplateData struct {
	Title  string
}
