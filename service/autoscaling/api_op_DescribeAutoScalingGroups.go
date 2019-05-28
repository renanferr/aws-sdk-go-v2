// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AutoScalingGroupNamesType
type DescribeAutoScalingGroupsInput struct {
	_ struct{} `type:"structure"`

	// The names of the Auto Scaling groups. Each name can be a maximum of 1600
	// characters. By default, you can only specify up to 50 names. You can optionally
	// increase this limit using the MaxRecords parameter.
	//
	// If you omit this parameter, all Auto Scaling groups are described.
	AutoScalingGroupNames []string `type:"list"`

	// The maximum number of items to return with this call. The default value is
	// 50 and the maximum value is 100.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AutoScalingGroupsType
type DescribeAutoScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The groups.
	//
	// AutoScalingGroups is a required field
	AutoScalingGroups []AutoScalingGroup `type:"list" required:"true"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAutoScalingGroups = "DescribeAutoScalingGroups"

// DescribeAutoScalingGroupsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes one or more Auto Scaling groups.
//
//    // Example sending a request using DescribeAutoScalingGroupsRequest.
//    req := client.DescribeAutoScalingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeAutoScalingGroups
func (c *Client) DescribeAutoScalingGroupsRequest(input *DescribeAutoScalingGroupsInput) DescribeAutoScalingGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeAutoScalingGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeAutoScalingGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeAutoScalingGroupsOutput{})
	return DescribeAutoScalingGroupsRequest{Request: req, Input: input, Copy: c.DescribeAutoScalingGroupsRequest}
}

// DescribeAutoScalingGroupsRequest is the request type for the
// DescribeAutoScalingGroups API operation.
type DescribeAutoScalingGroupsRequest struct {
	*aws.Request
	Input *DescribeAutoScalingGroupsInput
	Copy  func(*DescribeAutoScalingGroupsInput) DescribeAutoScalingGroupsRequest
}

// Send marshals and sends the DescribeAutoScalingGroups API request.
func (r DescribeAutoScalingGroupsRequest) Send(ctx context.Context) (*DescribeAutoScalingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAutoScalingGroupsResponse{
		DescribeAutoScalingGroupsOutput: r.Request.Data.(*DescribeAutoScalingGroupsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeAutoScalingGroupsRequestPaginator returns a paginator for DescribeAutoScalingGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeAutoScalingGroupsRequest(input)
//   p := autoscaling.NewDescribeAutoScalingGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeAutoScalingGroupsPaginator(req DescribeAutoScalingGroupsRequest) DescribeAutoScalingGroupsPaginator {
	return DescribeAutoScalingGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeAutoScalingGroupsInput
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

// DescribeAutoScalingGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeAutoScalingGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeAutoScalingGroupsPaginator) CurrentPage() *DescribeAutoScalingGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeAutoScalingGroupsOutput)
}

// DescribeAutoScalingGroupsResponse is the response type for the
// DescribeAutoScalingGroups API operation.
type DescribeAutoScalingGroupsResponse struct {
	*DescribeAutoScalingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAutoScalingGroups request.
func (r *DescribeAutoScalingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}