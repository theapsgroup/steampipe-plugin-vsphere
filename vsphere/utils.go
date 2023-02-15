package vsphere

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

const DEFAULT_ALLOW_UNVERFIED_SSL = false

func connect(ctx context.Context, d *plugin.QueryData) (*vim25.Client, error) {
	vsphereConfig := GetConfig(d.Connection)

	// Initial values. Env vars will be overridden by configuration if values are set in there
	vsphereServer := os.Getenv("VSPHERE_SERVER")
	user := os.Getenv("VSPHERE_USER")
	password := os.Getenv("VSPHERE_PASSWORD")
	allowUnverifiedSSL := os.Getenv("ALLOW_UNVERIFIED_SSL")

	// Override potential env values with config values
	if vsphereConfig.AllowUnverifiedSSL != nil {
		allowUnverifiedSSL = *vsphereConfig.AllowUnverifiedSSL
	}

	if vsphereConfig.VsphereServer != nil {
		vsphereServer = *vsphereConfig.VsphereServer
	}

	if vsphereConfig.User != nil {
		user = *vsphereConfig.User
	}

	if vsphereConfig.Password != nil {
		password = *vsphereConfig.Password
	}

	// Make sure we have all required arguments set via either env or config
	if user == "" || password == "" || vsphereServer == "" {
		errorMsg := ""
		if user == "" {
			errorMsg += "Missing user from config or env 'VSPHERE_USER'\n"
		}
		if password == "" {
			errorMsg += "Missing password from config or env 'VSPHERE_PASSWORD'\n"
		}
		if vsphereServer == "" {
			errorMsg += "Missing vsphere_server from config or env 'VSPHERE_SERVER'\n"
		}
		return nil, fmt.Errorf(errorMsg)
	}

	// Unverified ssl is a bit odd because it is optional
	// If it's unset in the config, and set in the environment, we want to make sure we parse it and error if we can't, instead of silently defaulting
	unverifiedSSLEnv := os.Getenv("VSPHERE_ALLOW_UNVERIFIED_SSL")
	if unverifiedSSLEnv != "" && vsphereConfig.AllowUnverifiedSSL == nil {
		parsed, err := strconv.ParseBool(unverifiedSSLEnv)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse VSPHERE_ALLOW_UNVERIFIED_SSL: Value: %s, Error: %v", unverifiedSSLEnv, err)
		}
		allowUnverifiedSSL = parsed
	}

	client := new(vim25.Client)

	parsedUrl, err := soap.ParseURL(vsphereServer)
	if err != nil {
		return nil, fmt.Errorf("Error parsing vsphere url: %v", err)
	}
	parsedUrl.User = url.UserPassword(user, password)

	session := &cache.Session{
		URL:      parsedUrl,
		Insecure: allowUnverifiedSSL,
	}
	err = session.Login(ctx, client, nil)
	if err != nil {
		return nil, fmt.Errorf("Error authenticating with vsphere: %v", err)
	}

	return client, nil
}
