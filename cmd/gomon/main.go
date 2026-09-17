package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/kab4nsunrise/gomon/internal/config"
	"github.com/kab4nsunrise/gomon/internal/sysinfo"
	"github.com/kab4nsunrise/gomon/internal/tui"
)

var (
	version = "0.1.0"
	cfg     = config.Default()
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gomon",
		Short: "System & process monitor with TUI",
		Long: `gomon is a terminal system monitor inspired by htop/btop.

Run without arguments to launch the interactive TUI.
Use subcommands for quick one-shot information.`,
		RunE: runTUI,
	}

	rootCmd.PersistentFlags().DurationVar(&cfg.RefreshInterval, "interval", cfg.RefreshInterval, "refresh interval")
	rootCmd.PersistentFlags().IntVar(&cfg.ProcessLimit, "limit", cfg.ProcessLimit, "max processes to show")

	rootCmd.AddCommand(
		versionCmd(),
		cpuCmd(),
		memCmd(),
		topCmd(),
		killCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runTUI(cmd *cobra.Command, args []string) error {
	m := tui.New(cfg)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}
	return nil
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gomon %s\n", version)
		},
	}
}

func cpuCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cpu",
		Short: "Show current CPU usage",
		RunE: func(cmd *cobra.Command, args []string) error {
			info, err := sysinfo.GetCPU()
			if err != nil {
				return err
			}
			fmt.Printf("CPU: %.1f%%  (%d cores)\n", info.UsagePercent, info.Cores)
			if info.ModelName != "" {
				fmt.Printf("Model: %s\n", info.ModelName)
			}
			for i, pct := range info.PerCore {
				fmt.Printf("  Core %d: %.1f%%\n", i, pct)
			}
			return nil
		},
	}
}

func memCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mem",
		Short: "Show current memory usage",
		RunE: func(cmd *cobra.Command, args []string) error {
			info, err := sysinfo.GetMem()
			if err != nil {
				return err
			}
			fmt.Printf("Memory: %s / %s (%.1f%%)\n",
				sysinfo.FormatBytes(info.Used),
				sysinfo.FormatBytes(info.Total),
				info.UsedPercent,
			)
			if info.SwapTotal > 0 {
				fmt.Printf("Swap:   %s / %s (%.1f%%)\n",
					sysinfo.FormatBytes(info.SwapUsed),
					sysinfo.FormatBytes(info.SwapTotal),
					info.SwapPercent,
				)
			}
			return nil
		},
	}
}

func topCmd() *cobra.Command {
	var n int
	cmd := &cobra.Command{
		Use:   "top",
		Short: "Show top processes by CPU",
		RunE: func(cmd *cobra.Command, args []string) error {
			procs, err := sysinfo.GetProcesses(n, sysinfo.SortByCPU)
			if err != nil {
				return err
			}
			fmt.Printf("%-8s %6s %6s %-12s %s\n", "PID", "CPU%", "MEM%", "USER", "NAME")
			for _, p := range procs {
				fmt.Printf("%-8d %5.1f%% %5.1f%% %-12s %s\n",
					p.PID, p.CPUPercent, p.MemPercent, p.Username, p.Name)
			}
			return nil
		},
	}
	cmd.Flags().IntVarP(&n, "number", "n", 10, "number of processes")
	return cmd
}

func killCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "kill [pid]",
		Short: "Kill a process by PID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var pid int32
			if _, err := fmt.Sscanf(args[0], "%d", &pid); err != nil {
				return fmt.Errorf("invalid PID: %s", args[0])
			}
			if err := sysinfo.KillProcess(pid); err != nil {
				return err
			}
			fmt.Printf("Process %d killed\n", pid)
			return nil
		},
	}
}
