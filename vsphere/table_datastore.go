package vsphere

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

type DataStore struct {
	Name        string
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
			{Name: "capacity", Type: proto.ColumnType_INT, Description: "Capacity in bytes"},
			{Name: "uncommitted", Type: proto.ColumnType_INT, Description: "How much storage on this datastore has been allocated to guests in bytes"},
			{Name: "free", Type: proto.ColumnType_INT, Description: "Free space left in bytes"},
			{Name: "accessible", Type: proto.ColumnType_BOOL, Description: "Whether this datastore is accessible"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of Datastore"},
		},
	}
}

func listDatastores(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, _ := connect(ctx, d)
	manager := view.NewManager(client)

	var dss []mo.Datastore
	datastoreView, err := manager.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"Datastore"}, true)
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}
	err = datastoreView.Retrieve(ctx, []string{"Datastore"}, []string{"summary"}, &dss)
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}

	for _, ds := range dss {
		d.StreamListItem(ctx, DataStore{
			Name:        ds.Summary.Name,
			Capacity:    ds.Summary.Capacity,
			Uncommitted: ds.Summary.Uncommitted,
			Free:        ds.Summary.FreeSpace,
			Accessible:  ds.Summary.Accessible,
			Type:        ds.Summary.Type,
		})
	}
	return nil, nil
}
