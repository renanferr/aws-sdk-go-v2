// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to obtain information about a configuration set.
type GetConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to obtain more information
	// about.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConfigurationSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a configuration set.
type GetConfigurationSetOutput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set.
	ConfigurationSetName *string `type:"string"`

	// An object that defines the dedicated IP pool that is used to send emails
	// that you send using the configuration set.
	DeliveryOptions *DeliveryOptions `type:"structure"`

	// An object that defines whether or not Amazon SES collects reputation metrics
	// for the emails that you send that use the configuration set.
	ReputationOptions *ReputationOptions `type:"structure"`

	// An object that defines whether or not Amazon SES can send email that you
	// send using the configuration set.
	SendingOptions *SendingOptions `type:"structure"`

	// An object that contains information about the suppression list preferences
	// for your account.
	SuppressionOptions *SuppressionOptions `type:"structure"`

	// An array of objects that define the tags (keys and values) that are associated
	// with the configuration set.
	Tags []Tag `type:"list"`

	// An object that defines the open and click tracking options for emails that
	// you send using the configuration set.
	TrackingOptions *TrackingOptions `type:"structure"`
}

// String returns the string representation
func (s GetConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetConfigurationSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeliveryOptions != nil {
		v := s.DeliveryOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeliveryOptions", v, metadata)
	}
	if s.ReputationOptions != nil {
		v := s.ReputationOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ReputationOptions", v, metadata)
	}
	if s.SendingOptions != nil {
		v := s.SendingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SendingOptions", v, metadata)
	}
	if s.SuppressionOptions != nil {
		v := s.SuppressionOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SuppressionOptions", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.TrackingOptions != nil {
		v := s.TrackingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrackingOptions", v, metadata)
	}
	return nil
}

const opGetConfigurationSet = "GetConfigurationSet"

// GetConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Get information about an existing configuration set, including the dedicated
// IP pool that it's associated with, whether or not it's enabled for sending
// email, and more.
//
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference
// to the configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to
// the email.
//
//    // Example sending a request using GetConfigurationSetRequest.
//    req := client.GetConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/GetConfigurationSet
func (c *Client) GetConfigurationSetRequest(input *GetConfigurationSetInput) GetConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opGetConfigurationSet,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/configuration-sets/{ConfigurationSetName}",
	}

	if input == nil {
		input = &GetConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &GetConfigurationSetOutput{})
	return GetConfigurationSetRequest{Request: req, Input: input, Copy: c.GetConfigurationSetRequest}
}

// GetConfigurationSetRequest is the request type for the
// GetConfigurationSet API operation.
type GetConfigurationSetRequest struct {
	*aws.Request
	Input *GetConfigurationSetInput
	Copy  func(*GetConfigurationSetInput) GetConfigurationSetRequest
}

// Send marshals and sends the GetConfigurationSet API request.
func (r GetConfigurationSetRequest) Send(ctx context.Context) (*GetConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConfigurationSetResponse{
		GetConfigurationSetOutput: r.Request.Data.(*GetConfigurationSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConfigurationSetResponse is the response type for the
// GetConfigurationSet API operation.
type GetConfigurationSetResponse struct {
	*GetConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConfigurationSet request.
func (r *GetConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
