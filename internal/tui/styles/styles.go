package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	colorPrimary = lipgloss.Color("#7C3AED")
	colorAccent  = lipgloss.Color("#22D3EE")
	colorSuccess = lipgloss.Color("#34D399")
	colorWarning = lipgloss.Color("#FBBF24")
	colorDanger  = lipgloss.Color("#F87171")
	colorMuted   = lipgloss.Color("#6B7280")
	colorText    = lipgloss.Color("#E5E7EB")
	colorBg      = lipgloss.Color("#111827")
	colorPanelBg = lipgloss.Color("#1F2937")

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			MarginBottom(1)

	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Padding(0, 1).
			MarginRight(1)

	PanelTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorAccent)

	LabelStyle = lipgloss.NewStyle().Foreground(colorMuted)
	ValueStyle = lipgloss.NewStyle().Foreground(colorText).Bold(true)
	HelpStyle  = lipgloss.NewStyle().Foreground(colorMuted)

	StatusBarStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Background(colorPanelBg).
			Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorAccent).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(colorMuted)

	SelectedRowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#111827")).
				Background(colorPrimary).
				Bold(true)

	NormalRowStyle = lipgloss.NewStyle().Foreground(colorText)

	BarLowStyle  = lipgloss.NewStyle().Foreground(colorSuccess)
	BarMedStyle  = lipgloss.NewStyle().Foreground(colorWarning)
	BarHighStyle = lipgloss.NewStyle().Foreground(colorDanger)

	ErrorStyle = lipgloss.NewStyle().Foreground(colorDanger).Bold(true)

	ConfirmStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorDanger).
			Padding(1, 2).
			Foreground(colorText)
)

func ProgressBar(percent float64, width int) string {
	if width <= 0 {
		width = 20
	}
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	filled := int(percent / 100 * float64(width))
	if filled > width {
		filled = width
	}

	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	style := BarLowStyle
	switch {
	case percent >= 85:
		style = BarHighStyle
	case percent >= 60:
		style = BarMedStyle
	}

	return style.Render(bar)
}
