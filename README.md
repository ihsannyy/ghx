# ghx - GitHub CLI Multi-Account Manager

`ghx` is a fast, lightweight multi-account manager CLI wrapper for [`gh`](https://cli.github.com/) (GitHub CLI) written in Go. 

It allows software developers and engineers to save, manage, and seamlessly switch between multiple GitHub accounts (synchronizing both `gh` authentication credentials and global `git user.name`/`user.email` identities) instantly with a single command.

Designed specifically for **Termux** (Android/Linux ARM64) as well as traditional Linux and macOS desktop environments.

---

## 🚀 Why Use `ghx`?

When working across personal projects, freelance work, and company organization repositories, switching GitHub accounts manually requires logging out of `gh auth logout`, re-authenticating, and updating `git config --global user.email`. 

`ghx` automates this entire process:
- 🔄 **One-Command Switching**: Switches active `gh` CLI login token and `git` author identity simultaneously.
- 🔑 **Automatic Profile Extraction**: Automatically fetches your GitHub username, display name, and verified primary email directly from GitHub API when adding a token.
- 📦 **Zero External Overhead**: Fast startup, compiled into a single static binary without heavy dependencies.
- 🎨 **Clean Box UI**: Beautiful, aligned table formatting with automatic TTY and `NO_COLOR` support.
- 🌐 **Dual Language Support**: Built-in internationalization supporting English (`en`) and Indonesian (`id`).
- ⚡ **Shell Completion**: Supports tab-autocompletion for Bash, Zsh, and Fish shells.

---

## 📋 Prerequisites

Before using `ghx`, ensure the following tools are installed on your system:

1. **Go** (v1.21 or newer, required for building from source)
2. **GitHub CLI** (`gh`)
3. **Git** (`git`)

On Termux (Android):
```bash
pkg update && pkg install golang gh git make
```

---

## 📦 Installation

### Option 1: Standard Build & Install (Recommended for Termux)

Clone the repository and install using the included `Makefile`:

```bash
git clone https://github.com/your-username/ghx.git
cd ghx
make build
make install
```

This compiles the binary into `bin/ghx` and copies it directly into your system PATH (`$PREFIX/bin/ghx`).

### Option 2: Build Manually

```bash
go build -o ghx .
mv ghx $PREFIX/bin/
chmod +x $PREFIX/bin/ghx
```

---

## 🏁 Quick Start Guide

### 1. Add your GitHub accounts
Generate a Personal Access Token (PAT) on GitHub with `repo`, `read:org`, and `gist` scopes, then run:

```bash
ghx add ghp_YourPersonalAccessToken123
```

*(You can also add a secondary work account)*:
```bash
ghx add ghp_YourWorkPersonalAccessToken456
```

### 2. View all saved accounts
```bash
ghx list
```

Example Output:
```text
┌────────┬────────────┬───────────────────┬────────────────────┐
│ STATUS │ USERNAME   │ NAME              │ EMAIL              │
├────────┼────────────┼───────────────────┼────────────────────┤
│        │ dev-user   │ Developer User    │ dev@company.org    │
│   *    │ octocat    │ Mona Lisa Octocat │ octocat@github.com │
└────────┴────────────┴───────────────────┴────────────────────┘
```

### 3. Switch between accounts
```bash
ghx switch dev-user
```
Now both your `gh` CLI session and your `git config --global` name and email are active as `dev-user`!

### 4. Check currently active account
```bash
ghx current
```

---

## 📖 Complete Command Reference

| Command | Arguments | Description |
| :--- | :--- | :--- |
| `ghx add` | `<token>` | Add a GitHub account using Personal Access Token and set it as active |
| `ghx login` | `[token]` | Authenticate an account (prompts for token or imports active `gh` session) |
| `ghx switch` | `<account>` | Switch active `gh` authentication & `git` global identity to `<account>` |
| `ghx list` | None | Display a formatted box table of all saved accounts and active status |
| `ghx current` | None | Display details (username, display name, email) of current active account |
| `ghx email` | `[email]` | Display or update the email address for current active account & git config |
| `ghx remove` | `<account>` | Remove a saved account from `ghx` configuration |
| `ghx lang` | `[en\|id]` | View or set application interface language (`en` for English, `id` for Indonesian) |
| `ghx completion` | `<shell>` | Generate shell completion script (`bash`, `zsh`, `fish`, `powershell`) |

---

## 💡 Command Usage & Examples

### `ghx add <token>`
Adds a GitHub account via token, automatically queries your GitHub profile info, saves it to `ghx`, and activates it.

```bash
ghx add ghp_xxxxxxxxxxxxxxxxxxxx

# Optional flags:
ghx add ghp_xxxxxxxxxxxxxxxxxxxx --name "My Custom Name" --email "custom@domain.com"
```

### `ghx switch <account>`
Switches your active GitHub CLI token and global git identity (`git config --global user.name` & `user.email`) to the specified username.

```bash
ghx switch octocat
```

### `ghx list`
Lists all accounts saved in `ghx`. The active account is indicated with an asterisk (`*`) and highlighted in green.

```bash
ghx list
```

### `ghx current`
Displays the username, name, and email of the currently active account.

```bash
ghx current
```

### `ghx email [email]`
Displays your active git email or updates it for both `ghx` config and global git config.

```bash
# Check active email:
ghx email

# Update active email:
ghx email newemail@example.com
```

### `ghx remove <account>`
Removes an account from `ghx` config. If the removed account was active, clears the current active account setting.

```bash
ghx remove dev-user
```

### `ghx lang [en|id]`
Switches the CLI interface language between English and Indonesian.

```bash
# Switch to Indonesian:
ghx lang id

# Switch back to English:
ghx lang en
```

---

## ⚡ Shell Autocompletion Setup

Enable tab-autocompletion for `ghx` commands in your preferred shell:

### Bash (Default Termux Shell)
Add the autocompletion script to your `~/.bashrc`:
```bash
echo 'source <(ghx completion bash)' >> ~/.bashrc
source ~/.bashrc
```

### Zsh
```zsh
echo 'source <(ghx completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

### Fish
```fish
mkdir -p ~/.config/fish/completions
ghx completion fish > ~/.config/fish/completions/ghx.fish
```

---

## 📁 Configuration File

`ghx` stores its configuration in JSON format at:
```text
~/.config/ghx/accounts.json
```

Example structure:
```json
{
  "current_account": "octocat",
  "language": "en",
  "accounts": {
    "octocat": {
      "username": "octocat",
      "name": "Mona Lisa Octocat",
      "email": "octocat@github.com",
      "token": "ghp_xxxxxxxxxxxxxxxxxxxx",
      "host": "github.com"
    }
  }
}
```

---

## ❓ Frequently Asked Questions (FAQ)

**Q: Where are my tokens stored?**  
A: Tokens are saved locally in your user configuration directory at `~/.config/ghx/accounts.json` with restricted `0600` file permissions. `ghx` never sends tokens to any external server other than official GitHub APIs.

**Q: Does `ghx switch` change my SSH keys?**  
A: `ghx` updates `gh` CLI credentials and `git config --global user.name` & `user.email`. If you rely on HTTPS or `gh` credential helper, git operations will automatically use the active account's token.

**Q: Can I use `ghx` with GitHub Enterprise Server?**  
A: Yes! You can pass the `--hostname` flag when adding an account:
```bash
ghx add <token> --hostname github.mycompany.com
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
