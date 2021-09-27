package vsphere

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type VM struct {
	ID               string
	Name             string
	Memory           int32
	NumCPU           int32
	IPAddress        string
	Uptime           int32
	Status           string
	CPUUsage         int32
	GuestMemoryUsage int32
	HostMemoryUsage  int32
}

func tableVm() *plugin.Table {
	return &plugin.Table{
		Name:        "vsphere_vm",
		Description: "VM's running in vsphere",
		List: &plugin.ListConfig{
			Hydrate: listVms,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The guest id"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the guest"},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "The the amount of memory"},
			{Name: "num_cpu", Type: proto.ColumnType_INT, Description: "The cpu core count"},
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "IP Address of the vm"},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "The guest uptime in seconds"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The overall guest status"},
			{Name: "cpu_usage", Type: proto.ColumnType_INT, Description: "VM cpu usage in mhz"},
			{Name: "guest_memory_usage", Type: proto.ColumnType_INT, Description: "Current memory usage in mb"},
			{Name: "host_memory_usage", Type: proto.ColumnType_INT, Description: "Consumed memory on the host by this vm"},
		},
	}
}

func listVms(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, _ := connect(ctx, d)
	manager := view.NewManager(client)

	var vms []mo.VirtualMachine
	//https://code.vmware.com/apis/704/vsphere/vim.VirtualMachine.html
	vmView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating vm view: %v", err))
	}
	err = vmView.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error listing vm summary: %v", err))
	}
	for _, vm := range vms {
		d.StreamListItem(ctx, VM{
			ID:               vm.Summary.Config.GuestId,
			Name:             vm.Summary.Config.Name,
			Memory:           vm.Summary.Config.MemorySizeMB,
			NumCPU:           vm.Summary.Config.NumCpu,
			IPAddress:        vm.Summary.Guest.IpAddress,
			Uptime:           vm.Summary.QuickStats.UptimeSeconds,
			Status:           string(vm.Summary.OverallStatus),
			CPUUsage:         vm.Summary.QuickStats.OverallCpuUsage,
			GuestMemoryUsage: vm.Summary.QuickStats.GuestMemoryUsage,
			HostMemoryUsage:  vm.Summary.QuickStats.HostMemoryUsage,
		})
	}
	return nil, nil
}
