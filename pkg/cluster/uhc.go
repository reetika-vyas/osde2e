package cluster

import (
	"fmt"

	uhc "github.com/openshift-online/uhc-sdk-go/pkg/client"
)

const (
	StagingURL = "https://api.stage.openshift.com"
	APIPrefix  = "/api/clusters_mgmt"
	APIVersion = "v1"
)

// NewUHC setups a client to connect to UHC.
func NewUHC(token string, staging, debug bool) (*UHC, error) {
	logger, err := uhc.NewGoLoggerBuilder().
		Debug(debug).
		Build()
	if err != nil {
		return nil, fmt.Errorf("couldn't build logger: %v", err)
	}

	builder := uhc.NewConnectionBuilder().
		Logger(logger).
		Tokens(token)

	if staging {
		builder.URL(StagingURL)
	}

	conn, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("couldn't setup connection: %v", err)
	}

	return &UHC{
		conn: conn,
	}, nil
}

// UHC acts as a client to manage an instance.
type UHC struct {
	conn *uhc.Connection
}
