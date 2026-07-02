// Package model defines the application state for the portfolio TUI.
// In Bubble Tea, the model is the single source of truth —
// every piece of information needed to render the UI lives here.
package model

// ---------------------------------------------------------------------------
// PAGE TYPE
// Instead of using raw integers (0, 1, 2, 3) to represent pages,
// we define a named type. This makes the code self-documenting.
// ---------------------------------------------------------------------------

// Page represents which page of the portfolio is currently active.
type Page int

const (
	PageHome     Page = iota // iota starts at 0 and increments automatically
	PageAbout                // 1
	PageProjects             // 2
	PageContact              // 3
)

// PageCount is the total number of pages.
// Used when wrapping navigation (going right from Contact loops back to Home).
const PageCount = 4

// ---------------------------------------------------------------------------
// MODEL
// This struct is the entire application state.
// Bubble Tea will pass this around between Update and View on every event.
// ---------------------------------------------------------------------------

// Model holds all state required to render the portfolio at any given moment.
type Model struct {
	// CurrentPage is the page currently being displayed.
	CurrentPage Page

	// FocusedButton tracks which button has keyboard focus on the current page.
	// Each page resets this to 0 when you navigate to it.
	FocusedButton int

	// SelectedProject is the index of the project selected on the Projects page.
	SelectedProject int

	// ShowProjectDetail controls whether the project detail view is visible.
	// When true, the Projects page renders a detail panel instead of the list.
	ShowProjectDetail bool

	// Width and Height are the current terminal dimensions.
	// Bubble Tea sends a WindowSizeMsg whenever the terminal is resized.
	// We store the values here so View can use them to size content correctly.
	Width  int
	Height int
}

// ---------------------------------------------------------------------------
// CONSTRUCTOR
// New returns a Model initialised with sensible default values.
// This is the idiomatic Go way to construct a struct with defaults —
// a plain function named New that returns the type.
// ---------------------------------------------------------------------------

// New returns the initial application model.
// The app always starts on the Home page with the first button focused.
func New() Model {
	return Model{
		CurrentPage:      PageHome,
		FocusedButton:    0,
		SelectedProject:  0,
		ShowProjectDetail: false,
		// Width and Height start at 0.
		// They will be set correctly the moment Bubble Tea sends
		// the first WindowSizeMsg, which happens automatically on startup.
		Width:  0,
		Height: 0,
	}
}

// ---------------------------------------------------------------------------
// HELPERS
// Small pure functions that answer questions about the model.
// Keeping logic here means pages and components don't need to think.
// ---------------------------------------------------------------------------

// NavItems returns the display labels for the navigation bar,
// in the order they appear left to right.
func NavItems() []string {
	return []string{"Home", "About", "Projects", "Contact"}
}
