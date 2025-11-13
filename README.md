# 🔐 Password Manager

A secure, interactive command-line password manager built with Go, featuring an elegant tree-based UI powered by Bubble Tea.

## Features

- 🔒 **Secure Storage**: Passwords are encrypted at rest using industry-standard encryption
- 🎨 **Interactive UI**: Beautiful tree-view interface with colors and intuitive navigation
- ⌨️ **Keyboard Navigation**: Vim-style keybindings and arrow key support
- 👁️ **Password Reveal**: Toggle individual or all passwords on demand
- 📅 **Metadata Tracking**: Automatic tracking of creation, update, and last read times
- 📋 **Clipboard Integration**: Automatically copy passwords to clipboard
- 🌳 **Tree Structure**: Organize passwords by domain with expandable/collapsible nodes

## Installation

```bash
go build -o password-manager
```

## Usage

### Initialize the Vault

Before using the password manager, initialize your vault:

```bash
./password-manager init
```

This will create an encrypted vault file to store your passwords.

### Add a Password

```bash
./password-manager add <website> <username> <password>
```

**Important**: When your password contains special characters (like `*`, `?`, `$`, `!`, etc.), wrap it in quotes:

```bash
./password-manager add example.com user 'MyP@ssw0rd*123!'
```

Examples:

```bash
./password-manager add github.com myusername 'SecurePass123!'
./password-manager add gmail.com user@email.com 'P@ssw*rd$456'
```

### List Passwords (Interactive UI)

Display all passwords in an interactive tree view:

```bash
./password-manager list
```

#### Keyboard Controls

| Key               | Action                                           |
| ----------------- | ------------------------------------------------ |
| `↑` / `k`         | Move cursor up                                   |
| `↓` / `j`         | Move cursor down                                 |
| `Enter` / `Space` | Expand/collapse domain or toggle password reveal |
| `r`               | Reveal/hide password at cursor                   |
| `R`               | Reveal/hide all passwords                        |
| `q` / `Ctrl+C`    | Quit                                             |

#### Interactive UI Features

- **Domain Grouping**: Passwords are organized by domain
- **Entry Count**: See how many accounts you have for each domain
- **Password Masking**: Passwords are masked by default with asterisks
- **Selective Reveal**: Press `r` to reveal individual passwords or `R` to toggle all
- **Metadata Display**: View when each password was created and last updated
- **Color Coding**:
  - 🔵 Blue: Selected domain
  - 🟢 Green: Expanded domain
  - 🟣 Purple: Collapsed domain
  - 🟡 Yellow: Highlighted selection

### Get a Password

Retrieve a specific password and copy it to clipboard:

```bash
./password-manager get <website> <username>
```

Example:

```bash
./password-manager get github.com myusername
```

The password will be automatically copied to your clipboard.

## Special Characters in Passwords

The shell interprets certain characters as special commands. Always wrap passwords containing these characters in single quotes (`'`):

### Characters that need quoting:

- `*` (asterisk)
- `?` (question mark)
- `$` (dollar sign)
- `!` (exclamation mark)
- `&` (ampersand)
- `|` (pipe)
- `;` (semicolon)
- `<` `>` (angle brackets)
- `(` `)` (parentheses)
- `{` `}` (braces)
- `[` `]` (brackets)
- `` ` `` (backtick)
- `\` (backslash)
- Space characters

### Examples:

✅ **Good**:

```bash
./password-manager add site.com user 'Pass*word!123'
./password-manager add site.com user 'My$ecure&P@ss'
```

❌ **Bad** (will cause errors):

```bash
./password-manager add site.com user Pass*word!123
./password-manager add site.com user My$ecure&P@ss
```

## Architecture

### Project Structure

```
password-manager/
├── cmd/                    # CLI commands
│   ├── add.go             # Add password command
│   ├── get.go             # Get password command
│   ├── init.go            # Initialize vault command
│   ├── list.go            # Interactive list command
│   └── root.go            # Root command setup
├── internal/
│   ├── encryption/        # Encryption utilities
│   ├── passkey/           # Passkey management
│   ├── session/           # Session handling
│   ├── storage/           # File handling and storage
│   └── ui/                # Bubble Tea UI components
│       ├── model.go       # UI state model
│       ├── update.go      # Event handling
│       ├── view.go        # Rendering logic
│       └── styles.go      # Visual styles
├── vault/                 # Vault data structures
└── main.go                # Entry point
```

### Technologies Used

- **[Cobra](https://github.com/spf13/cobra)**: CLI framework
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: Terminal UI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Terminal styling
- **[go-humanize](https://github.com/dustin/go-humanize)**: Human-readable time formatting

## Security

- All passwords are encrypted at rest
- Passwords are masked by default in the UI
- Clipboard integration for secure password retrieval
- File permissions are set to user-only access (0600)

## Development

### Building from Source

```bash
go build -o password-manager
```

### Running Tests

```bash
go test ./...
```

## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
