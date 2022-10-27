package vsphere

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type Datastore struct {
	Name        string
	Moref       string
	Capacity    int64
	Free        int64
	Uncommitted int64
	Accessible  bool
	Type        string
}

func tableDatastore() *plugin.Table {
	return &plugin.Table{
		Name:        "vsphere_datastore",
		Description: "Vsphere datastores",
		List: &plugin.ListConfig{
			Hydrate: listDatastores,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the datastore"},
			{Name: "moref", Type: proto.ColumnType_STRING, Description: "Managed object reference of the datastore"},
			{Name: "capacity", Type: proto.ColumnType_INT, Description: "Maximum capacity of this datastore in bytes"},
			{Name: "uncommitted", Type: proto.ColumnType_INT, Description: "Total additional storage space, in bytes, potentially used by all virtual machines on this datastore"},
			{Name: "free", Type: proto.ColumnType_INT, Description: "Available space of this datastore, in bytes"},
			{Name: "accessible", Type: proto.ColumnType_BOOL, Description: "The connectivity status of this datastore"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of file system volume, such as VMFS or NFS"},
		},
	}
}

func listDatastores(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error connecting to vsphere: %v", err))
	}

	manager := view.NewManager(client)

	var dss []mo.Datastore
	datastoreView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"Datastore"}, true)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error creating datastore container view: %v", err))
	}
	err = datastoreView.Retrieve(ctx, []string{"Datastore"}, []string{"summary"}, &dss)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error listing datastore summary: %v", err))
	}

	for _, ds := range dss {
		d.StreamListItem(ctx, Datastore{
			Name:        ds.Summary.Name,
			Moref:       ds.Summary.Datastore.Value,
			Capacity:    ds.Summary.Capacity,
			Uncommitted: ds.Summary.Uncommitted,
			Free:        ds.Summary.FreeSpace,
			Accessible:  ds.Summary.Accessible,
			Type:        ds.Summary.Type,
		})
	}
	return nil, nil
}
