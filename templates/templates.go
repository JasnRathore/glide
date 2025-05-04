package templates

import (
	"embed"
	"text/template"
	"os"
	"fmt"
	
	//"glide/models"
)

//go:embed *.tmpl
var templatesFS embed.FS

func LoadTemplate(templateName string) (string, error) {
	content, err := templatesFS.ReadFile(templateName)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", templateName, err)
	}
	return string(content), nil
}

func GenerateTemplate(templateName string, outputFile string,data interface{},) error {
	
	//templateName := "main.go.tmpl"
	//outputFile := "main.go"
	
	templateContent, err := LoadTemplate(templateName)
	if err != nil {
		return fmt.Errorf("failed to Load template: %w", err)
	}
	
	tmpl, err := template.New("generated").Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	
	return nil
}

func CopyTemplate(templateName string, outputFile string) error {
    templateContent, err := LoadTemplate(templateName)
    if err != nil {
        return fmt.Errorf("failed to load template: %w", err)
    }
    
    outFile, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer outFile.Close()
    
    _, err = outFile.WriteString(templateContent)
    if err != nil {
        return fmt.Errorf("failed to write to output file: %w", err)
    }
    
    return nil
}