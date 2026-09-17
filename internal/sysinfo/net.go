package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetInfo struct {
	BytesSent   uint64
	BytesRecv   uint64
	PacketsSent uint64
	PacketsRecv uint64
}

func GetNet() (NetInfo, error) {
	counters, err := net.IOCounters(false)
	if err != nil {
		return NetInfo{}, fmt.Errorf("net io counters: %w", err)
	}
	if len(counters) == 0 {
		return NetInfo{}, nil
	}
	c := counters[0]
	return NetInfo{
		BytesSent:   c.BytesSent,
		BytesRecv:   c.BytesRecv,
		PacketsSent: c.PacketsSent,
		PacketsRecv: c.PacketsRecv,
	}, nil
}
