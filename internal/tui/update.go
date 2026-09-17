package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kab4nsunrise/gomon/internal/sysinfo"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tickMsg:
		return m, tea.Batch(
			tickCmd(m.cfg.RefreshInterval),
			fetchDataCmd(m.cfg.ProcessLimit, m.sortBy),
		)
	case dataMsg:
		if msg.Err != nil {
			m.err = msg.Err
			m.statusMsg = fmt.Sprintf("Error: %v", msg.Err)
		} else {
			m.err = nil
			m.cpu = msg.CPU
			m.mem = msg.Mem
			m.procs = msg.Procs
			m.lastUpdate = time.Now()
			if m.cursor >= len(m.procs) && len(m.procs) > 0 {
				m.cursor = len(m.procs) - 1
			}
			if m.cursor < 0 {
				m.cursor = 0
			}
		}
		return m, nil
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.showHelp {
		m.showHelp = false
		return m, nil
	}

	if m.confirmKill {
		switch msg.String() {
		case "y", "Y":
			if m.cursor >= 0 && m.cursor < len(m.procs) {
				p := m.procs[m.cursor]
				if err := sysinfo.KillProcess(p.PID); err != nil {
					m.statusMsg = fmt.Sprintf("Failed to kill %d: %v", p.PID, err)
				} else {
					m.statusMsg = fmt.Sprintf("Killed process %d (%s)", p.PID, p.Name)
				}
			}
			m.confirmKill = false
			return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
		case "n", "N", "esc":
			m.confirmKill = false
			m.statusMsg = "Kill cancelled"
			return m, nil
		}
		return m, nil
	}

	switch msg.String() {
	case "q", "ctrl+c", "esc":
		m.quitting = true
		return m, tea.Quit
	case "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.procs)-1 {
			m.cursor++
		}
	case "pgup":
		m.cursor -= 10
		if m.cursor < 0 {
			m.cursor = 0
		}
	case "pgdown":
		m.cursor += 10
		if m.cursor >= len(m.procs) {
			m.cursor = len(m.procs) - 1
		}
	case "c":
		m.sortBy = sysinfo.SortByCPU
		m.statusMsg = "Sorted by CPU"
		return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
	case "m":
		m.sortBy = sysinfo.SortByMem
		m.statusMsg = "Sorted by Memory"
		return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
	case "p":
		m.sortBy = sysinfo.SortByPID
		m.statusMsg = "Sorted by PID"
		return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
	case "n":
		m.sortBy = sysinfo.SortByName
		m.statusMsg = "Sorted by Name"
		return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
	case "k", "K":
		if m.cursor >= 0 && m.cursor < len(m.procs) {
			m.confirmKill = true
		}
	case "r":
		m.statusMsg = "Refreshing…"
		return m, fetchDataCmd(m.cfg.ProcessLimit, m.sortBy)
	case "?":
		m.showHelp = true
	}
	return m, nil
}
