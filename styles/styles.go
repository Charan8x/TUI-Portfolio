// Package styles defines the complete visual design system for the portfolio.
// Every color, text style, border, and spacing value lives here.
// No other file should hardcode a color or padding — import this package instead.
package styles

import "github.com/charmbracelet/lipgloss"

// ---------------------------------------------------------------------------
// COLOR PALETTE
// These are the only colors used across the entire application.
// Changing a value here changes it everywhere — that's the point.
// ---------------------------------------------------------------------------

const (
	ColorBase    = lipgloss.Color("#141414") // main background
	ColorSurface = lipgloss.Color("#1F1F1F") // panel / card backgrounds
	ColorSubtle  = lipgloss.Color("#2E2E2E") // borders and dividers
	ColorMuted   = lipgloss.Color("#666666") // dim text, secondary labels
	ColorText    = lipgloss.Color("#E8E8E8") // primary readable text
	ColorAccent  = lipgloss.Color("#00D4AA") // teal — active, highlighted, interactive
)

// ---------------------------------------------------------------------------
// TYPOGRAPHY
// These styles define how text looks throughout the app.
// Think of these like CSS classes — Title, Subtitle, Body, Muted.
// ---------------------------------------------------------------------------

// Title is used for page headings and your name on the home screen.
var Title = lipgloss.NewStyle().
	Foreground(ColorAccent).
	Bold(true)

// Subtitle is used for section headings within a page.
var Subtitle = lipgloss.NewStyle().
	Foreground(ColorText).
	Bold(true)

// Body is the default readable text style.
var Body = lipgloss.NewStyle().
	Foreground(ColorText)

// Muted is for secondary information — labels, hints, less important text.
var Muted = lipgloss.NewStyle().
	Foreground(ColorMuted)

// Accent is for highlighted inline text — not a heading, just emphasis.
var Accent = lipgloss.NewStyle().
	Foreground(ColorAccent)

// ---------------------------------------------------------------------------
// NAVIGATION BAR STYLES
// The navbar sits at the top of every screen.
// A tab can be in one of two states: active (you're on that page)
// or inactive (you're not on that page).
// ---------------------------------------------------------------------------

// ActiveTab is the currently selected navigation item.
// Uses a teal background and dark text to stand out.
var ActiveTab = lipgloss.NewStyle().
	Foreground(ColorBase).
	Background(ColorAccent).
	Bold(true).
	Padding(0, 2)

// InactiveTab is any navigation item that is not currently selected.
var InactiveTab = lipgloss.NewStyle().
	Foreground(ColorMuted).
	Padding(0, 2)

// ---------------------------------------------------------------------------
// HEADER
// The header bar spans the full terminal width at the very top.
// Left side: branding. Right side: navigation tabs.
// ---------------------------------------------------------------------------

// Header is the full-width top bar container.
// It has a bottom border to separate it from the page content.
var Header = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true).
	BorderForeground(ColorSubtle).
	Padding(0, 1)

// HeaderBrand is the "Sai Charan | Portfolio" text on the left of the header.
var HeaderBrand = lipgloss.NewStyle().
	Foreground(ColorText).
	Bold(true)

// HeaderBrandAccent is the accent-colored part of the brand text.
var HeaderBrandAccent = lipgloss.NewStyle().
	Foreground(ColorAccent).
	Bold(true)

// ---------------------------------------------------------------------------
// LAYOUT & PANELS
// These styles define the containers that hold page content.
// ---------------------------------------------------------------------------

// Page is the outermost content container below the header.
// It provides consistent padding on all pages.
var Page = lipgloss.NewStyle().
	Padding(2, 4)

// Panel is a bordered box used to group related content on a page.
// Used for things like project cards, skill sections, etc.
var Panel = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(ColorSubtle).
	Background(ColorSurface).
	Padding(1, 2)

// ContentPanel is the large centered bordered box on the home page.
// It frames the main content and gives the layout its premium feel.
var ContentPanel = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(ColorSubtle).
	Padding(2, 6)

// ---------------------------------------------------------------------------
// STATUS BAR
// A slim bar pinned to the very bottom of the terminal.
// Shows keyboard hints so users always know how to navigate.
// ---------------------------------------------------------------------------

// StatusBar is the full-width bottom bar container.
var StatusBar = lipgloss.NewStyle().
	Background(ColorSurface).
	Foreground(ColorMuted).
	Padding(0, 2)

// StatusKey is a highlighted key name inside the status bar (e.g. "tab").
var StatusKey = lipgloss.NewStyle().
	Foreground(ColorAccent).
	Background(ColorSurface).
	Bold(true)

// ---------------------------------------------------------------------------
// PAGE NAV STRIP
// A horizontal navigation strip shown inside the content area.
// Lists all pages with the active one highlighted.
// ---------------------------------------------------------------------------

// NavStripItem is the style for the container of the horizontal page nav strip.
var NavStripItem = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Padding(0, 4)

// NavStripActive is the active page item in the nav strip.
var NavStripActive = lipgloss.NewStyle().
	Foreground(ColorBase).
	Background(ColorAccent).
	Bold(true).
	Padding(0, 2)

// NavStripInactive is an inactive page item in the nav strip.
var NavStripInactive = lipgloss.NewStyle().
	Foreground(ColorMuted).
	Padding(0, 2)

// ---------------------------------------------------------------------------
// BUTTONS
// Buttons are selectable elements on the home and projects pages.
// They have two states: focused (the cursor is on them) and blurred.
// ---------------------------------------------------------------------------

// ButtonFocused is a button that currently has keyboard focus.
var ButtonFocused = lipgloss.NewStyle().
	Foreground(ColorBase).
	Background(ColorAccent).
	Bold(true).
	Padding(0, 2)

// ButtonBlurred is a button that does not have keyboard focus.
var ButtonBlurred = lipgloss.NewStyle().
	Foreground(ColorAccent).
	Padding(0, 2)
