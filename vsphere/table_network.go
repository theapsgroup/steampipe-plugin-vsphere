package vsphere

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type Network struct {
	Name       string
	Type       string
	IPPoolName string
	IPPoolId   int32
	Accessible bool
}

func tableNetwork() *plugin.Table {
	return &plugin.Table{
		Name:        "vsphere_network",
		Description: "VM's running in vsphere",
		List: &plugin.ListConfig{
			Hydrate: listNetworks,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the network"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the network"},
			{Name: "ip_pool_name", Type: proto.ColumnType_STRING, Description: "Name of the associated IP pool. Empty if the network is not associated with an IP pool"},
			{Name: "ip_pool_id", Type: proto.ColumnType_INT, Description: "Identifier of the associated IP pool. Zero if the network is not associated with an IP pool"},
			{Name: "accessible", Type: proto.ColumnType_BOOL, Description: "At least one host is configured to provide this network"},
		},
	}
}

func listNetworks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error connecting to vsphere: %v", err))
	}

	manager := view.NewManager(client)

	var networks []mo.Network
	// https://code.vmware.com/apis/704/vsphere/vim.Network.html
	networkView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"Network"}, true)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating network view: %v", err))
	}
	err = networkView.Retrieve(ctx, []string{"Network"}, []string{"summary"}, &networks)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error listing network summary: %v", err))
	}

	for _, n := range networks {
		summary := n.Summary.GetNetworkSummary()

		var ippoolid int32 = -1
		if summary.IpPoolId != nil {
			ippoolid = *summary.IpPoolId
		}

		d.StreamListItem(ctx, Network{
			Name:       summary.Name,
			Type:       summary.Network.Type,
			IPPoolName: summary.IpPoolName,
			IPPoolId:   ippoolid,
			Accessible: summary.Accessible,
		})
	}
	return nil, nil
}
