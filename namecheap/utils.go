package namecheap

import (
	"context"
	"errors"
	"os"

	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*namecheap.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*namecheap.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	var conn *namecheap.Client

	// Default to the env var settings
	username := os.Getenv("NAMECHEAP_USERNAME")
	apiUser := os.Getenv("NAMECHEAP_API_USER")
	apiKey := os.Getenv("NAMECHEAP_API_KEY")

	// Prefer config settings
	namecheapConfig := GetConfig(d.Connection)
	if namecheapConfig.Username != nil {
		username = *namecheapConfig.Username
	}
	if namecheapConfig.APIUser != nil {
		apiUser = *namecheapConfig.APIUser
	}
	if namecheapConfig.APIKey != nil {
		apiKey = *namecheapConfig.APIKey
	}

	// Error if the minimum config is not set
	if username == "" {
		return conn, errors.New("username must be configured")
	}
	if apiUser == "" {
		// Assume the same as the username by default
		apiUser = username
	}
	if apiKey == "" {
		return conn, errors.New("api_key must be configured")
	}

	conn = namecheap.NewClient(&namecheap.ClientOptions{
		UserName: username,
		ApiUser:  apiUser,
		ApiKey:   apiKey,
		// This is required by the SDK, but the server seems to only look at the actual
		// client IP that made the request. So, provide a dummy value for convenience
		// to reduce configuration complexity.
		ClientIp:   "1.2.3.4",
		UseSandbox: false,
	})

	return conn, nil
}

func dateTimeToTimestamp(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	dt := d.Value.(*namecheap.DateTime)
	return dt.Time, nil
}
