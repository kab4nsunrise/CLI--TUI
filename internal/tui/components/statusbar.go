package components

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/kab4nsunrise/gomon/internal/tui/styles"
)

func RenderStatusBar(width int, message string, lastUpdate time.Time) string {
	help := "q:quit  ↑↓/j:nav  c:cpu  m:mem  p:pid  n:name  k:kill  r:refresh  ?:help"
	if message != "" {
		help = message
	}

	timeStr := lastUpdate.Format("15:04:05")
	left := styles.HelpStyle.Render(help)
	right := styles.LabelStyle.Render(timeStr)

	gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	line := left + lipgloss.NewStyle().Width(gap).Render("") + right
	return styles.StatusBarStyle.Width(width).Render(line)
}

func RenderConfirm(pid int32, name string, width int) string {
	msg := fmt.Sprintf("Kill process %d (%s)?\n\n[y] Yes    [n] No", pid, name)
	box := styles.ConfirmStyle.Render(msg)
	return lipgloss.Place(width, 8, lipgloss.Center, lipgloss.Center, box)
}

func RenderHelp(width, height int) string {
	help := `
  gomon — System Monitor

  Navigation
    ↑         Move up
    ↓ / j     Move down
    PgUp/PgDn Page up/down

  Sorting
    c         Sort by CPU
    m         Sort by Memory
    p         Sort by PID
    n         Sort by Name

  Actions
    k / K     Kill selected process
    r         Force refresh
    ?         Toggle this help
    q / Esc   Quit

  Press any key to close help
`
	box := styles.PanelStyle.
		BorderForeground(lipgloss.Color("#7C3AED")).
		Padding(1, 2).
		Render(help)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
