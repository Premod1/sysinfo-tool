# 🖥️ sysinfo-tool

A simple and lightweight Go CLI tool that displays essential system information including OS details, CPU cores, memory usage, and Go version.

## 📋 Requirements

- **Go 1.24+** (required for building and installation)

## 🚀 Installation

### Option 1: Using `go install` (Recommended)

#### Linux / macOS
```bash
go install github.com/Premod1/sysinfo-tool@latest
```

#### Windows
```cmd
go install github.com/Premod1/sysinfo-tool@latest
```

### Option 2: Building from Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Premod1/sysinfo-tool.git
   cd sysinfo-tool
   ```

2. **Build the binary:**
   ```bash
   go build -o sysinfo-tool
   ```

3. **Run the tool:**
   ```bash
   ./sysinfo-tool info
   ```

## 🛠️ Adding to PATH

### Linux / macOS

After installing with `go install`, the binary is typically installed to `$GOPATH/bin` or `$HOME/go/bin`. To ensure it's in your PATH:

1. **Check if Go bin directory is in PATH:**
   ```bash
   echo $PATH | grep -q "$(go env GOPATH)/bin" && echo "✅ Already in PATH" || echo "❌ Not in PATH"
   ```

2. **If not in PATH, add it to your shell profile:**
   ```bash
   # For bash users (~/.bashrc or ~/.bash_profile)
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   source ~/.bashrc
   
   # For zsh users (~/.zshrc)
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
   source ~/.zshrc
   ```

### Windows

1. **Find your Go bin directory:**
   ```cmd
   go env GOPATH
   ```

2. **Add to PATH via System Properties:**
   - Open "System Properties" → "Advanced" → "Environment Variables"
   - Edit the "Path" variable under "User variables"
   - Add `%USERPROFILE%\go\bin` (or the path from step 1 + `\bin`)
   - Click "OK" and restart your command prompt

3. **Alternative: Add to PATH via Command Line:**
   ```cmd
   setx PATH "%PATH%;%USERPROFILE%\go\bin"
   ```

## 📖 Usage

Run the CLI tool using one of the supported keywords:

```bash
sysinfo-tool info
```

or

```bash
sysinfo-tool system
```

### Example Output

```
=== System Information ===
OS: linux
Architecture: amd64
CPU Cores: 8
Go Version: go1.24.7
Allocated Memory: 0 MB
```

## 💡 Notes / Tips

- **For beginners:** Make sure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/)
- **Troubleshooting:** If the command is not found after installation, ensure your Go bin directory is in your system's PATH
- **Memory Usage:** The "Allocated Memory" shows the memory currently allocated by the Go runtime for this specific program
- **Cross-platform:** This tool works on Linux, macOS, and Windows
- **Lightweight:** The compiled binary is very small and has no external dependencies

## 🔧 Development

To contribute or modify the tool:

1. Fork the repository
2. Make your changes
3. Test with: `go run main.go info`
4. Build with: `go build -o sysinfo-tool`
5. Submit a pull request

## 📝 License

This project is open source. Feel free to use and modify as needed.

---

**Happy system monitoring!** 🎉