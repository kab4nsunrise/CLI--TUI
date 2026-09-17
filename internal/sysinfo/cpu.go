package sysinfo

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUInfo struct {
	UsagePercent float64
	PerCore      []float64
	Cores        int
	ModelName    string
}

func GetCPU() (CPUInfo, error) {
	info := CPUInfo{}

	percents, err := cpu.Percent(0, false)
	if err != nil {
		return info, fmt.Errorf("cpu percent: %w", err)
	}
	if len(percents) > 0 {
		info.UsagePercent = percents[0]
	}

	perCore, err := cpu.Percent(0, true)
	if err != nil {
		return info, fmt.Errorf("cpu per-core: %w", err)
	}
	info.PerCore = perCore

	cores, err := cpu.Counts(true)
	if err != nil {
		return info, fmt.Errorf("cpu counts: %w", err)
	}
	info.Cores = cores

	cpuInfo, err := cpu.Info()
	if err == nil && len(cpuInfo) > 0 {
		info.ModelName = cpuInfo[0].ModelName
	}

	return info, nil
}

func GetCPUBlocking(interval time.Duration) (CPUInfo, error) {
	info := CPUInfo{}

	percents, err := cpu.Percent(interval, false)
	if err != nil {
		return info, err
	}
	if len(percents) > 0 {
		info.UsagePercent = percents[0]
	}

	perCore, err := cpu.Percent(0, true)
	if err != nil {
		return info, err
	}
	info.PerCore = perCore

	cores, _ := cpu.Counts(true)
	info.Cores = cores

	cpuInfo, err := cpu.Info()
	if err == nil && len(cpuInfo) > 0 {
		info.ModelName = cpuInfo[0].ModelName
	}

	return info, nil
}
