<h1 align="center">GuardaTuti</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat-square&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/Status-In%20Development-yellow?style=flat-square" alt="Status" />
  <img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" alt="License" />
</p>

<p align="center">
  <b>A secure and simple CLI password vault for the modern developer</b>
</p>

## 🔐 About

GuardaTuti is a lightweight command-line password manager built in Go. It generates strong passwords and securely stores your credentials locally, making them accessible with simple commands.

## ✨ Features

### 🚀 Currently Available

- Generate strong random passwords automatically
- Store credentials (service, username, password)
- List saved credentials with creation timestamps

#### Commands

- `generate` - Create custom secure passwords
- `get` - Retrieve credentials with one-click clipboard copy
- `list` - View all stored services in a formatted table
- `show` - Display detailed information for a specific service
- `delete` - Remove unwanted credentials
- `help` - Comprehensive help system

#### User Interface

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) TUI enhancements:
  - Interactive credential listing with filtering
  - Beautiful password generation forms
  - Detailed credential view with masked passwords
  - Confirmation dialogs for sensitive operations

## 📦 Installation

[wip]

```bash
# Clone the repository
# Build the application
# Run GuardaTuti
## 🔧 Usage
```

## 🛣️ Roadmap

- [ ] Implement persistent storage
- [ ] Complete command set -[ ] generate -[ ] get + copy -[ ] list -[ ] show -[ ] delete
- [ ] Enhance UI with Bubble Tea
- [ ] Add clipboard integration
- [ ] Make mac/linux/windows client available

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.
