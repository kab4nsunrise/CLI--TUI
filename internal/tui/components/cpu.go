package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kab4nsunrise/gomon/internal/sysinfo"
	"github.com/kab4nsunrise/gomon/internal/tui/styles"
)

func RenderCPU(info sysinfo.CPUInfo, width int) string {
	var b strings.Builder

	b.WriteString(styles.PanelTitleStyle.Render("CPU"))
	b.WriteString("\n")

	bar := styles.ProgressBar(info.UsagePercent, 18)
	b.WriteString(fmt.Sprintf("%s %s %5.1f%%\n",
		styles.LabelStyle.Render("Total"),
		bar,
		info.UsagePercent,
	))

	maxCores := 8
	if len(info.PerCore) < maxCores {
		maxCores = len(info.PerCore)
	}
	for i := 0; i < maxCores; i++ {
		pct := info.PerCore[i]
		bar := styles.ProgressBar(pct, 12)
		b.WriteString(fmt.Sprintf("%s %s %5.1f%%\n",
			styles.LabelStyle.Render(fmt.Sprintf("C%-2d", i)),
			bar,
			pct,
		))
	}
	if len(info.PerCore) > maxCores {
		b.WriteString(styles.LabelStyle.Render(fmt.Sprintf("  … +%d cores\n", len(info.PerCore)-maxCores)))
	}

	content := strings.TrimRight(b.String(), "\n")
	return styles.PanelStyle.Width(width).Render(content)
}

func RenderCPUCompact(info sysinfo.CPUInfo) string {
	bar := styles.ProgressBar(info.UsagePercent, 15)
	return lipgloss.JoinHorizontal(lipgloss.Left,
		styles.LabelStyle.Render("CPU "),
		bar,
		fmt.Sprintf(" %.1f%%", info.UsagePercent),
	)
}
