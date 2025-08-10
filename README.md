# Pathy

> A minimal terminal user interface (TUI) file explorer built with Go, Bubble Tea, and Lip Gloss

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

---

## 💭 About

I wanted to build something similar to file managers like [Yazi](https://github.com/sxyazi/yazi), [Superfile](https://github.com/MHNightCat/superfile), and [LF](https://github.com/gokcehan/lf), but make it my own. This project also gave me an excuse to use the [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss) libraries because I think TUIs are beautiful.

Currently, Pathy is a minimal file explorer that can navigate through directories. It's in (extremely) early development and will slowly grow with more features over time.

## ✨ Current Features

- **📁 Directory Navigation** - Browse files and folders with vim-like keybindings
- **📂 History** - Go back to previous directories with backspace
- **🎨 Clean Interface** - Beautiful TUI powered by Lip Gloss
- **📱 Responsive** - Adapts to your terminal size

## 🚀 Installation

### From Source

```bash
git clone https://github.com/miguel-fg/pathy.git
cd pathy
go build -o pathy
./pathy
```

## 🎯 Usage

Start Pathy in the current directory:
```bash
pathy
```

Or specify a starting directory:
```bash
pathy /path/to/directory
```

## ⌨️ Keybindings

| Key | Action |
|-----|--------|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Enter` / `l` | Enter directory |
| `Backspace` / `h` | Go back to previous directory |
| `q` / `Ctrl+C` | Quit application |

## 🏗️ Built With

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Go framework for building TUIs
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - For styling terminal UIs
- **Go** - Because it's fast, I love it, and I'm learning it

## 🙏 Inspiration

This project was inspired by existing terminal file managers like:
- [Yazi](https://github.com/sxyazi/yazi)
- [Superfile](https://github.com/MHNightCat/superfile)
- [LF](https://github.com/gokcehan/lf)

## 🛠️ Development

### Prerequisites

- Go 1.21 or higher
- Terminal with true color support (recommended)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/pathy.git
cd pathy

# Install dependencies
go mod tidy

# Build the application
go build -o pathy

# Run
./pathy
```

## 🛣️ Future Plans

This is just the beginning! As the project grows, I plan to add:

- File operations (add, copy, move, delete)
- Search functionality
- Multiple selection
- File preview
- Customizable themes
- And much more (suggestions welcome)...

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Inspired by [Yazi](https://github.com/sxyazi/yazi), [Superfile](https://github.com/MHNightCat/superfile), and [LF](https://github.com/gokcehan/lf)

---

<div align="center">
  <p>Made with ❤️ and Go</p>
  <p><a href="#-pathy">↑ Back to top</a></p>
</div>