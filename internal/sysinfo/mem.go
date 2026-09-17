package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemInfo struct {
	Total       uint64
	Used        uint64
	Available   uint64
	UsedPercent float64
	SwapTotal   uint64
	SwapUsed    uint64
	SwapPercent float64
}

func GetMem() (MemInfo, error) {
	info := MemInfo{}

	v, err := mem.VirtualMemory()
	if err != nil {
		return info, fmt.Errorf("virtual memory: %w", err)
	}

	info.Total = v.Total
	info.Used = v.Used
	info.Available = v.Available
	info.UsedPercent = v.UsedPercent

	s, err := mem.SwapMemory()
	if err != nil {
		return info, fmt.Errorf("swap memory: %w", err)
	}

	info.SwapTotal = s.Total
	info.SwapUsed = s.Used
	info.SwapPercent = s.UsedPercent

	return info, nil
}

func FormatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
