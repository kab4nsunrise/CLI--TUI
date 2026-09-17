package sysinfo

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	PID        int32
	Name       string
	CPUPercent float64
	MemPercent float32
	RSS        uint64
	Status     string
	Username   string
	Cmdline    string
}

type SortBy int

const (
	SortByCPU SortBy = iota
	SortByMem
	SortByPID
	SortByName
)

func GetProcesses(limit int, sortBy SortBy) ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("list processes: %w", err)
	}

	result := make([]ProcessInfo, 0, len(procs))

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPct, _ := p.CPUPercent()
		memPct, _ := p.MemoryPercent()
		memInfo, _ := p.MemoryInfo()
		status, _ := p.Status()
		username, _ := p.Username()
		cmdline, _ := p.Cmdline()

		var rss uint64
		if memInfo != nil {
			rss = memInfo.RSS
		}

		statusStr := ""
		if len(status) > 0 {
			statusStr = status[0]
		}

		result = append(result, ProcessInfo{
			PID:        p.Pid,
			Name:       name,
			CPUPercent: cpuPct,
			MemPercent: memPct,
			RSS:        rss,
			Status:     statusStr,
			Username:   username,
			Cmdline:    cmdline,
		})
	}

	switch sortBy {
	case SortByCPU:
		sort.Slice(result, func(i, j int) bool {
			return result[i].CPUPercent > result[j].CPUPercent
		})
	case SortByMem:
		sort.Slice(result, func(i, j int) bool {
			return result[i].MemPercent > result[j].MemPercent
		})
	case SortByPID:
		sort.Slice(result, func(i, j int) bool {
			return result[i].PID < result[j].PID
		})
	case SortByName:
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
	}

	if limit > 0 && len(result) > limit {
		result = result[:limit]
	}

	return result, nil
}

func KillProcess(pid int32) error {
	p, err := process.NewProcess(pid)
	if err != nil {
		return fmt.Errorf("process %d not found: %w", pid, err)
	}
	return p.Kill()
}
