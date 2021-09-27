package vsphere

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type Host struct {
	Name        string
	Vendor      string
	Model       string
	CPU         string
	CPUCores    int16
	CPUThreads  int16
	CPUMhz      int32
	NumNics     int32
	NumHbas     int32
	Memory      int64
	Status      string
	CPUUsage    int32
	MemoryUsage int32
	Uptime      int32
}

func tableHost() *plugin.Table {
	return &plugin.Table{
		Name:        "vsphere_host",
		Description: "Vsphere hosts",
		List: &plugin.ListConfig{
			Hydrate: listHosts,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the host"},
			{Name: "vendor", Type: proto.ColumnType_STRING, Description: "The hardware vendor identification"},
			{Name: "model", Type: proto.ColumnType_STRING, Description: "The system model identification"},
			{Name: "cpu", Type: proto.ColumnType_STRING, Description: "The CPU model"},
			{Name: "cpu_cores", Type: proto.ColumnType_INT, Description: "Number of physical CPU cores on the host"},
			{Name: "cpu_threads", Type: proto.ColumnType_INT, Description: "Number of physical CPU threads on the host"},
			{Name: "cpu_mhz", Type: proto.ColumnType_INT, Description: "The speed of the CPU cores. This is an average value if there are multiple speeds"},
			{Name: "num_nics", Type: proto.ColumnType_INT, Description: "The number of network adapters"},
			{Name: "num_hbas", Type: proto.ColumnType_INT, Description: "The number of host bus adapters"},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "The physical memory size in bytes"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the host"},
			{Name: "cpu_usage", Type: proto.ColumnType_INT, Description: "Current cpu usage in mhz"},
			{Name: "memory_usage", Type: proto.ColumnType_INT, Description: "Current memory usage in MB"},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "The uptime in seconds"},
		},
	}
}

func listHosts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, _ := connect(ctx, d)
	manager := view.NewManager(client)

	var hosts []mo.HostSystem

	//https://code.vmware.com/apis/704/vsphere/vim.HostSystem.html
	hostView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating host container view: %v", err))
	}
	err = hostView.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hosts)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error listing host summary view: %v", err))
	}

	for _, h := range hosts {
		d.StreamListItem(ctx, Host{
			Name:        h.Summary.Config.Name,
			Vendor:      h.Summary.Hardware.Vendor,
			Model:       h.Summary.Hardware.Model,
			CPU:         h.Summary.Hardware.CpuModel,
			CPUCores:    h.Summary.Hardware.NumCpuCores,
			CPUThreads:  h.Summary.Hardware.NumCpuThreads,
			CPUMhz:      h.Summary.Hardware.CpuMhz,
			NumNics:     h.Summary.Hardware.NumNics,
			NumHbas:     h.Summary.Hardware.NumHBAs,
			Memory:      h.Summary.Hardware.MemorySize,
			Status:      string(h.Summary.OverallStatus),
			CPUUsage:    h.Summary.QuickStats.OverallCpuUsage,
			MemoryUsage: h.Summary.QuickStats.OverallMemoryUsage,
			Uptime:      h.Summary.QuickStats.Uptime,
		})
	}
	return nil, nil
}
