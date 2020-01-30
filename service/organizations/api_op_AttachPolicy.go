// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type AttachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the policy that you want to attach to the target.
	// You can get the ID for the policy by calling the ListPolicies operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lowercase or uppercase letters, digits,
	// or the underscore character (_).
	//
	// PolicyId is a required field
	PolicyId *string `type:"string" required:"true"`

	// The unique identifier (ID) of the root, OU, or account that you want to attach
	// the policy to. You can get the ID by calling the ListRoots, ListOrganizationalUnitsForParent,
	// or ListAccounts operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a target ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//    letters or digits.
	//
	//    * Account - A string that consists of exactly 12 digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lowercase letters or digits (the ID of the root that the
	//    OU is in). This string is followed by a second "-" dash and from 8 to
	//    32 additional lowercase letters or digits.
	//
	// TargetId is a required field
	TargetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachPolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if s.TargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachPolicy = "AttachPolicy"

// AttachPolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Attaches a policy to a root, an organizational unit (OU), or an individual
// account.
//
// How the policy affects accounts depends on the type of policy:
//
//    * For more information about attaching SCPs, see How SCPs Work (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_about-scps.html)
//    in the AWS Organizations User Guide.
//
//    * For information about attaching tag policies, see How Policy Inheritance
//    Works (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies-inheritance.html)
//    in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using AttachPolicyRequest.
//    req := client.AttachPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/AttachPolicy
func (c *Client) AttachPolicyRequest(input *AttachPolicyInput) AttachPolicyRequest {
	op := &aws.Operation{
		Name:       opAttachPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachPolicyInput{}
	}

	req := c.newRequest(op, input, &AttachPolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AttachPolicyRequest{Request: req, Input: input, Copy: c.AttachPolicyRequest}
}

// AttachPolicyRequest is the request type for the
// AttachPolicy API operation.
type AttachPolicyRequest struct {
	*aws.Request
	Input *AttachPolicyInput
	Copy  func(*AttachPolicyInput) AttachPolicyRequest
}

// Send marshals and sends the AttachPolicy API request.
func (r AttachPolicyRequest) Send(ctx context.Context) (*AttachPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachPolicyResponse{
		AttachPolicyOutput: r.Request.Data.(*AttachPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachPolicyResponse is the response type for the
// AttachPolicy API operation.
type AttachPolicyResponse struct {
	*AttachPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachPolicy request.
func (r *AttachPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
