package namecheap

import (
	"context"

	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNamecheapDomainHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "namecheap_domain_host",
		Description: "Host records for domains in Namecheap.",
		List: &plugin.ListConfig{
			ParentHydrate: listDomain,
			Hydrate:       listDomainHost,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain", Require: plugin.Optional},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Name of the domain, e.g. steampipe.io."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the host, e.g. www or @ for the root domain."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the host, e.g. CNAME."},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "Target address for the host, e.g. steampipe.io., 1.2.3.4"},
			{Name: "ttl", Type: proto.ColumnType_INT, Description: "TTL of the record."},
			{Name: "mx_pref", Type: proto.ColumnType_INT, Transform: transform.FromField("MXPref"), Description: "MX priority of the host."},
			// Other columns
			// Always null? - {Name: "associated_app_title", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "email_type", Type: proto.ColumnType_STRING, Description: "Email type for the domain."},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "Friendly name of the host, e.g. MX2."},
			{Name: "host_id", Type: proto.ColumnType_INT, Transform: transform.FromField("HostId"), Description: "Unique ID of the host record at Namecheap, e.g. 1234."},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "True if the record is active."},
			{Name: "is_ddns_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IsDDNSEnabled"), Description: "True if dynamic DNS is enabled."},
			{Name: "is_using_our_dns", Type: proto.ColumnType_BOOL, Description: "True if the domain is using Namecheap DNS."},
		}),
	}
}

type hostRow struct {
	namecheap.DomainsDNSHostRecordDetailed
	Domain        *string
	EmailType     *string
	IsUsingOurDNS *bool
}

func listDomainHost(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Get domain from hydrate data
	domain := h.Item.(namecheap.Domain)

	plugin.Logger(ctx).Debug("namecheap_domain.listDomainHost", "domain", domain)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("namecheap_domain.listDomainHost", "connection_error", err)
		return nil, err
	}

	resp, err := conn.DomainsDNS.GetHosts(*domain.Name)
	if err != nil {
		// Ignore errors, because the domain may not have any hosts
		plugin.Logger(ctx).Error("namecheap_domain.listDomainHost", "query_error", err)
	} else {
		for _, host := range *resp.DomainDNSGetHostsResult.Hosts {
			d.StreamListItem(ctx, hostRow{host, resp.DomainDNSGetHostsResult.Domain, resp.DomainDNSGetHostsResult.EmailType, resp.DomainDNSGetHostsResult.IsUsingOurDNS})
		}
	}

	return nil, nil
}
