package pages

import (
	"ssh-portfolio/model"
	"ssh-portfolio/styles"

	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------------------------
// CONTENT
// ---------------------------------------------------------------------------

type contactItem struct {
	label string
	value string
	url   string
}

var contactItems = []contactItem{
	{label: "Email", value: "bunnyappasani5@gmail.com", url: "mailto:bunnyappasani5@gmail.com"},
	{label: "LinkedIn", value: "linkedin.com/in/saicharan777", url: "https://www.linkedin.com/in/saicharan777/"},
	{label: "GitHub", value: "github.com/Charan8x", url: "https://github.com/Charan8x"},
	{label: "Availability", value: "Currently working as intern", url: ""},
}

// ---------------------------------------------------------------------------
// VIEW
// ---------------------------------------------------------------------------

// Contact renders the complete contact page.
func Contact(m model.Model) string {
	// Title.
	title := styles.Subtitle.Render("CONTACT")

	// Build each contact row.
	rows := make([]string, len(contactItems))
	for i, item := range contactItems {
		var row string

		if i == m.FocusedButton && item.url != "" {
			// Focused and actionable — show in focused button style.
			line := item.label + "  " + item.value
			row = styles.ButtonFocused.Render(line)
		} else if item.url != "" {
			// Not focused but has a URL — show as clickable hint.
			label := styles.Accent.Bold(true).Render(item.label)
			value := styles.Body.Render(item.value)
			row = label + "  " + value
		} else {
			// No URL — plain display (Availability).
			label := styles.Accent.Bold(true).Render(item.label)
			value := styles.Body.Render(item.value)
			row = label + "  " + value
		}

		rows[i] = row
	}

	// Stack items vertically.
	items := lipgloss.JoinVertical(lipgloss.Left, rows...)

	// Hint text at bottom.
	hint := styles.Muted.Render("tab focus  ·  enter open  ·  esc back")

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		items,
		"",
		"",
		hint,
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

// ContactItemURL returns the URL for a given contact item index.
func ContactItemURL(index int) string {
	if index < 0 || index >= len(contactItems) {
		return ""
	}
	return contactItems[index].url
}

// ContactItemCount returns the total number of contact items.
func ContactItemCount() int {
	return len(contactItems)
}
