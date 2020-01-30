// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateFleetCapacityInput struct {
	_ struct{} `type:"structure"`

	// Number of EC2 instances you want this fleet to host.
	DesiredInstances *int64 `type:"integer"`

	// A unique identifier for a fleet to update capacity for. You can use either
	// the fleet ID or ARN value.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// The maximum value allowed for the fleet's instance count. Default if not
	// set is 1.
	MaxSize *int64 `type:"integer"`

	// The minimum value allowed for the fleet's instance count. Default if not
	// set is 0.
	MinSize *int64 `type:"integer"`
}

// String returns the string representation
func (s UpdateFleetCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFleetCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFleetCapacityInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateFleetCapacityOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet that was updated.
	FleetId *string `type:"string"`
}

// String returns the string representation
func (s UpdateFleetCapacityOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateFleetCapacity = "UpdateFleetCapacity"

// UpdateFleetCapacityRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates capacity settings for a fleet. Use this action to specify the number
// of EC2 instances (hosts) that you want this fleet to contain. Before calling
// this action, you may want to call DescribeEC2InstanceLimits to get the maximum
// capacity based on the fleet's EC2 instance type.
//
// Specify minimum and maximum number of instances. Amazon GameLift will not
// change fleet capacity to values fall outside of this range. This is particularly
// important when using auto-scaling (see PutScalingPolicy) to allow capacity
// to adjust based on player demand while imposing limits on automatic adjustments.
//
// To update fleet capacity, specify the fleet ID and the number of instances
// you want the fleet to host. If successful, Amazon GameLift starts or terminates
// instances so that the fleet's active instance count matches the desired instance
// count. You can view a fleet's current capacity information by calling DescribeFleetCapacity.
// If the desired instance count is higher than the instance type's limit, the
// "Limit Exceeded" exception occurs.
//
// Learn more
//
//  Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * DescribeFleetAttributes
//
//    * Update fleets: UpdateFleetAttributes UpdateFleetCapacity UpdateFleetPortSettings
//    UpdateRuntimeConfiguration
//
//    * Manage fleet actions: StartFleetActions StopFleetActions
//
//    // Example sending a request using UpdateFleetCapacityRequest.
//    req := client.UpdateFleetCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateFleetCapacity
func (c *Client) UpdateFleetCapacityRequest(input *UpdateFleetCapacityInput) UpdateFleetCapacityRequest {
	op := &aws.Operation{
		Name:       opUpdateFleetCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateFleetCapacityInput{}
	}

	req := c.newRequest(op, input, &UpdateFleetCapacityOutput{})
	return UpdateFleetCapacityRequest{Request: req, Input: input, Copy: c.UpdateFleetCapacityRequest}
}

// UpdateFleetCapacityRequest is the request type for the
// UpdateFleetCapacity API operation.
type UpdateFleetCapacityRequest struct {
	*aws.Request
	Input *UpdateFleetCapacityInput
	Copy  func(*UpdateFleetCapacityInput) UpdateFleetCapacityRequest
}

// Send marshals and sends the UpdateFleetCapacity API request.
func (r UpdateFleetCapacityRequest) Send(ctx context.Context) (*UpdateFleetCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFleetCapacityResponse{
		UpdateFleetCapacityOutput: r.Request.Data.(*UpdateFleetCapacityOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFleetCapacityResponse is the response type for the
// UpdateFleetCapacity API operation.
type UpdateFleetCapacityResponse struct {
	*UpdateFleetCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFleetCapacity request.
func (r *UpdateFleetCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
