package pages

import (
	"ssh-portfolio/model"
	"ssh-portfolio/styles"

	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------------------------
// CONTENT
// All about page content is defined here as constants.
// Update this section to change what appears on the About page.
// ---------------------------------------------------------------------------

const aboutIntro = "Full stack developer with a passion for building efficient and\nscalable web applications. Experienced in both backend and frontend\ndevelopment, with a growing focus on cloud infrastructure\nand distributed systems."

// aboutSection represents a single labeled section on the about page.
// Each section has a title and one or more lines of content.
type aboutSection struct {
	title   string
	content string
}

// sections defines every block that appears on the About page, in order.
var aboutSections = []aboutSection{
	{
		title:   "ABOUT ME",
		content: aboutIntro,
	},
	{
		title: "EDUCATION",
		content: "B.Tech  ·  SRM University AP  ·  2024–2028  ·  CGPA 7.1\n12th    ·  MINDS (CBSE)        ·  2022–2024  ·  75%",
	},
	{
		title:   "CORE CONCEPTS",
		content: "Database Management (DBMS)  ·  Object Oriented Programming (OOP)  ·  Data Structures & Algorithms (DSA)  ·  Deep Learning",
	},
	{
		title:   "LANGUAGES",
		content: "Python  ·  C  ·  C++  ·  Go  ·  JavaScript  ·  HTML & CSS",
	},
	{
		title:   "FRAMEWORKS & TOOLS",
		content: "Flask  ·  FastAPI  ·  Tailwind CSS  ·  Node.js  ·  React  ·  Bubble Tea  ·  PyTorch",
	},
	{
		title:   "INTERESTS",
		content: "Cloud Infrastructure  ·  VFX & Compositing  ·  CGI  ·  3D Modelling",
	},
}

// ---------------------------------------------------------------------------
// VIEW
// ---------------------------------------------------------------------------

// About renders the complete about page and returns it as a string.
// Layout: single column, sections stacked top to bottom inside a content panel.
func About(m model.Model) string {
	// Build each section block and collect them.
	blocks := make([]string, 0, len(aboutSections)*4)

	for i, section := range aboutSections {
		// Section title in accent color.
		title := styles.Accent.Bold(true).Render(section.title)

		// Divider line beneath the title — same width as the title text.
		dividerLen := len(section.title)
		divider := styles.Muted.Render(repeatRune('─', dividerLen))

		// Section body in normal body text.
		content := styles.Body.Render(section.content)

		// Stack title, divider, content vertically.
		block := lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			divider,
			content,
		)

		blocks = append(blocks, block)

		// Add spacing between sections, but not after the last one.
		if i < len(aboutSections)-1 {
			blocks = append(blocks, "")
		}
	}

	// Page title.
	pageTitle := styles.Subtitle.Render("ABOUT")

	// Combine page title with section blocks into a single slice.
	allBlocks := append([]string{pageTitle, ""}, blocks...)

	// Stack all section blocks into one content string.
	inner := lipgloss.JoinVertical(lipgloss.Left, allBlocks...)

	// Wrap in the content panel — same panel style used on home page
	// for visual consistency across the app.
	panelWidth := m.Width - 8
	panelHeight := m.Height - 4

	panel := styles.ContentPanel.
		Width(panelWidth).
		Height(panelHeight).
		Align(lipgloss.Left, lipgloss.Top).
		Render(inner)

	return lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, panel)
}

// repeatRune returns a string of n repetitions of the given rune.
// Used to draw divider lines that match section title lengths.
func repeatRune(r rune, n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = r
	}
	return string(result)
}
