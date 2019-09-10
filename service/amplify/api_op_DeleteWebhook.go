// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for the delete webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteWebhookRequest
type DeleteWebhookInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for a webhook.
	//
	// WebhookId is a required field
	WebhookId *string `location:"uri" locationName:"webhookId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteWebhookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWebhookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteWebhookInput"}

	if s.WebhookId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebhookId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteWebhookInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.WebhookId != nil {
		v := *s.WebhookId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "webhookId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the delete webhook request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteWebhookResult
type DeleteWebhookOutput struct {
	_ struct{} `type:"structure"`

	// Webhook structure.
	//
	// Webhook is a required field
	Webhook *Webhook `locationName:"webhook" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteWebhookOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteWebhookOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Webhook != nil {
		v := s.Webhook

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "webhook", v, metadata)
	}
	return nil
}

const opDeleteWebhook = "DeleteWebhook"

// DeleteWebhookRequest returns a request value for making API operation for
// AWS Amplify.
//
// Deletes a webhook.
//
//    // Example sending a request using DeleteWebhookRequest.
//    req := client.DeleteWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteWebhook
func (c *Client) DeleteWebhookRequest(input *DeleteWebhookInput) DeleteWebhookRequest {
	op := &aws.Operation{
		Name:       opDeleteWebhook,
		HTTPMethod: "DELETE",
		HTTPPath:   "/webhooks/{webhookId}",
	}

	if input == nil {
		input = &DeleteWebhookInput{}
	}

	req := c.newRequest(op, input, &DeleteWebhookOutput{})
	return DeleteWebhookRequest{Request: req, Input: input, Copy: c.DeleteWebhookRequest}
}

// DeleteWebhookRequest is the request type for the
// DeleteWebhook API operation.
type DeleteWebhookRequest struct {
	*aws.Request
	Input *DeleteWebhookInput
	Copy  func(*DeleteWebhookInput) DeleteWebhookRequest
}

// Send marshals and sends the DeleteWebhook API request.
func (r DeleteWebhookRequest) Send(ctx context.Context) (*DeleteWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteWebhookResponse{
		DeleteWebhookOutput: r.Request.Data.(*DeleteWebhookOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteWebhookResponse is the response type for the
// DeleteWebhook API operation.
type DeleteWebhookResponse struct {
	*DeleteWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteWebhook request.
func (r *DeleteWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}