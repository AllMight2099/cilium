// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your managed prefix lists and any Amazon Web Services-managed prefix
// lists.
//
// To view the entries for your prefix list, use GetManagedPrefixListEntries.
func (c *Client) DescribeManagedPrefixLists(ctx context.Context, params *DescribeManagedPrefixListsInput, optFns ...func(*Options)) (*DescribeManagedPrefixListsOutput, error) {
	if params == nil {
		params = &DescribeManagedPrefixListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeManagedPrefixLists", params, optFns, c.addOperationDescribeManagedPrefixListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeManagedPrefixListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeManagedPrefixListsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//
	//   - owner-id - The ID of the prefix list owner.
	//
	//   - prefix-list-id - The ID of the prefix list.
	//
	//   - prefix-list-name - The name of the prefix list.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more prefix list IDs.
	PrefixListIds []string

	noSmithyDocumentSerde
}

type DescribeManagedPrefixListsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the prefix lists.
	PrefixLists []types.ManagedPrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeManagedPrefixListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeManagedPrefixLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeManagedPrefixLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeManagedPrefixLists"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeManagedPrefixLists(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeManagedPrefixListsAPIClient is a client that implements the
// DescribeManagedPrefixLists operation.
type DescribeManagedPrefixListsAPIClient interface {
	DescribeManagedPrefixLists(context.Context, *DescribeManagedPrefixListsInput, ...func(*Options)) (*DescribeManagedPrefixListsOutput, error)
}

var _ DescribeManagedPrefixListsAPIClient = (*Client)(nil)

// DescribeManagedPrefixListsPaginatorOptions is the paginator options for
// DescribeManagedPrefixLists
type DescribeManagedPrefixListsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeManagedPrefixListsPaginator is a paginator for
// DescribeManagedPrefixLists
type DescribeManagedPrefixListsPaginator struct {
	options   DescribeManagedPrefixListsPaginatorOptions
	client    DescribeManagedPrefixListsAPIClient
	params    *DescribeManagedPrefixListsInput
	nextToken *string
	firstPage bool
}

// NewDescribeManagedPrefixListsPaginator returns a new
// DescribeManagedPrefixListsPaginator
func NewDescribeManagedPrefixListsPaginator(client DescribeManagedPrefixListsAPIClient, params *DescribeManagedPrefixListsInput, optFns ...func(*DescribeManagedPrefixListsPaginatorOptions)) *DescribeManagedPrefixListsPaginator {
	if params == nil {
		params = &DescribeManagedPrefixListsInput{}
	}

	options := DescribeManagedPrefixListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeManagedPrefixListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeManagedPrefixListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeManagedPrefixLists page.
func (p *DescribeManagedPrefixListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeManagedPrefixListsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeManagedPrefixLists(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeManagedPrefixLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeManagedPrefixLists",
	}
}
