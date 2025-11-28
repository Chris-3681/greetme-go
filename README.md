# greetme-go

A tiny Go command-line app that reads your name from the terminal and prints a greeting.  
Optionally, it can shout the greeting in uppercase.

## Requirements

- Go 1.20+ installed (`go version` should work)
- Git (if cloning from GitHub)

## Getting Started


clone
```bash
git clone https://github.com/Chris-3681/greetme-go.git
cd greetme-go
```
# run
```bash
go run main.go
```

You should see:
```bash
Enter your name:
```

Type your name and press Enter.

Example output:
```bash
Enter your name: Chris
Hello, Chris! Welcome to Go 🎉
Generated at: Thu, 27 Nov 2025 21:30:00 EAT
```
## Optional: Shout Mode

The app supports a --shout flag:

go run main.go `-- --shout`


This prints the greeting in uppercase.

Build a Binary
```bash
go build -o greetme .
./greetme        # Linux/macOS
greetme.exe      # Windows
```
Project Structure

`main.go` — application entrypoint and logic

`go.mod` — Go module definition


Commit & push:

```
git add main.go README.md
git commit -m "docs: add README and finalize greetme CLI"
git push
```