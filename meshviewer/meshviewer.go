package meshviewer

import (
	"github.com/FreifunkBremen/respond-collector/data"
	"github.com/FreifunkBremen/respond-collector/jsontime"
)

// Node struct
type Node struct {
	Firstseen  jsontime.Time    `json:"firstseen"`
	Lastseen   jsontime.Time    `json:"lastseen"`
	Flags      Flags            `json:"flags"`
	Statistics *Statistics      `json:"statistics"`
	Nodeinfo   *data.NodeInfo   `json:"nodeinfo"`
	Neighbours *data.Neighbours `json:"-"`
}

type Flags struct {
	Online  bool `json:"online"`
	Gateway bool `json:"gateway"`
}

// Nodes struct: cache DB of Node's structs
type NodesV1 struct {
	Version   int              `json:"version"`
	Timestamp jsontime.Time    `json:"timestamp"`
	List      map[string]*Node `json:"nodes"` // the current nodemap, indexed by node ID
}

type NodesV2 struct {
	Version   int           `json:"version"`
	Timestamp jsontime.Time `json:"timestamp"`
	List      []*Node       `json:"nodes"` // the current nodemap, as array
}

type Statistics struct {
	NodeId      string  `json:"node_id"`
	Clients     uint32  `json:"clients"`
	RootFsUsage float64 `json:"rootfs_usage,omitempty"`
	LoadAverage float64 `json:"loadavg,omitempty"`
	MemoryUsage float64 `json:"memory_usage,omitempty"`
	Uptime      float64 `json:"uptime,omitempty"`
	Idletime    float64 `json:"idletime,omitempty"`
	Gateway     string  `json:"gateway,omitempty"`
	Processes   struct {
		Total   uint32 `json:"total"`
		Running uint32 `json:"running"`
	} `json:"processes,omitempty"`
	MeshVpn *data.MeshVPN `json:"mesh_vpn,omitempty"`
	Traffic struct {
		Tx      *data.Traffic `json:"tx"`
		Rx      *data.Traffic `json:"rx"`
		Forward *data.Traffic `json:"forward"`
		MgmtTx  *data.Traffic `json:"mgmt_tx"`
		MgmtRx  *data.Traffic `json:"mgmt_rx"`
	} `json:"traffic,omitempty"`
}

func NewStatistics(stats *data.Statistics) *Statistics {
	total := stats.Clients.Total
	if total == 0 {
		total = stats.Clients.Wifi24 + stats.Clients.Wifi5
	}
	/* The Meshviewer could not handle absolute memory output
         * calc the used memory as a float witch 100% equal 1.0
         */
	memoryUsage := (float64(stats.Memory.Total) - float64(stats.Memory.Free)) / float64(stats.Memory.Total)

	return &Statistics{
		NodeId:      stats.NodeId,
		Gateway:     stats.Gateway,
		RootFsUsage: stats.RootFsUsage,
		LoadAverage: stats.LoadAverage,
		MemoryUsage: memoryUsage,
		Uptime:      stats.Uptime,
		Idletime:    stats.Idletime,
		Processes:   stats.Processes,
		MeshVpn:     stats.MeshVpn,
		Traffic:     stats.Traffic,
		Clients:     total,
	}
}
