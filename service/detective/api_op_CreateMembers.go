// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMembersInput struct {
	_ struct{} `type:"structure"`

	// The list of AWS accounts to invite to become member accounts in the behavior
	// graph. For each invited account, the account list contains the account identifier
	// and the AWS account root user email address.
	//
	// Accounts is a required field
	Accounts []Account `min:"1" type:"list" required:"true"`

	// The ARN of the behavior graph to invite the member accounts to contribute
	// their data to.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`

	// Customized message text to include in the invitation email message to the
	// invited member accounts.
	Message *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMembersInput"}

	if s.Accounts == nil {
		invalidParams.Add(aws.NewErrParamRequired("Accounts"))
	}
	if s.Accounts != nil && len(s.Accounts) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Accounts", 1))
	}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}
	if s.Message != nil && len(*s.Message) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Message", 1))
	}
	if s.Accounts != nil {
		for i, v := range s.Accounts {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Accounts", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Accounts != nil {
		v := s.Accounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Accounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMembersOutput struct {
	_ struct{} `type:"structure"`

	// The set of member account invitation requests that Detective was able to
	// process. This includes accounts that are being verified, that failed verification,
	// and that passed verification and are being sent an invitation.
	Members []MemberDetail `type:"list"`

	// The list of accounts for which Detective was unable to process the invitation
	// request. For each account, the list provides the reason why the request could
	// not be processed. The list includes accounts that are already member accounts
	// in the behavior graph.
	UnprocessedAccounts []UnprocessedAccount `type:"list"`
}

// String returns the string representation
func (s CreateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Members != nil {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Members", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreateMembers = "CreateMembers"

// CreateMembersRequest returns a request value for making API operation for
// Amazon Detective.
//
// Amazon Detective is currently in preview.
//
// Sends a request to invite the specified AWS accounts to be member accounts
// in the behavior graph. This operation can only be called by the master account
// for a behavior graph.
//
// CreateMembers verifies the accounts and then sends invitations to the verified
// accounts.
//
// The request provides the behavior graph ARN and the list of accounts to invite.
//
// The response separates the requested accounts into two lists:
//
//    * The accounts that CreateMembers was able to start the verification for.
//    This list includes member accounts that are being verified, that have
//    passed verification and are being sent an invitation, and that have failed
//    verification.
//
//    * The accounts that CreateMembers was unable to process. This list includes
//    accounts that were already invited to be member accounts in the behavior
//    graph.
//
//    // Example sending a request using CreateMembersRequest.
//    req := client.CreateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/CreateMembers
func (c *Client) CreateMembersRequest(input *CreateMembersInput) CreateMembersRequest {
	op := &aws.Operation{
		Name:       opCreateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/graph/members",
	}

	if input == nil {
		input = &CreateMembersInput{}
	}

	req := c.newRequest(op, input, &CreateMembersOutput{})
	return CreateMembersRequest{Request: req, Input: input, Copy: c.CreateMembersRequest}
}

// CreateMembersRequest is the request type for the
// CreateMembers API operation.
type CreateMembersRequest struct {
	*aws.Request
	Input *CreateMembersInput
	Copy  func(*CreateMembersInput) CreateMembersRequest
}

// Send marshals and sends the CreateMembers API request.
func (r CreateMembersRequest) Send(ctx context.Context) (*CreateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMembersResponse{
		CreateMembersOutput: r.Request.Data.(*CreateMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMembersResponse is the response type for the
// CreateMembers API operation.
type CreateMembersResponse struct {
	*CreateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMembers request.
func (r *CreateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
