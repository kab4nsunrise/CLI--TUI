package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kab4nsunrise/gomon/internal/sysinfo"
	"github.com/kab4nsunrise/gomon/internal/tui/styles"
)

const (
	colPID  = 8
	colCPU  = 7
	colMem  = 7
	colUser = 12
	colName = 24
)

func RenderProcesses(procs []sysinfo.ProcessInfo, cursor int, width, height int, sortBy sysinfo.SortBy) string {
	var b strings.Builder

	cpuH, memH, pidH, nameH := "CPU%", "MEM%", "PID", "NAME"
	switch sortBy {
	case sysinfo.SortByCPU:
		cpuH = "CPU%▼"
	case sysinfo.SortByMem:
		memH = "MEM%▼"
	case sysinfo.SortByPID:
		pidH = "PID▼"
	case sysinfo.SortByName:
		nameH = "NAME▼"
	}

	header := fmt.Sprintf("%-*s %-*s %-*s %-*s %s",
		colPID, pidH,
		colCPU, cpuH,
		colMem, memH,
		colUser, "USER",
		nameH,
	)
	b.WriteString(styles.HeaderStyle.Render(header))
	b.WriteString("\n")

	maxRows := height - 2
	if maxRows < 1 {
		maxRows = 10
	}
	if len(procs) < maxRows {
		maxRows = len(procs)
	}

	for i := 0; i < maxRows; i++ {
		p := procs[i]
		name := p.Name
		if len(name) > colName {
			name = name[:colName-1] + "…"
		}
		user := p.Username
		if len(user) > colUser {
			user = user[:colUser-1] + "…"
		}

		row := fmt.Sprintf("%-*d %5.1f%% %5.1f%% %-*s %s",
			colPID, p.PID,
			p.CPUPercent,
			p.MemPercent,
			colUser, user,
			name,
		)

		if i == cursor {
			b.WriteString(styles.SelectedRowStyle.Width(width - 4).Render(row))
		} else {
			b.WriteString(styles.NormalRowStyle.Render(row))
		}
		b.WriteString("\n")
	}

	content := strings.TrimRight(b.String(), "\n")
	title := styles.PanelTitleStyle.Render(fmt.Sprintf("Processes (%d)", len(procs)))
	return styles.PanelStyle.Width(width).Render(lipgloss.JoinVertical(lipgloss.Left, title, content))
}
