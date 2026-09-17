package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskInfo struct {
	Mountpoint  string
	Device      string
	Fstype      string
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
}

func GetDisks() ([]DiskInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, fmt.Errorf("partitions: %w", err)
	}

	var result []DiskInfo
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		if usage.Total < 1<<20 {
			continue
		}
		result = append(result, DiskInfo{
			Mountpoint:  p.Mountpoint,
			Device:      p.Device,
			Fstype:      p.Fstype,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}
	return result, nil
}
