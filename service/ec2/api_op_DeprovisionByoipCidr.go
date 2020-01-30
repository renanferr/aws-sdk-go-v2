// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeprovisionByoipCidrInput struct {
	_ struct{} `type:"structure"`

	// The address range, in CIDR notation. The prefix must be the same prefix that
	// you specified when you provisioned the address range.
	//
	// Cidr is a required field
	Cidr *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DeprovisionByoipCidrInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeprovisionByoipCidrInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeprovisionByoipCidrInput"}

	if s.Cidr == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cidr"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeprovisionByoipCidrOutput struct {
	_ struct{} `type:"structure"`

	// Information about the address range.
	ByoipCidr *ByoipCidr `locationName:"byoipCidr" type:"structure"`
}

// String returns the string representation
func (s DeprovisionByoipCidrOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeprovisionByoipCidr = "DeprovisionByoipCidr"

// DeprovisionByoipCidrRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Releases the specified address range that you provisioned for use with your
// AWS resources through bring your own IP addresses (BYOIP) and deletes the
// corresponding address pool.
//
// Before you can release an address range, you must stop advertising it using
// WithdrawByoipCidr and you must not have any IP addresses allocated from its
// address range.
//
//    // Example sending a request using DeprovisionByoipCidrRequest.
//    req := client.DeprovisionByoipCidrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeprovisionByoipCidr
func (c *Client) DeprovisionByoipCidrRequest(input *DeprovisionByoipCidrInput) DeprovisionByoipCidrRequest {
	op := &aws.Operation{
		Name:       opDeprovisionByoipCidr,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeprovisionByoipCidrInput{}
	}

	req := c.newRequest(op, input, &DeprovisionByoipCidrOutput{})
	return DeprovisionByoipCidrRequest{Request: req, Input: input, Copy: c.DeprovisionByoipCidrRequest}
}

// DeprovisionByoipCidrRequest is the request type for the
// DeprovisionByoipCidr API operation.
type DeprovisionByoipCidrRequest struct {
	*aws.Request
	Input *DeprovisionByoipCidrInput
	Copy  func(*DeprovisionByoipCidrInput) DeprovisionByoipCidrRequest
}

// Send marshals and sends the DeprovisionByoipCidr API request.
func (r DeprovisionByoipCidrRequest) Send(ctx context.Context) (*DeprovisionByoipCidrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeprovisionByoipCidrResponse{
		DeprovisionByoipCidrOutput: r.Request.Data.(*DeprovisionByoipCidrOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeprovisionByoipCidrResponse is the response type for the
// DeprovisionByoipCidr API operation.
type DeprovisionByoipCidrResponse struct {
	*DeprovisionByoipCidrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeprovisionByoipCidr request.
func (r *DeprovisionByoipCidrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
