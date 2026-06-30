# 🚀 GoInit

> Minimal questions. Production-ready defaults.

GoInit is a Vite-inspired CLI tool for Go developers that generates production-ready project structures with sensible defaults.

Instead of spending time creating folders, configuring Docker, or organizing your project structure, GoInit helps you start building immediately.

---

## ✨ Features

- Interactive CLI experience
- Minimal questions
- Installable globally
- Support for multiple frameworks
- Support for multiple databases
- Vite-like developer experience
- Production-ready defaults

---

## 📦 Installation

### Using Go

```bash
go install github.com/dev-marees/goinit@latest
```

Verify:

```bash
goinit --version
```

Output:

```text
goinit version v0.1.0
```

---

## 🚀 Usage

### Interactive Mode

```bash
goinit create
```

Example:

```text
? Project name: ecommerce-api

? Select framework:
❯ Gin
  Fiber

? Select database:
❯ PostgreSQL
  MySQL
  None
```

---

### Partial Interactive Mode

```bash
goinit create -n ecommerce-api
```

Only asks for missing information.

---

### Non-Interactive Mode

```bash
goinit create \
    -n ecommerce-api \
    --framework gin \
    --database postgres
```

---

## ⚙️ Available Options

| Flag | Description |
|------|-------------|
| `-n, --name` | Project name |
| `--framework` | gin, fiber |
| `--database` | postgres, mysql, none |

---

## 📖 Examples

### Gin + PostgreSQL

```bash
goinit create \
    -n ecommerce \
    --framework gin \
    --database postgres
```

### Fiber + MySQL

```bash
goinit create \
    -n inventory \
    --framework fiber \
    --database mysql
```

---

## 🗂 Current Project Structure

```text
goinit/
├── cmd/
│   ├── root.go
│   └── create.go
│
├── internal/
│   └── prompts/
│       ├── config.go
│       └── prompt.go
│
├── main.go
├── go.mod
└── README.md
```

---

## 🛣 Roadmap

### v0.1.0

- [x] Cobra CLI
- [x] Global installation
- [x] Version command
- [x] Interactive prompts
- [x] Optional flags

### v0.2.0

- [ ] Project structure generation
- [ ] go.mod generation
- [ ] .env generation
- [ ] README generation

### v0.3.0

- [ ] Gin template
- [ ] Fiber template
- [ ] Database templates

### v0.4.0

- [ ] Docker support
- [ ] Docker Compose
- [ ] Makefile

### v1.0.0

- [ ] Production-ready scaffolding
- [ ] Plugin system
- [ ] Additional frameworks

---

## 🎯 Philosophy

GoInit follows three principles:

- Minimal Questions
- Sensible Defaults
- Production-Ready Projects

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome.

```bash
git clone <repository-url>

cd goinit

go install

goinit create
```

---

## 📄 License

MIT License

---

## ⭐ Support

If you find GoInit useful, consider giving the repository a star.

---

Built with ❤️ using Go and Cobra.