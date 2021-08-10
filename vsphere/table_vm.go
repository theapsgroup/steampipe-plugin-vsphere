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
	ID     string
	Name   string
	Memory int32
	NumCPU int32
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

	for _, vm := range vms {
		logger.Warn(vm.Summary.Config.InstanceUuid)
		d.StreamListItem(ctx, VM{
			ID:     vm.Summary.Config.GuestId,
			Name:   vm.Summary.Config.Name,
			Memory: vm.Summary.Config.MemorySizeMB,
			NumCPU: vm.Summary.Config.NumCpu,
		})
	}
	return nil, nil
}
