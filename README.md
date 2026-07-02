# Sai Charan — Terminal Portfolio

A full-screen terminal portfolio built with Go, inspired by terminal.shop's design philosophy. Accessible via SSH from anywhere in the world.

## Demo Video

A live demonstration of the portfolio in action.

*(Upload `demo.mp4` to the repo and watch it here, or view directly in GitHub)*

## Features

- **Full-screen TUI** — No scrolling, no browser needed, pure terminal
- **4 pages** — Home, About, Projects, Contact
- **Keyboard-driven** — Arrow keys navigate, Tab cycles focus, Enter opens links
- **ASCII branding** — Custom SC block logo
- **Dark theme** — Near-black background with teal accents
- **SSH-accessible** — Deploy on any Linux server and connect from any terminal

## Tech Stack

| | |
|---|---|
| Language | Go 1.22+ |
| Framework | [Bubble Tea](https://github.com/charmbracelet/bubbletea) |
| Styling | [Lip Gloss](https://github.com/charmbracelet/lipgloss) |
| Components | [Bubbles](https://github.com/charmbracelet/bubbles) |

## Keyboard Controls

| Key | Action |
|---|---|
| `←` / `→` | Navigate between pages |
| `↑` / `↓` | Navigate project list (Projects page) |
| `Tab` | Cycle focus between buttons |
| `Enter` | Open URL or enter project detail |
| `Esc` | Go back from project detail |
| `Ctrl+C` | Quit |

## Build from Source

### Prerequisites

- Go 1.22 or later
- Git
- A Unix-like terminal (Linux, macOS, Windows Terminal with WSL)

### Steps

```bash
# Clone the repository
git clone https://github.com/Charan8x/TUI-Portfolio.git
cd TUI-Portfolio

# Build
go build -o portfolio .

# Run
./portfolio
```

## Deployment

This portfolio is designed to run on a Linux server accessible via SSH.

1. Copy the binary to your Linux server
2. Configure SSH to run the binary as a login shell for a dedicated user
3. Users connect via: `ssh user@your-server-ip`

## About Me

- **Name:** Sai Charan
- **Title:** Full Stack Developer & Cloud Enthusiast
- **Location:** India
- **Currently:** Working as an intern

### Skills

- **Languages:** Python, C, C++, Go, JavaScript, HTML & CSS
- **Frameworks:** React, Node.js, Flask, FastAPI, Tailwind CSS, Bubble Tea, PyTorch
- **Concepts:** Database Management (DBMS), Object Oriented Programming (OOP), Data Structures & Algorithms (DSA), Deep Learning

## Projects

### 1. License Plate Detection & Recognition
**Tech:** Python, YOLOv11x, PyTorch, EasyOCR, OpenCV, ByteTrack, CUDA

Real-time ALPR system using a fine-tuned YOLOv11x model and EasyOCR with ByteTrack tracking. Detects and reads license plates from live CCTV feeds — including moving vehicles across varying weather conditions. Accuracy: 85–90%.

[View on GitHub](https://github.com/Charan8x/License-plate-Detection-ML)

### 2. SSH Portfolio TUI
**Tech:** Go, Bubble Tea, Lip Gloss

This very project. A full-screen terminal portfolio accessible over SSH with keyboard-driven navigation and a clean monochrome aesthetic.

[View on GitHub](https://github.com/Charan8x/TUI-Portfolio)

### 3. CookBook — Virtual Cooking Assistant
**Tech:** React, Vite, Tailwind CSS, Node.js, Express, MongoDB, JWT, bcryptjs

Full-stack recipe discovery and management platform with 166+ recipes across 8 world cuisines, JWT authentication, smart search, filters, favourites, full CRUD, and YouTube tutorial integration.

[View on GitHub](https://github.com/Charan8x/Cookbook-Virtual-Cooking-Assistant)

## Connect

| Platform | Link |
|---|---|
| GitHub | https://github.com/Charan8x |
| LinkedIn | https://www.linkedin.com/in/saicharan777 |
| Email | bunnyappasani5@gmail.com |
| Resume | [View PDF](https://drive.google.com/file/d/1ADwv2OfzWcr5WKGlBNDaLM3cTQWXXp7g/view) |