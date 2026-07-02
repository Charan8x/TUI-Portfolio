// Package pages contains the full-screen view functions for each page.
package pages

import (
	"fmt"
	"strings"

	"ssh-portfolio/model"
	"ssh-portfolio/styles"

	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------------------------
// CONTENT
// ---------------------------------------------------------------------------

const (
	homeName  = "Sai Charan"
	homeTitle = "Full Stack Developer & Cloud Enthusiast"
	homeIntro = "Building efficient web applications · Exploring cloud infrastructure"
)

var homeButtons = []struct {
	label string
	url   string
}{
	{label: "GitHub", url: "https://github.com/Charan8x"},
	{label: "LinkedIn", url: "https://www.linkedin.com/in/saicharan777/"},
	{label: "Email", url: "mailto:bunnyappasani5@gmail.com"},
	{label: "Resume", url: "https://drive.google.com/file/d/1ADwv2OfzWcr5WKGlBNDaLM3cTQWXXp7g/view"},
}

// ---------------------------------------------------------------------------
// ASCII BRANDING — SC block logo
// ---------------------------------------------------------------------------

const asciiLogo = `
███████╗ ██████╗
██╔════╝██╔════╝
███████╗██║     
╚════██║██║     
███████║╚██████╗
╚══════╝ ╚═════╝`

// ---------------------------------------------------------------------------
// VIEW
// ---------------------------------------------------------------------------

// Home renders the complete home page.
// Layout:
//
//	[header]           ← rendered by main view, not here
//	[                ]
//	[  ╭───────────╮ ]
//	[  │  SC logo  │ ]
//	[  │  name     │ ]
//	[  │  title    │ ]
//	[  │  buttons  │ ]
//	[  ╰───────────╯ ]
//	[                ]
//	[status bar]       ← rendered by main view, not here
func Home(m model.Model) string {
	// -- ASCII Logo --
	logo := styles.Accent.Render(asciiLogo)

	// -- Name --
	name := styles.Title.
		Bold(true).
		Render(homeName)

	// -- Title / subtitle --
	title := styles.Muted.Render(homeTitle)

	// -- Intro tagline --
	intro := styles.Body.Render(homeIntro)

	// -- Divider --
	// A subtle horizontal rule inside the panel to separate logo from text.
	dividerWidth := 36
	divider := styles.Muted.Render(strings.Repeat("─", dividerWidth))

	// -- Buttons --
	buttons := renderHomeButtons(m)

	// -- Page title --
	pageTitle := styles.Muted.Render("HOME")

	// -- Inner content block --
	// Everything stacked vertically, centered.
	inner := lipgloss.JoinVertical(
		lipgloss.Center,
		pageTitle,
		"",
		logo,
		"",
		divider,
		"",
		name,
		title,
		"",
		intro,
		"",
		buttons,
	)

	// -- Content panel --
	// The large bordered box that frames the home content.
	// Height is calculated to fill the screen between header and status bar.
	// Header ≈ 3 lines, status bar = 1 line, so subtract 4.
	panelHeight := m.Height - 4
	panelWidth := m.Width - 8 // leave a small margin on each side

	panel := styles.ContentPanel.
		Width(panelWidth).
		Height(panelHeight).
		Align(lipgloss.Center, lipgloss.Top).
		Render(inner)

	// -- Place panel with horizontal centering --
	return lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, panel)
}

// renderHomeButtons builds two rows of two buttons.
func renderHomeButtons(m model.Model) string {
	rendered := make([]string, len(homeButtons))

	for i, btn := range homeButtons {
		label := fmt.Sprintf("  %s  ", btn.label)
		if i == m.FocusedButton {
			rendered[i] = styles.ButtonFocused.Render(label)
		} else {
			rendered[i] = styles.ButtonBlurred.Render(label)
		}
	}

	row1 := lipgloss.JoinHorizontal(lipgloss.Top, rendered[0], "   ", rendered[1])
	row2 := lipgloss.JoinHorizontal(lipgloss.Top, rendered[2], "   ", rendered[3])

	return lipgloss.JoinVertical(lipgloss.Center, row1, "", row2)
}

// HomeButtonURL returns the URL for a given button index.
func HomeButtonURL(index int) string {
	if index < 0 || index >= len(homeButtons) {
		return ""
	}
	return homeButtons[index].url
}

// HomeButtonCount returns the total number of buttons on the home page.
func HomeButtonCount() int {
	return len(homeButtons)
}
