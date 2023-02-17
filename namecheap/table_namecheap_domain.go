package namecheap

import (
	"context"

	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableNamecheapDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "namecheap_domain",
		Description: "Domains registered in Namecheap.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "domain", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "Name of the domain, e.g. steampipe.io."},
			{Name: "expires", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Expires").Transform(dateTimeToTimestamp), Description: "Time when the domain expires."},
			// Other columns
			{Name: "auto_renew", Type: proto.ColumnType_BOOL, Description: "True if the domain will renew automatically."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(dateTimeToTimestamp), Description: "Time when the domain was created at Namecheap."},
			{Name: "dns_provider_type", Type: proto.ColumnType_STRING, Hydrate: getDomainInfo, Transform: transform.FromField("DnsDetails.ProviderType"), Description: ""},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the domain, e.g. 1234."},
			{Name: "is_expired", Type: proto.ColumnType_BOOL, Description: "True if the domain has expired."},
			{Name: "is_locked", Type: proto.ColumnType_BOOL, Description: "True if the domain is locked."},
			{Name: "is_our_dns", Type: proto.ColumnType_BOOL, Description: "True if the domain uses Namecheap DNS."},
			{Name: "is_premium", Type: proto.ColumnType_BOOL, Description: "True if the domain uses premium DNS."},
			{Name: "nameservers", Type: proto.ColumnType_JSON, Hydrate: getDomainInfo, Transform: transform.FromField("DnsDetails.Nameservers"), Description: "Array of nameservers for the domain."},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "Namecheap user who owns the domain, e.g. janedoe."},
			{Name: "whois_guard", Type: proto.ColumnType_STRING, Description: "Status of WhoisGuard, e.g. ENABLED, NOTPRESENT, etc."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("namecheap_domain.listDomain", "connection_error", err)
		return nil, err
	}

	// Set up paging, using the max page size possible.
	page := 0
	pageSize := 100

	// Keep a count, so we know when to exit the loop.
	// Is there a better way from the API response? I couldn't find one.
	count := 0

	params := &namecheap.DomainsGetListArgs{
		Page:     &page,
		PageSize: &pageSize,
	}

	// If given a domain qualifier, then limit our search to match it for
	// efficiency. It's not exact, but it will help.
	if d.EqualsQuals["domain"] != nil {
		s := d.EqualsQuals["domain"].GetStringValue()
		params.SearchTerm = &s
	}

	for {
		page++
		resp, err := conn.Domains.GetList(params)
		plugin.Logger(ctx).Debug("namecheap_domain.listDomain", "params.Page", *params.Page, "params.SearchTeam", params.SearchTerm)
		if err != nil {
			plugin.Logger(ctx).Error("namecheap_domain.listDomain", "query_error", err)
			return nil, err
		}
		plugin.Logger(ctx).Debug("namecheap_domain.listDomain", "paging.TotalItems", *resp.Paging.TotalItems)
		for _, i := range *resp.Domains {
			count++
			d.StreamListItem(ctx, i)
		}
		if count >= *resp.Paging.TotalItems {
			break
		}
	}

	return nil, nil
}

func getDomainInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// Get domain from hydrate data
	domain := h.Item.(namecheap.Domain)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("namecheap_domain.getDomainInfo", "connection_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("namecheap_domain.getDomainInfo", "domain.Name", *domain.Name)

	resp, err := conn.Domains.GetInfo(*domain.Name)
	if err != nil {
		plugin.Logger(ctx).Error("namecheap_domain.getDomainInfo", "query_error", err)
		return nil, err
	}

	return resp.DomainDNSGetListResult, nil
}
