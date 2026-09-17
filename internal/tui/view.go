package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kab4nsunrise/gomon/internal/tui/components"
	"github.com/kab4nsunrise/gomon/internal/tui/styles"
)

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.width == 0 {
		return "Loading…"
	}
	if m.showHelp {
		return components.RenderHelp(m.width, m.height)
	}
	if m.confirmKill && m.cursor >= 0 && m.cursor < len(m.procs) {
		p := m.procs[m.cursor]
		return components.RenderConfirm(p.PID, p.Name, m.width)
	}

	var b strings.Builder

	title := styles.TitleStyle.Render("⚡ gomon — System Monitor")
	b.WriteString(title)
	b.WriteString("\n")

	panelWidth := (m.width - 6) / 2
	if panelWidth < 30 {
		panelWidth = 30
	}

	cpuPanel := components.RenderCPU(m.cpu, panelWidth)
	memPanel := components.RenderMemory(m.mem, panelWidth)
	top := lipgloss.JoinHorizontal(lipgloss.Top, cpuPanel, memPanel)
	b.WriteString(top)
	b.WriteString("\n")

	procHeight := m.height - 18
	if procHeight < 8 {
		procHeight = 8
	}
	procWidth := m.width - 4
	if procWidth < 40 {
		procWidth = 40
	}

	procsPanel := components.RenderProcesses(m.procs, m.cursor, procWidth, procHeight, m.sortBy)
	b.WriteString(procsPanel)
	b.WriteString("\n")

	status := components.RenderStatusBar(m.width, m.statusMsg, m.lastUpdate)
	b.WriteString(status)

	return b.String()
}
