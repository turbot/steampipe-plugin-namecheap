package namecheap

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "username",
			Description: "The name of the user.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getUsername,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getUsernameMemoized = plugin.HydrateFunc(getUsernameUncached).Memoize(memoize.WithCacheKeyFunction(getUsernameCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getUsername(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getUsernameMemoized(ctx, d, h)
}

// Build a cache key for the call to getUsernameCacheKey.
func getUsernameCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getUsername"
	return key, nil
}

func getUsernameUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connect(ctx, d,)
	if err != nil {
		return nil, err
	}

	return client.ClientOptions.UserName, nil
}
