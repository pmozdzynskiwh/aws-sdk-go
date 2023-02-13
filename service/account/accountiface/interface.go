// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package accountiface provides an interface to enable mocking the AWS Account service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package accountiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/account"
)

// AccountAPI provides an interface to enable mocking the
// account.Account service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Account.
//	func myFunc(svc accountiface.AccountAPI) bool {
//	    // Make svc.DeleteAlternateContact request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := account.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAccountClient struct {
//	    accountiface.AccountAPI
//	}
//	func (m *mockAccountClient) DeleteAlternateContact(input *account.DeleteAlternateContactInput) (*account.DeleteAlternateContactOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAccountClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AccountAPI interface {
	DeleteAlternateContact(*account.DeleteAlternateContactInput) (*account.DeleteAlternateContactOutput, error)
	DeleteAlternateContactWithContext(aws.Context, *account.DeleteAlternateContactInput, ...request.Option) (*account.DeleteAlternateContactOutput, error)
	DeleteAlternateContactRequest(*account.DeleteAlternateContactInput) (*request.Request, *account.DeleteAlternateContactOutput)

	DisableRegion(*account.DisableRegionInput) (*account.DisableRegionOutput, error)
	DisableRegionWithContext(aws.Context, *account.DisableRegionInput, ...request.Option) (*account.DisableRegionOutput, error)
	DisableRegionRequest(*account.DisableRegionInput) (*request.Request, *account.DisableRegionOutput)

	EnableRegion(*account.EnableRegionInput) (*account.EnableRegionOutput, error)
	EnableRegionWithContext(aws.Context, *account.EnableRegionInput, ...request.Option) (*account.EnableRegionOutput, error)
	EnableRegionRequest(*account.EnableRegionInput) (*request.Request, *account.EnableRegionOutput)

	GetAlternateContact(*account.GetAlternateContactInput) (*account.GetAlternateContactOutput, error)
	GetAlternateContactWithContext(aws.Context, *account.GetAlternateContactInput, ...request.Option) (*account.GetAlternateContactOutput, error)
	GetAlternateContactRequest(*account.GetAlternateContactInput) (*request.Request, *account.GetAlternateContactOutput)

	GetContactInformation(*account.GetContactInformationInput) (*account.GetContactInformationOutput, error)
	GetContactInformationWithContext(aws.Context, *account.GetContactInformationInput, ...request.Option) (*account.GetContactInformationOutput, error)
	GetContactInformationRequest(*account.GetContactInformationInput) (*request.Request, *account.GetContactInformationOutput)

	GetRegionOptStatus(*account.GetRegionOptStatusInput) (*account.GetRegionOptStatusOutput, error)
	GetRegionOptStatusWithContext(aws.Context, *account.GetRegionOptStatusInput, ...request.Option) (*account.GetRegionOptStatusOutput, error)
	GetRegionOptStatusRequest(*account.GetRegionOptStatusInput) (*request.Request, *account.GetRegionOptStatusOutput)

	ListRegions(*account.ListRegionsInput) (*account.ListRegionsOutput, error)
	ListRegionsWithContext(aws.Context, *account.ListRegionsInput, ...request.Option) (*account.ListRegionsOutput, error)
	ListRegionsRequest(*account.ListRegionsInput) (*request.Request, *account.ListRegionsOutput)

	ListRegionsPages(*account.ListRegionsInput, func(*account.ListRegionsOutput, bool) bool) error
	ListRegionsPagesWithContext(aws.Context, *account.ListRegionsInput, func(*account.ListRegionsOutput, bool) bool, ...request.Option) error

	PutAlternateContact(*account.PutAlternateContactInput) (*account.PutAlternateContactOutput, error)
	PutAlternateContactWithContext(aws.Context, *account.PutAlternateContactInput, ...request.Option) (*account.PutAlternateContactOutput, error)
	PutAlternateContactRequest(*account.PutAlternateContactInput) (*request.Request, *account.PutAlternateContactOutput)

	PutContactInformation(*account.PutContactInformationInput) (*account.PutContactInformationOutput, error)
	PutContactInformationWithContext(aws.Context, *account.PutContactInformationInput, ...request.Option) (*account.PutContactInformationOutput, error)
	PutContactInformationRequest(*account.PutContactInformationInput) (*request.Request, *account.PutContactInformationOutput)
}

var _ AccountAPI = (*account.Account)(nil)
