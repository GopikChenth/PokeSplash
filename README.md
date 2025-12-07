# PokéSplash ⚡

**Ditch the boring black screen!** PokéSplash greets you with a wild Pokémon in colorful retro style every time you open your terminal. It loads instantly, so you’re never waiting to code. From cute starters to epic legendaries, your command line just became a mini adventure. Catch 'em all while you work!

## Features

- ⚡ **Instant startup** (<3ms) — perfect for shell profiles
- 🎨 **898 Pokémon** with high-quality 24-bit ANSI colors
- 🎲 **Random** encounter on every terminal launch
- 📦 **Tiny binary** (~7MB) with zero dependencies
- 💻 **Cross-platform** (Windows, Linux, macOS)

## Installation

### Quick Install (Script)

**Windows (PowerShell):**

```powershell
.\install.ps1
```

**Linux/macOS:**

```bash
chmod +x install.sh
./install.sh
```

### Android (Termux)

You can run PokéSplash on your phone!

1. **Build for Android:**
   ```powershell
   .\build-termux.ps1
   ```
2. **Transfer** the `pokesplash-android` binary to your phone.
3. **Run in Termux:**
   ```bash
   chmod +x pokesplash-android
   ./pokesplash-android
   ```

### From Binary (Manual)

Download the latest release and add it to your PATH.

### From Source (Go)

```bash
go install github.com/GopikChenth/PokeSplash@latest
```

## Usage

Run it directly:

```bash
pokesplash
```

Or with flags:

```bash
pokesplash --pokemon mewtwo    # Summon a legendary
pokesplash --pokemon magikarp  # Splash!
pokesplash --list              # See the Pokédex
```

### Add to Terminal Startup

**PowerShell (`$PROFILE`):**

```powershell
pokesplash
```

**Bash/Zsh (`.bashrc` / `.zshrc`):**

```bash
pokesplash
```

## License

MIT
