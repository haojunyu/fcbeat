// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTagsRequest
type DescribeTagsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * key - The tag key.
	//
	//    * resource-id - The ID of the resource.
	//
	//    * resource-type - The resource type (customer-gateway | dedicated-host
	//    | dhcp-options | elastic-ip | fleet | fpga-image | image | instance |
	//    host-reservation | internet-gateway | launch-template | natgateway | network-acl
	//    | network-interface | reserved-instances | route-table | security-group
	//    | snapshot | spot-instances-request | subnet | volume | vpc | vpc-peering-connection
	//    | vpn-connection | vpn-gateway).
	//
	//    * tag:<key> - The key/value combination of the tag. For example, specify
	//    "tag:Owner" for the filter name and "TeamA" for the filter value to find
	//    resources with the tag "Owner=TeamA".
	//
	//    * value - The tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. This value can
	// be between 5 and 1000. To retrieve the remaining results, make another call
	// with the returned NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTagsResult
type DescribeTagsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The tags.
	Tags []TagDescription `locationName:"tagSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTags = "DescribeTags"

// DescribeTagsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified tags for your EC2 resources.
//
// For more information about tags, see Tagging Your Resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeTagsRequest.
//    req := client.DescribeTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTags
func (c *Client) DescribeTagsRequest(input *DescribeTagsInput) DescribeTagsRequest {
	op := &aws.Operation{
		Name:       opDescribeTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeTagsInput{}
	}

	req := c.newRequest(op, input, &DescribeTagsOutput{})
	return DescribeTagsRequest{Request: req, Input: input, Copy: c.DescribeTagsRequest}
}

// DescribeTagsRequest is the request type for the
// DescribeTags API operation.
type DescribeTagsRequest struct {
	*aws.Request
	Input *DescribeTagsInput
	Copy  func(*DescribeTagsInput) DescribeTagsRequest
}

// Send marshals and sends the DescribeTags API request.
func (r DescribeTagsRequest) Send(ctx context.Context) (*DescribeTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTagsResponse{
		DescribeTagsOutput: r.Request.Data.(*DescribeTagsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTagsRequestPaginator returns a paginator for DescribeTags.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTagsRequest(input)
//   p := ec2.NewDescribeTagsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTagsPaginator(req DescribeTagsRequest) DescribeTagsPaginator {
	return DescribeTagsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTagsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeTagsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTagsPaginator struct {
	aws.Pager
}

func (p *DescribeTagsPaginator) CurrentPage() *DescribeTagsOutput {
	return p.Pager.CurrentPage().(*DescribeTagsOutput)
}

// DescribeTagsResponse is the response type for the
// DescribeTags API operation.
type DescribeTagsResponse struct {
	*DescribeTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTags request.
func (r *DescribeTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
