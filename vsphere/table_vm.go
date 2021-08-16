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
	ID        string
	Name      string
	Memory    int32
	NumCPU    int32
	IPAddress string
	Uptime    int32
	Status    string
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
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "The cpu core count"},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "The host uptime in seconds"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The overall guest status"},
		},
	}
}

func listVms(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, _ := connect(ctx, d)
	manager := view.NewManager(client)

	var vms []mo.VirtualMachine
	vmView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}
	err = vmView.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}

	//https://code.vmware.com/apis/704/vsphere/vmodl.query.PropertyCollector.PropertySpec.html
	//https://code.vmware.com/apis/704/vsphere/vim.VirtualMachine.html
	var vmstatus []mo.VirtualMachine
	err = vmView.Retrieve(ctx, []string{"VirtualMachine"}, []string{"guestHeartbeatStatus"}, &vmstatus)

	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}

	for _, vm := range vms {
		logger.Warn(vm.Summary.Config.InstanceUuid)

		d.StreamListItem(ctx, VM{
			ID:        vm.Summary.Config.GuestId,
			Name:      vm.Summary.Config.Name,
			Memory:    vm.Summary.Config.MemorySizeMB,
			NumCPU:    vm.Summary.Config.NumCpu,
			IPAddress: vm.Summary.Guest.IpAddress,
			Uptime:    vm.Summary.QuickStats.UptimeSeconds,
			Status:    string(vm.Summary.OverallStatus),
		})
	}
	return nil, nil
}
