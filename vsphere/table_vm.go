package vsphere

import (
	"context"
	"fmt"
    "encoding/json"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	//"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type VM struct {
	ID               string
	Moref            string
	Name             string
	GuestFullName    string
	Memory           int32
	NumCPU           int32
	Hardware         string
	IPAddress        string
	Uptime           int32
	Power            string
	Status           string
	CPUUsage         int32
	GuestMemoryUsage int32
	HostMemoryUsage  int32
	Hostmoref        string
	Storage          string
}

func tableVm() *plugin.Table {
	return &plugin.Table{
		Name:        "vsphere_vm",
		Description: "VM's running in vsphere",
		List: &plugin.ListConfig{
			Hydrate: listVms,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The guest operating system identifier (short name)"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the virtual machine"},
			{Name: "moref", Type: proto.ColumnType_STRING, Description: "Managed object reference of the virtual machine"},
			{Name: "guest_full_name", Type: proto.ColumnType_STRING, Description: "Guest operating system name configured on the virtual machine."},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "Memory size of the virtual machine in MB"},
			{Name: "num_cpu", Type: proto.ColumnType_INT, Description: "Number of virtual processors in the virtual machine"},
			{Name: "hardware", Type: proto.ColumnType_STRING, Description: "Version of the virtual hardware"},
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "Primary IP address assigned to the guest operating system, if known"},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "The guest uptime in seconds"},
			{Name: "power", Type: proto.ColumnType_STRING, Description: "The powerstate of this vm"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The overall guest status"},
			{Name: "cpu_usage", Type: proto.ColumnType_INT, Description: "VM cpu usage in mhz"},
			{Name: "guest_memory_usage", Type: proto.ColumnType_INT, Description: "Current memory usage in mb"},
			{Name: "host_memory_usage", Type: proto.ColumnType_INT, Description: "Consumed memory on the host by this vm"},
			{Name: "hostmoref", Type: proto.ColumnType_STRING, Description: "The host that is responsible for running a virtual machine."},
			{Name: "storage", Type: proto.ColumnType_JSON, Description: "Storage Usage"},
			//{Name: "raw", Type: proto.ColumnType_JSON, Description: "Raw data.", Transform: transform.FromValue()},
		},
	}
}

func listVms(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error connecting to vsphere: %v", err))
	}

	manager := view.NewManager(client)

	var vms []mo.VirtualMachine
	// https://code.vmware.com/apis/704/vsphere/vim.VirtualMachine.html
	vmView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating vm view: %v", err))
	}
	err = vmView.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary", "runtime", "config", "storage"}, &vms)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error listing vm summary: %v", err))
	}
	for _, vm := range vms {

        jsonBytes, _ := json.Marshal(vm.Storage.PerDatastoreUsage)

		d.StreamListItem(ctx, VM{
			ID:               vm.Summary.Config.GuestId,
			Name:             vm.Summary.Config.Name,
			Moref:            vm.Summary.Vm.Value,
			GuestFullName:    vm.Summary.Guest.GuestFullName,
			Memory:           vm.Summary.Config.MemorySizeMB,
			NumCPU:           vm.Summary.Config.NumCpu,
			Hardware:         vm.Summary.Config.HwVersion,
			IPAddress:        vm.Summary.Guest.IpAddress,
			Uptime:           vm.Summary.QuickStats.UptimeSeconds,
			Power:            string(vm.Runtime.PowerState),
			Status:           string(vm.Summary.OverallStatus),
			CPUUsage:         vm.Summary.QuickStats.OverallCpuUsage,
			GuestMemoryUsage: vm.Summary.QuickStats.GuestMemoryUsage,
			HostMemoryUsage:  vm.Summary.QuickStats.HostMemoryUsage,
			Hostmoref:        vm.Runtime.Host.Value,
			Storage:          string(jsonBytes),
		})

	}
	return nil, nil
}
