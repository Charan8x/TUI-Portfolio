// Package components contains reusable UI elements used across multiple pages.
// Each component is a pure function — it takes data in, returns a string out.
// No state, no side effects.
package components

import (
	"strings"

	"ssh-portfolio/model"
	"ssh-portfolio/styles"

	"github.com/charmbracelet/lipgloss"
)

// separator is the visual divider rendered between navigation tabs.
var separator = styles.Muted.Render("  ")

// Render builds and returns the full header bar as a styled string.
// Layout: "Sai Charan | Portfolio" on the left, nav tabs on the right.
// The header spans the full terminal width.
func Render(m model.Model) string {
	// -- Left side: branding --
	// "Sai Charan" in accent color, "| Portfolio" in normal text.
	brand := lipgloss.JoinHorizontal(
		lipgloss.Center,
		styles.HeaderBrandAccent.Render("Sai Charan"),
		styles.HeaderBrand.Render(" | Portfolio"),
	)

	// -- Right side: navigation tabs --
	items := model.NavItems()
	tabs := make([]string, len(items))

	for i, item := range items {
		if model.Page(i) == m.CurrentPage {
			tabs[i] = styles.ActiveTab.Render(item)
		} else {
			tabs[i] = styles.InactiveTab.Render(item)
		}
	}

	// Join tabs horizontally with a small space between them.
	tabRow := strings.Join(tabs, separator)

	// -- Compose the full header --
	// We need branding on the left and tabs on the right.
	// The trick: calculate the space between them to push tabs to the right edge.
	//
	// Total width = brand width + gap + tabs width
	// gap = m.Width - brand width - tabs width - header horizontal padding (2)
	brandWidth := lipgloss.Width(brand)
	tabsWidth := lipgloss.Width(tabRow)
	padding := 2 // matches Header style Padding(0, 1) × 2 sides

	gapWidth := m.Width - brandWidth - tabsWidth - padding
	if gapWidth < 1 {
		gapWidth = 1
	}

	gap := strings.Repeat(" ", gapWidth)

	// Combine brand + gap + tabs into one full-width row.
	row := brand + gap + tabRow

	// Wrap in the Header container style.
	return styles.Header.
		Width(m.Width - padding).
		Render(row)
}

// RenderPageNav builds a horizontal navigation strip showing all pages.
// The active page is highlighted with the accent background.
// This is displayed at the top of the content area on every page.
func RenderPageNav(m model.Model) string {
	items := model.NavItems()
	tabs := make([]string, len(items))

	for i, item := range items {
		if model.Page(i) == m.CurrentPage {
			tabs[i] = styles.NavStripActive.Render(item)
		} else {
			tabs[i] = styles.NavStripInactive.Render(item)
		}
	}

	// Use a dot separator between items.
	sep := styles.Muted.Render(" · ")
	row := strings.Join(tabs, sep)

	return styles.NavStripItem.
		Width(m.Width).
		Render(row)
}

// RenderStatusBar builds the bottom status bar with keyboard hints.
// It is pinned to the bottom of every page.
func RenderStatusBar(m model.Model) string {
	// Each hint is: key in accent + description in muted
	hints := []string{
		styles.StatusKey.Render("←/→") + styles.StatusBar.Render(" navigate"),
		styles.StatusKey.Render("tab") + styles.StatusBar.Render(" focus"),
		styles.StatusKey.Render("enter") + styles.StatusBar.Render(" open"),
		styles.StatusKey.Render("esc") + styles.StatusBar.Render(" back"),
		styles.StatusKey.Render("ctrl+c") + styles.StatusBar.Render(" quit"),
	}

	content := strings.Join(hints, styles.StatusBar.Render("  ·  "))

	return styles.StatusBar.
		Width(m.Width).
		Render(content)
}
