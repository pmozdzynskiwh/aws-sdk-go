// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourceexplorer2 provides the client and types for making API
// requests to AWS Resource Explorer.
//
// Amazon Web Services Resource Explorer is a resource search and discovery
// service. By using Resource Explorer, you can explore your resources using
// an internet search engine-like experience. Examples of resources include
// Amazon Relational Database Service (Amazon RDS) instances, Amazon Simple
// Storage Service (Amazon S3) buckets, or Amazon DynamoDB tables. You can search
// for your resources using resource metadata like names, tags, and IDs. Resource
// Explorer can search across all of the Amazon Web Services Regions in your
// account in which you turn the service on, to simplify your cross-Region workloads.
//
// Resource Explorer scans the resources in each of the Amazon Web Services
// Regions in your Amazon Web Services account in which you turn on Resource
// Explorer. Resource Explorer creates and maintains an index (https://docs.aws.amazon.com/arexug/mainline/getting-started-terms-and-concepts.html#term-index)
// in each Region, with the details of that Region's resources.
//
// You can search across all of the indexed Regions in your account (https://docs.aws.amazon.com/arexug/mainline/manage-aggregator-region.html)
// by designating one of your Amazon Web Services Regions to contain the aggregator
// index for the account. When you promote a local index in a Region to become
// the aggregator index for the account (https://docs.aws.amazon.com/arexug/mainline/manage-aggregator-region-turn-on.html),
// Resource Explorer automatically replicates the index information from all
// local indexes in the other Regions to the aggregator index. Therefore, the
// Region with the aggregator index has a copy of all resource information for
// all Regions in the account where you turned on Resource Explorer. As a result,
// views in the aggregator index Region include resources from all of the indexed
// Regions in your account.
//
//	<p>For more information about Amazon Web Services Resource Explorer, including
//	how to enable and configure the service, see the <a href="https://docs.aws.amazon.com/resource-explorer/latest/userguide/">Amazon
//	Web Services Resource Explorer User Guide</a>.</p> <note> <p>The example
//	HTTP query requests and responses in this guide are displayed with the
//	<a href="https://json.org">JSON</a> formatted across multiple lines for
//	readability. The actual query responses from the Resource Explorer service
//	do not include this extra whitespace.</p> </note> <p> <b>We want your
//	feedback about this documentation</b> </p> <p>Our goal is to help you
//	get everything you can from Resource Explorer. If this guide helps you
//	to do that, then let us know. If the guide isn't helping you, then we
//	want to hear from you so we can address the issue. Use the <b>Feedback</b>
//	link that's in the upper-right corner of every page. That sends your comments
//	directly to the writers of this guide. We review every submission, looking
//	for opportunities to improve the documentation. Thank you in advance for
//	your help!</p>
//
// See https://docs.aws.amazon.com/goto/WebAPI/resource-explorer-2-2022-07-28 for more information on this service.
//
// See resourceexplorer2 package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourceexplorer2/
//
// # Using the Client
//
// To contact AWS Resource Explorer with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Resource Explorer client ResourceExplorer2 for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/resourceexplorer2/#New
package resourceexplorer2
