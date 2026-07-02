// Main entry point for the SSH Portfolio TUI.
// Initialises the Bubble Tea program and starts the event loop.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"ssh-portfolio/components"
	"ssh-portfolio/model"
	"ssh-portfolio/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------------------------
// PROGRAM MODEL
// TeaModel wraps model.Model and implements the tea.Model interface
// (Init, Update, View). This keeps the pure state in the model package
// and the event handling logic here in main.go.
// ---------------------------------------------------------------------------

// TeaModel is the Bubble Tea model for the application.
// It embeds model.Model to inherit all state fields.
type TeaModel struct {
	model.Model
}

// Init is called when the program starts. Returns any initial commands.
func (m TeaModel) Init() tea.Cmd {
	return nil
}

// Update is called on every event (key press, resize, etc.).
func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// -- Window resize --
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	// -- Keyboard input --
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "left":
			// Previous page
			m.FocusedButton = 0
			m.ShowProjectDetail = false
			m.CurrentPage = (m.CurrentPage + model.PageCount - 1) % model.PageCount
			return m, nil

		case "right":
			// Next page
			m.FocusedButton = 0
			m.ShowProjectDetail = false
			m.CurrentPage = (m.CurrentPage + 1) % model.PageCount
			return m, nil

		// -- Up/Down: project list navigation (Projects page only) --
		case "up":
			if m.CurrentPage == model.PageProjects && !m.ShowProjectDetail {
				if m.SelectedProject > 0 {
					m.SelectedProject--
				}
			}
			return m, nil

		case "down":
			if m.CurrentPage == model.PageProjects && !m.ShowProjectDetail {
				if m.SelectedProject < pages.ProjectCount()-1 {
					m.SelectedProject++
				}
			}
			return m, nil

		// -- Tab: cycle focus within current page --
		case "tab":
			switch {
			case m.CurrentPage == model.PageProjects && !m.ShowProjectDetail:
				// Project list: cycle through projects
				m.SelectedProject = (m.SelectedProject + 1) % pages.ProjectCount()
			case m.CurrentPage == model.PageProjects && m.ShowProjectDetail:
				// Project detail: cycle GitHub/Back buttons
				m.FocusedButton = (m.FocusedButton + 1) % pages.ProjectDetailButtonCount()
			case m.CurrentPage == model.PageHome:
				m.FocusedButton = (m.FocusedButton + 1) % pages.HomeButtonCount()
			case m.CurrentPage == model.PageContact:
				m.FocusedButton = (m.FocusedButton + 1) % pages.ContactItemCount()
			default:
				m.FocusedButton = 0
			}
			return m, nil

		case "enter":
			return m.handleEnter()

		case "esc":
			if m.CurrentPage == model.PageProjects && m.ShowProjectDetail {
				m.ShowProjectDetail = false
				m.FocusedButton = 0
			}
			return m, nil
		}
	}

	return m, nil
}

// handleEnter processes the Enter key press based on the current page.
func (m TeaModel) handleEnter() (tea.Model, tea.Cmd) {
	switch m.CurrentPage {

	case model.PageHome:
		url := pages.HomeButtonURL(m.FocusedButton)
		if url == "" {
			return m, nil
		}
		return m, openURL(url)

	case model.PageProjects:
		if m.ShowProjectDetail {
			url := pages.ProjectDetailButtonURL(m.SelectedProject, m.FocusedButton)
			if url == "" {
				if m.FocusedButton == 1 {
					m.ShowProjectDetail = false
					m.FocusedButton = 0
				}
				return m, nil
			}
			return m, openURL(url)
		}
		m.ShowProjectDetail = true
		m.FocusedButton = 0
		return m, nil

	case model.PageContact:
		url := pages.ContactItemURL(m.FocusedButton)
		if url == "" {
			return m, nil
		}
		return m, openURL(url)

	default:
		return m, nil
	}
}

// openURL returns a command that opens the given URL in the default browser.
// Uses the OS-specific open command: rundll32 on Windows, open on macOS,
// xdg-open on Linux.
func openURL(url string) tea.Cmd {
	return func() tea.Msg {
		switch runtime.GOOS {
		case "windows":
			exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			exec.Command("open", url).Start()
		default:
			exec.Command("xdg-open", url).Start()
		}
		return nil
	}
}

// View is called on every render cycle.
// Composes header + page nav strip + page content + status bar.
func (m TeaModel) View() string {
	if m.Width == 0 {
		return ""
	}

	header := components.Render(m.Model)
	pageNav := components.RenderPageNav(m.Model)

	var page string
	switch m.CurrentPage {
	case model.PageHome:
		page = pages.Home(m.Model)
	case model.PageAbout:
		page = pages.About(m.Model)
	case model.PageProjects:
		page = pages.Projects(m.Model)
	case model.PageContact:
		page = pages.Contact(m.Model)
	}

	statusBar := components.RenderStatusBar(m.Model)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		pageNav,
		page,
		statusBar,
	)
}

// ---------------------------------------------------------------------------
// MAIN
// Entry point. Creates the Bubble Tea program and starts the event loop.
// ---------------------------------------------------------------------------

func main() {
	p := tea.NewProgram(
		TeaModel{Model: model.New()},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	_ = m
}
