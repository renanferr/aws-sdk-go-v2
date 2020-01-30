// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkforceInput struct {
	_ struct{} `type:"structure"`

	// The name of the private workforce whose access you want to restrict. WorkforceName
	// is automatically set to "default" when a workforce is created and cannot
	// be modified.
	//
	// WorkforceName is a required field
	WorkforceName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeWorkforceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkforceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkforceInput"}

	if s.WorkforceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkforceName"))
	}
	if s.WorkforceName != nil && len(*s.WorkforceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkforceName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkforceOutput struct {
	_ struct{} `type:"structure"`

	// A single private workforce, which is automatically created when you create
	// your first private work team. You can create one private work force in each
	// AWS Region. By default, any workforce related API operation used in a specific
	// region will apply to the workforce created in that region. To learn how to
	// create a private workforce, see Create a Private Workforce (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html).
	//
	// Workforce is a required field
	Workforce *Workforce `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeWorkforceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkforce = "DescribeWorkforce"

// DescribeWorkforceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists private workforce information, including workforce name, Amazon Resource
// Name (ARN), and, if applicable, allowed IP address ranges (CIDRs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)).
// Allowable IP address ranges are the IP addresses that workers can use to
// access tasks.
//
// This operation applies only to private workforces.
//
//    // Example sending a request using DescribeWorkforceRequest.
//    req := client.DescribeWorkforceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeWorkforce
func (c *Client) DescribeWorkforceRequest(input *DescribeWorkforceInput) DescribeWorkforceRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkforce,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkforceInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkforceOutput{})
	return DescribeWorkforceRequest{Request: req, Input: input, Copy: c.DescribeWorkforceRequest}
}

// DescribeWorkforceRequest is the request type for the
// DescribeWorkforce API operation.
type DescribeWorkforceRequest struct {
	*aws.Request
	Input *DescribeWorkforceInput
	Copy  func(*DescribeWorkforceInput) DescribeWorkforceRequest
}

// Send marshals and sends the DescribeWorkforce API request.
func (r DescribeWorkforceRequest) Send(ctx context.Context) (*DescribeWorkforceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkforceResponse{
		DescribeWorkforceOutput: r.Request.Data.(*DescribeWorkforceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkforceResponse is the response type for the
// DescribeWorkforce API operation.
type DescribeWorkforceResponse struct {
	*DescribeWorkforceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkforce request.
func (r *DescribeWorkforceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
