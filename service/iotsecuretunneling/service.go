// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsecuretunneling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// IoTSecureTunneling provides the API operation methods for making requests to
// AWS IoT Secure Tunneling. See this package's package overview docs
// for details on the service.
//
// IoTSecureTunneling methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type IoTSecureTunneling struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "IoTSecureTunneling" // Name of service.
	EndpointsID = "api.tunneling.iot"  // ID to lookup a service endpoint with.
	ServiceID   = "IoTSecureTunneling" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the IoTSecureTunneling client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//
//	mySession := session.Must(session.NewSession())
//
//	// Create a IoTSecureTunneling client from just a session.
//	svc := iotsecuretunneling.New(mySession)
//
//	// Create a IoTSecureTunneling client with additional configuration
//	svc := iotsecuretunneling.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *IoTSecureTunneling {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "IoTSecuredTunneling"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *IoTSecureTunneling {
	svc := &IoTSecureTunneling{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2018-10-05",
				ResolvedRegion: resolvedRegion,
				EndpointID:     EndpointsID,
				JSONVersion:    "1.1",
				TargetPrefix:   "IoTSecuredTunneling",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(jsonrpc.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a IoTSecureTunneling operation and runs any
// custom request initialization.
func (c *IoTSecureTunneling) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
