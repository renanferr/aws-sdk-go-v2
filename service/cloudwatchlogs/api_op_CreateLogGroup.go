// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type CreateLogGroupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	// For more information, see Amazon Resource Names - AWS Key Management Service
	// (AWS KMS) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The key-value pairs to use for the tags.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateLogGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLogGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLogGroupInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLogGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLogGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLogGroup = "CreateLogGroup"

// CreateLogGroupRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates a log group with the specified name.
//
// You can create up to 20,000 log groups per account.
//
// You must use the following guidelines when naming a log group:
//
//    * Log group names must be unique within a region for an AWS account.
//
//    * Log group names can be between 1 and 512 characters long.
//
//    * Log group names consist of the following characters: a-z, A-Z, 0-9,
//    '_' (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and
//    '#' (number sign)
//
// If you associate a AWS Key Management Service (AWS KMS) customer master key
// (CMK) with the log group, ingested data is encrypted using the CMK. This
// association is stored as long as the data encrypted with the CMK is still
// within Amazon CloudWatch Logs. This enables Amazon CloudWatch Logs to decrypt
// this data whenever it is requested.
//
// If you attempt to associate a CMK with the log group but the CMK does not
// exist or the CMK is disabled, you will receive an InvalidParameterException
// error.
//
// Important: CloudWatch Logs supports only symmetric CMKs. Do not associate
// an asymmetric CMK with your log group. For more information, see Using Symmetric
// and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html).
//
//    // Example sending a request using CreateLogGroupRequest.
//    req := client.CreateLogGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/CreateLogGroup
func (c *Client) CreateLogGroupRequest(input *CreateLogGroupInput) CreateLogGroupRequest {
	op := &aws.Operation{
		Name:       opCreateLogGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLogGroupInput{}
	}

	req := c.newRequest(op, input, &CreateLogGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateLogGroupRequest{Request: req, Input: input, Copy: c.CreateLogGroupRequest}
}

// CreateLogGroupRequest is the request type for the
// CreateLogGroup API operation.
type CreateLogGroupRequest struct {
	*aws.Request
	Input *CreateLogGroupInput
	Copy  func(*CreateLogGroupInput) CreateLogGroupRequest
}

// Send marshals and sends the CreateLogGroup API request.
func (r CreateLogGroupRequest) Send(ctx context.Context) (*CreateLogGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLogGroupResponse{
		CreateLogGroupOutput: r.Request.Data.(*CreateLogGroupOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLogGroupResponse is the response type for the
// CreateLogGroup API operation.
type CreateLogGroupResponse struct {
	*CreateLogGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLogGroup request.
func (r *CreateLogGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
