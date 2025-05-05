package commands

import (
	"fmt"
	"strings"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	utils "glide/utils"
	models "glide/models"
	tmpl "glide/templates"
)

type state int

const (
	stateInput state = iota
	stateMenu
	stateDone
)

type model struct {
	state     state
	input     string
	selection int
	options   []string
	done      bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case stateInput:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				m.state = stateMenu
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyCtrlC: 
				return m, tea.Quit
			default:
				m.input += msg.String()
			}
		}
	case stateMenu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyUp:
				if m.selection > 0 {
					m.selection--
				}
			case tea.KeyDown:
				if m.selection < len(m.options)-1 {
					m.selection++
				}
			case tea.KeyEnter:
				m.state = stateDone
				m.done = true
				return m, tea.Quit
			case tea.KeyCtrlC: 
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	switch m.state {
	case stateInput:
		b.WriteString("Enter Project Name: ")
		b.WriteString(m.input)
	case stateMenu:
		b.WriteString(fmt.Sprintf("You typed: %s\n\nChoose an option:\n\n", m.input))
		for i, opt := range m.options {
			cursor := " "
			if i == m.selection {
				cursor = ">"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
		}
	case stateDone:
		b.WriteString(fmt.Sprintf("You typed: %s\n", m.input))
		b.WriteString(fmt.Sprintf("You selected: %s\n", m.options[m.selection]))
	}
	return b.String()
}

func CreateFrontend(packageManager string, name string) {
	switch packageManager {
		case "npm":
			err := utils.RunCommand("npm", "create", "vite@latest", strings.ToLower(name))
			utils.Check(err)
		case "pnpm":
			err := utils.RunCommand("pnpm", "create", "vite", strings.ToLower(name))
			utils.Check(err)
		case "yarn":
			err := utils.RunCommand("yarn", "create", "vite", strings.ToLower(name))
			utils.Check(err)
		case "bun":
			err := utils.RunCommand("bun", "create", "vite", strings.ToLower(name))
			utils.Check(err)
		case "deno":
			err := utils.RunCommand("deno", "init", "--npm","vite", strings.ToLower(name))
			utils.Check(err)
		default:
			fmt.Println("default")
	}
}


func InstallFrontendDependencies(packageManager string) {
	switch packageManager {
		case "npm":
			err := utils.RunCommand("npm", "install")
			utils.Check(err)
		case "pnpm":
			err := utils.RunCommand("pnpm", "install")
			utils.Check(err)
		case "yarn":
			err := utils.RunCommand("yarn")
			utils.Check(err)
		case "bun":
			err := utils.RunCommand("bun", "install")
			utils.Check(err)
		case "deno":
			err := utils.RunCommand("deno", "install")
			utils.Check(err)
		default:
			fmt.Println("default")
	}
}

func GetFrontendDependenciesCommand(packageManager string) string {
	switch packageManager {
		case "npm":
			return "npm install"
		case "pnpm":
			return "pnpm install"
		case "yarn":
			return "yarn"
		case "bun":
			return "bun install"
		case "deno":
			return "deno install"
		default:
			return "default"
	}
}

func ui() (models.ProjectDetails, error) {
	m := model {
		state:   stateInput,
		options: []string{"npm", "yarn", "pnpm", "deno", "bun"},
	}

	prog := tea.NewProgram(m)
	finalModel, err := prog.Run()
	utils.Check(err)

	m = finalModel.(model) // type assert to get final state

	if m.done {
		return models.ProjectDetails{
			Name:  m.input,
			PackageManager: m.options[m.selection],
		}, nil
	}
	return models.ProjectDetails{}, err
}



func InitProject() {
	//createing the webapp
	project, err := ui()
	utils.Check(err)
	CreateFrontend(project.PackageManager, project.Name)
	
	//go into projfolder
	err = os.Chdir(strings.ToLower(project.Name))
	utils.Check(err)
	
	//install dependencies
	//InstallFrontendDependencies(project.PackageManager)
	
	
	//making the glide config file
	jsonData, err := utils.StructToJSON(project)
	utils.Check(err)
	err = utils.WriteJSONToFile("glide.config.json", jsonData)
	utils.Check(err)
	
	//making the src-glide dir
	dirName := "src-glide"
	err = os.Mkdir(dirName, 0755)
	utils.Check(err)
	err = os.Chdir(dirName)
	utils.Check(err)
	
	//init go proj
	utils.RunCommand("go", "mod", "init", strings.ToLower(project.Name))
	
	//init air {hotrealoading}
	//utils.RunCommand("air", "init")
	tmpl.CopyTemplate("air.toml.tmpl",".air.toml")	
	
	//installing dependencies	
	repoName := "github.com/JasnRathore/glide-lib"
	utils.RunCommand("go", "get", repoName)
	
	
	//generating src-glide files
	data := models.TemplateData {
		Title: strings.ToLower(project.Name),
	}
	tmpl.GenerateTemplate("main.go.tmpl","main.go",data)	
	err = os.Mkdir("app", 0755)
	tmpl.GenerateTemplate("app.go.tmpl","app/app.go",data)	
	tmpl.GenerateTemplate("build.go.tmpl","build.go",data)	
	
	fmt.Println("cd",strings.ToLower(project.Name))
	fmt.Println(GetFrontendDependenciesCommand(project.PackageManager))
	fmt.Println("glide dev")
	
}

