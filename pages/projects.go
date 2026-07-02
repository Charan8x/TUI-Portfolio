package pages

import (
	"ssh-portfolio/model"
	"ssh-portfolio/styles"

	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------------------------
// PROJECT DATA
// Each project has a name, description, tech stack, features, and optional links.
// ---------------------------------------------------------------------------

// project represents a single portfolio project.
type project struct {
	name        string
	description string
	tech        string
	features    []string
	githubURL   string
	demoURL     string
}

// projects is the master list of all portfolio projects.
var projects = []project{
	{
		name:        "License Plate Detection & Recognition",
		description: "Real-time Automatic License Plate Recognition (ALPR) system using a fine-tuned YOLOv11x model and EasyOCR with ByteTrack tracking. Detects and reads license plates from live CCTV feeds — including moving vehicles across varying weather conditions. Draws a bounding box with OCR overlay for easy identification.",
		tech:        "Python  ·  YOLOv11x  ·  PyTorch  ·  EasyOCR  ·  OpenCV  ·  ByteTrack  ·  CUDA",
		features: []string{
			"Real-time detection on video footage",
			"Persistent tracking — same plate keeps same ID across frames",
			"OCR with voting system for maximum accuracy",
			"5-variant image preprocessing for challenging conditions",
			"GPU accelerated with CUDA support",
			"All-weather operation — rain, snow, day, night",
			"Clean bounding box with plate text overlay",
		},
		githubURL: "https://github.com/Charan8x/License-plate-Detection-ML",
		demoURL:   "",
	},
	{
		name:        "SSH Portfolio TUI",
		description: "A full-screen terminal portfolio application built with Go and Bubble Tea. Features multi-page keyboard-driven navigation, a clean monochrome design with teal accents, ASCII branding, and a responsive layout. Designed for SSH deployment on cloud infrastructure.",
		tech:        "Go  ·  Bubble Tea  ·  Lip Gloss  ·  Bubbles",
		features: []string{
			"Full-screen TUI with four content pages",
			"Keyboard-driven navigation — no mouse required",
			"ASCII block logo branding",
			"Responsive layout that adapts to terminal size",
			"Status bar with interactive key hints",
			"Project detail view with focusable buttons",
		},
		githubURL: "",
		demoURL:   "",
	},
	{
		name:        "CookBook — Virtual Cooking Assistant",
		description: "A modern full-stack recipe discovery and management platform. Explore 166+ recipes across 8 world cuisines, save favourites, create your own recipes, and watch video tutorials — all in one place with a responsive dark-mode interface.",
		tech:        "React  ·  Vite  ·  Tailwind CSS  ·  Node.js  ·  Express  ·  MongoDB  ·  JWT  ·  bcryptjs",
		features: []string{
			"JWT authentication with secure register and login",
			"Smart search by name, ingredient, cuisine, or tag",
			"Filter by cuisine type and difficulty with pagination",
			"166+ pre-loaded recipes across 8 world cuisines",
			"Favourites save and management system",
			"Full CRUD — create, edit, and delete your own recipes",
			"YouTube tutorial integration on every recipe",
			"Full dark/light mode toggle",
		},
		githubURL: "https://github.com/Charan8x/Cookbook-Virtual-Cooking-Assistant",
		demoURL:   "",
	},
}

// ---------------------------------------------------------------------------
// SELECTION MARKER
// A small diamond indicator shown next to the currently selected project.
// ---------------------------------------------------------------------------

var selectionMarker = styles.Accent.Render("⟐ ")
var emptyMarker = styles.Muted.Render("  ")

// ---------------------------------------------------------------------------
// VIEW — LIST MODE
// ---------------------------------------------------------------------------

// Projects renders the complete projects page.
// It shows either the project list or the project detail view
// depending on the ShowProjectDetail flag in the model.
func Projects(m model.Model) string {
	if m.ShowProjectDetail {
		return renderProjectDetail(m)
	}
	return renderProjectList(m)
}

// renderProjectList shows all projects as a selectable list inside a panel.
func renderProjectList(m model.Model) string {
	// Title heading.
	title := styles.Subtitle.Render("PROJECTS")

	// Build each project row.
	rows := make([]string, len(projects))
	for i, p := range projects {
		marker := emptyMarker
		itemStyle := styles.Muted

		if i == m.SelectedProject {
			marker = selectionMarker
			itemStyle = styles.Body
		}

		row := marker + itemStyle.Render(p.name)
		rows[i] = row
	}

	// Stack rows vertically with spacing.
	list := lipgloss.JoinVertical(lipgloss.Left, rows...)

	// Hint text at the bottom.
	hint := styles.Muted.Render("↑/↓ select  ·  enter view details  ·  esc back")

	// Compose inner content.
	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		list,
		"",
		"",
		hint,
	)

	// Wrap in content panel.
	panelWidth := m.Width - 8
	panelHeight := m.Height - 4

	panel := styles.ContentPanel.
		Width(panelWidth).
		Height(panelHeight).
		Align(lipgloss.Left, lipgloss.Top).
		Render(inner)

	return lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, panel)
}

// ---------------------------------------------------------------------------
// VIEW — DETAIL MODE
// ---------------------------------------------------------------------------

// renderProjectDetail shows the full details for the selected project
// with focusable GitHub and Back buttons.
func renderProjectDetail(m model.Model) string {
	if m.SelectedProject < 0 || m.SelectedProject >= len(projects) {
		return renderProjectList(m)
	}

	p := projects[m.SelectedProject]

	// -- Project name --
	name := styles.Title.
		Bold(true).
		Render(p.name)

	// -- Separator --
	sep := styles.Muted.Render("─")

	// -- Description --
	desc := styles.Body.
		Width(60).
		Render(p.description)

	// -- Tech stack --
	techLabel := styles.Accent.Bold(true).Render("Tech Stack")
	techContent := styles.Body.Render(p.tech)

	// -- Features --
	featLabel := styles.Accent.Bold(true).Render("Features")
	featItems := make([]string, len(p.features))
	for i, f := range p.features {
		featItems[i] = styles.Body.Render("  •  " + f)
	}
	featBlock := lipgloss.JoinVertical(lipgloss.Left, featItems...)

	// -- Buttons --
	// Button 0 = GitHub, Button 1 = Back
	buttons := renderDetailButtons(m, p)

	// -- Compose --
	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		name,
		sep,
		"",
		desc,
		"",
		techLabel,
		techContent,
		"",
		featLabel,
		featBlock,
		"",
		"",
		buttons,
	)

	panelWidth := m.Width - 8
	panelHeight := m.Height - 4

	panel := styles.ContentPanel.
		Width(panelWidth).
		Height(panelHeight).
		Align(lipgloss.Left, lipgloss.Top).
		Render(inner)

	return lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, panel)
}

// renderDetailButtons renders the GitHub and Back buttons for project detail view.
func renderDetailButtons(m model.Model, p project) string {
	var btn0, btn1 string

	// GitHub button (index 0)
	githubLabel := "  GitHub  "
	if p.githubURL == "" {
		// No URL yet — show as disabled/muted
		btn0 = styles.Muted.Render("  GitHub  ")
	} else if m.FocusedButton == 0 {
		btn0 = styles.ButtonFocused.Render(githubLabel)
	} else {
		btn0 = styles.ButtonBlurred.Render(githubLabel)
	}

	// Back button (index 1)
	backLabel := "  Back  "
	if m.FocusedButton == 1 {
		btn1 = styles.ButtonFocused.Render(backLabel)
	} else {
		btn1 = styles.ButtonBlurred.Render(backLabel)
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, btn0, "   ", btn1)
}

// ---------------------------------------------------------------------------
// HELPERS
// These are called by main.go's Update function to handle keyboard input.
// ---------------------------------------------------------------------------

// ProjectCount returns the total number of projects.
func ProjectCount() int {
	return len(projects)
}

// ProjectDetailButtonURL returns the URL for a button in project detail view.
// Index 0 = GitHub, Index 1 = Back (returns empty string).
func ProjectDetailButtonURL(projectIndex, buttonIndex int) string {
	if projectIndex < 0 || projectIndex >= len(projects) {
		return ""
	}
	p := projects[projectIndex]
	switch buttonIndex {
	case 0:
		return p.githubURL
	case 1:
		return "" // Back button has no URL
	default:
		return ""
	}
}

// ProjectDetailButtonCount returns the number of buttons in detail view.
func ProjectDetailButtonCount() int {
	return 2 // GitHub + Back
}
