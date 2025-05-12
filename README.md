
# Glide 🚀

> Build desktop apps with Go and your favorite frontend framework

Glide is a modern toolkit that bridges the gap between Go backend development and your preferred JavaScript/TypeScript frontend. Create beautiful, responsive desktop applications that harness the power of Go's performance and the flexibility of web technologies - all with minimal configuration.

## 🔍 Overview

Glide creates a seamless development experience by combining:

- **Go backend** - For system-level operations, performance, and cross-platform compatibility
- **Web-based frontend** - Using your preferred JS framework (React, Vue, Svelte, etc.)
- **Streamlined workflow** - Hot-reloading during development, simple bundling for production

The result is a desktop application that runs natively across operating systems while maintaining a modern, responsive UI.

## ⚙️ Requirements

To use Glide, you need the following tools installed:

- **Go** (version 1.21+)
- **Node.js** (version 14+)
- **Air** - For Go hot-reloading
- One of the following JS package managers:
  - npm
  - yarn
  - pnpm
  - bun
  - deno

## 🚀 Getting Started

### Installation

1. Install Glide using Go:

```bash
go install github.com/JasnRathore/glide@latest
```

### Create a New Project

Initialize a new Glide project with:

```bash
glide init
```

This interactive command will:
1. Prompt you for a project name
2. Let you select your preferred package manager
3. Verify dependencies are installed
4. Create a new Vite project with your selected configuration
5. Set up the Go backend structure
6. Initialize configurations for hot-reloading
7. Install necessary dependencies

## 🏗️ Project Structure

After initialization, your project will have the following structure:

```
your-project/
├── src/                # Frontend source code (Vite project)
│   └── glide/          # Glide JS/TS utilities
│       ├── glide.js    # JavaScript interface
│       └── glide.ts    # TypeScript interface
├── src-glide/          # Go backend code
│   ├── app/            # Application logic
│   │   └── app.go      # App configuration
│   ├── main.go         # Entry point for development
│   ├── build.go        # Build configuration
│   └── .air.toml       # Hot-reload config
└── glide.config.json   # Project configuration
```

## 💻 Development Workflow

### Running in Development Mode

Start your development server with:

```bash
glide dev
```

This command:
1. Starts your frontend dev server (using your chosen package manager)
2. Runs the Go backend with hot-reloading via Air
3. Connects the two together for seamless development

Any changes to your frontend or Go code will automatically reload.

### Calling Go Functions from JavaScript

Glide provides a simple interface to call Go functions from your frontend:

```javascript
// Using the JavaScript helper
import { callWindowFunction } from './glide/glide.js';

// Call a Go function registered in app.go
const greeting = callWindowFunction('Greet', 'World');
console.log(greeting); // "Hello, World"
```

TypeScript users can use the strongly-typed interface:

```typescript
import { callWindowFunction } from './glide/glide.ts';

// Type-safe function calls
const greeting = callWindowFunction<[string], string>('Greet', 'World');
console.log(greeting); // "Hello, World"
```

### Registering Go Functions

To expose Go functions to your frontend, modify the `app.go` file:

```go
func Greet(name string) string {
    return fmt.Sprintf("Hello, %s", name)
}

func App() *glide.App {
    // App configuration...
    
    // Register functions to be called from JavaScript
    funcs := []interface{}{Greet}
    app.InvokeHandler(funcs)
    
    // ...
}
```

## 📦 Building for Production

When you're ready to distribute your application, create a production build:

```bash
glide build
```

This command:
1. Builds your frontend for production
2. Copies the build artifacts to the Go project
3. Compiles everything into a single executable
4. Places the executable in the `src-glide/target` directory

The resulting binary includes your frontend assets and can be distributed as a standalone application.

## 🎛️ Configuration

### App Configuration

Customize your application in `app.go`:

```go
func App() *glide.App {
    config := glide.AppConfig{
        Title:     "MyApp",
        Width:     800,
        Height:    600,
        Debug:     true,
        AutoFocus: true,
        IconID:    1,
        
        Tray: &glide.TrayConfig{
            IconID:  2,
            Title:   "MyApp",
            Tooltip: "My Awesome App",
        },
    }

    app := glide.New(config)
    // ...
}
```

## 🛠️ Advanced Usage

### System Tray Integration

Glide makes it easy to add system tray functionality:

```go
app.AddMenuItem(glide.MenuItem{
    Title: "Show Window",
    Handler: func() {
        app.ShowWindow()
    },
})

app.AddMenuItem(glide.MenuItem{
    Title: "Exit",
    Handler: func() {
        app.Exit()
    },
})
```

### Finding Available Ports

The build system automatically finds available ports for your application:

```go
port, err := findAvailablePort(8080)
if err != nil {
    log.Fatalf("Error finding available port: %v", err)
}
```

## 📚 Library Dependencies

Glide uses the following key libraries:

- **Backend**:
  - [github.com/JasnRathore/glide-lib](https://github.com/JasnRathore/glide-lib) - Core Glide functionality
  - [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework for interactive CLI

- **Frontend**:
  - [Vite](https://vitejs.dev/) - Next generation frontend tooling

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve Glide.

---

Happy building with Glide! 🚀