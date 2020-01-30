// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type RejectInvitationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the behavior graph to reject the invitation to.
	//
	// The member account's current member status in the behavior graph must be
	// INVITED.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RejectInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectInvitationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectInvitationInput"}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RejectInvitationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RejectInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRejectInvitation = "RejectInvitation"

// RejectInvitationRequest returns a request value for making API operation for
// Amazon Detective.
//
// Amazon Detective is currently in preview.
//
// Rejects an invitation to contribute the account data to a behavior graph.
// This operation must be called by a member account that has the INVITED status.
//
//    // Example sending a request using RejectInvitationRequest.
//    req := client.RejectInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/RejectInvitation
func (c *Client) RejectInvitationRequest(input *RejectInvitationInput) RejectInvitationRequest {
	op := &aws.Operation{
		Name:       opRejectInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/invitation/removal",
	}

	if input == nil {
		input = &RejectInvitationInput{}
	}

	req := c.newRequest(op, input, &RejectInvitationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RejectInvitationRequest{Request: req, Input: input, Copy: c.RejectInvitationRequest}
}

// RejectInvitationRequest is the request type for the
// RejectInvitation API operation.
type RejectInvitationRequest struct {
	*aws.Request
	Input *RejectInvitationInput
	Copy  func(*RejectInvitationInput) RejectInvitationRequest
}

// Send marshals and sends the RejectInvitation API request.
func (r RejectInvitationRequest) Send(ctx context.Context) (*RejectInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectInvitationResponse{
		RejectInvitationOutput: r.Request.Data.(*RejectInvitationOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectInvitationResponse is the response type for the
// RejectInvitation API operation.
type RejectInvitationResponse struct {
	*RejectInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectInvitation request.
func (r *RejectInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
