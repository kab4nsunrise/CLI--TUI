package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kab4nsunrise/gomon/internal/config"
	"github.com/kab4nsunrise/gomon/internal/sysinfo"
)

type tickMsg time.Time

type dataMsg struct {
	CPU   sysinfo.CPUInfo
	Mem   sysinfo.MemInfo
	Procs []sysinfo.ProcessInfo
	Err   error
}

type Model struct {
	cfg config.Config

	cpu   sysinfo.CPUInfo
	mem   sysinfo.MemInfo
	procs []sysinfo.ProcessInfo

	cursor      int
	sortBy      sysinfo.SortBy
	width       int
	height      int
	lastUpdate  time.Time
	statusMsg   string
	showHelp    bool
	confirmKill bool
	err         error
	quitting    bool
}

func New(cfg config.Config) Model {
	return Model{
		cfg:        cfg,
		sortBy:     sysinfo.SortByCPU,
		lastUpdate: time.Now(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(m.cfg.RefreshInterval),
		fetchDataCmd(m.cfg.ProcessLimit, m.sortBy),
	)
}

func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func fetchDataCmd(limit int, sortBy sysinfo.SortBy) tea.Cmd {
	return func() tea.Msg {
		cpu, errCPU := sysinfo.GetCPU()
		mem, errMem := sysinfo.GetMem()
		procs, errProc := sysinfo.GetProcesses(limit, sortBy)

		var err error
		if errCPU != nil {
			err = errCPU
		} else if errMem != nil {
			err = errMem
		} else if errProc != nil {
			err = errProc
		}

		return dataMsg{
			CPU:   cpu,
			Mem:   mem,
			Procs: procs,
			Err:   err,
		}
	}
}
