package components

import (
	"fmt"
	"strings"

	"github.com/kab4nsunrise/gomon/internal/sysinfo"
	"github.com/kab4nsunrise/gomon/internal/tui/styles"
)

func RenderMemory(info sysinfo.MemInfo, width int) string {
	var b strings.Builder

	b.WriteString(styles.PanelTitleStyle.Render("Memory"))
	b.WriteString("\n")

	bar := styles.ProgressBar(info.UsedPercent, 18)
	b.WriteString(fmt.Sprintf("%s %s %5.1f%%\n",
		styles.LabelStyle.Render("RAM "),
		bar,
		info.UsedPercent,
	))
	b.WriteString(fmt.Sprintf("  %s / %s\n",
		styles.ValueStyle.Render(sysinfo.FormatBytes(info.Used)),
		styles.LabelStyle.Render(sysinfo.FormatBytes(info.Total)),
	))

	if info.SwapTotal > 0 {
		swapBar := styles.ProgressBar(info.SwapPercent, 18)
		b.WriteString(fmt.Sprintf("%s %s %5.1f%%\n",
			styles.LabelStyle.Render("Swap"),
			swapBar,
			info.SwapPercent,
		))
		b.WriteString(fmt.Sprintf("  %s / %s\n",
			styles.ValueStyle.Render(sysinfo.FormatBytes(info.SwapUsed)),
			styles.LabelStyle.Render(sysinfo.FormatBytes(info.SwapTotal)),
		))
	} else {
		b.WriteString(styles.LabelStyle.Render("Swap  (none)\n"))
	}

	content := strings.TrimRight(b.String(), "\n")
	return styles.PanelStyle.Width(width).Render(content)
}
