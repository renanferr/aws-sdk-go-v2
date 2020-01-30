// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeMatchmakingConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages. This parameter is limited to
	// 10.
	Limit *int64 `min:"1" type:"integer"`

	// A unique identifier for a matchmaking configuration(s) to retrieve. You can
	// use either the configuration name or ARN value. To request all existing configurations,
	// leave this parameter empty.
	Names []string `type:"list"`

	// A token that indicates the start of the next sequential page of results.
	// Use the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// A unique identifier for a matchmaking rule set. You can use either the rule
	// set name or ARN value. Use this parameter to retrieve all matchmaking configurations
	// that use this rule set.
	RuleSetName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeMatchmakingConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMatchmakingConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMatchmakingConfigurationsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.RuleSetName != nil && len(*s.RuleSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleSetName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeMatchmakingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// A collection of requested matchmaking configurations.
	Configurations []MatchmakingConfiguration `type:"list"`

	// A token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeMatchmakingConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMatchmakingConfigurations = "DescribeMatchmakingConfigurations"

// DescribeMatchmakingConfigurationsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the details of FlexMatch matchmaking configurations. With this
// operation, you have the following options: (1) retrieve all existing configurations,
// (2) provide the names of one or more configurations to retrieve, or (3) retrieve
// all configurations that use a specified rule set name. When requesting multiple
// items, use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, a configuration is returned for each requested name.
// When specifying a list of names, only configurations that currently exist
// are returned.
//
// Learn more
//
//  Setting Up FlexMatch Matchmakers (https://docs.aws.amazon.com/gamelift/latest/developerguide/matchmaker-build.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using DescribeMatchmakingConfigurationsRequest.
//    req := client.DescribeMatchmakingConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeMatchmakingConfigurations
func (c *Client) DescribeMatchmakingConfigurationsRequest(input *DescribeMatchmakingConfigurationsInput) DescribeMatchmakingConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeMatchmakingConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMatchmakingConfigurationsInput{}
	}

	req := c.newRequest(op, input, &DescribeMatchmakingConfigurationsOutput{})
	return DescribeMatchmakingConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeMatchmakingConfigurationsRequest}
}

// DescribeMatchmakingConfigurationsRequest is the request type for the
// DescribeMatchmakingConfigurations API operation.
type DescribeMatchmakingConfigurationsRequest struct {
	*aws.Request
	Input *DescribeMatchmakingConfigurationsInput
	Copy  func(*DescribeMatchmakingConfigurationsInput) DescribeMatchmakingConfigurationsRequest
}

// Send marshals and sends the DescribeMatchmakingConfigurations API request.
func (r DescribeMatchmakingConfigurationsRequest) Send(ctx context.Context) (*DescribeMatchmakingConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMatchmakingConfigurationsResponse{
		DescribeMatchmakingConfigurationsOutput: r.Request.Data.(*DescribeMatchmakingConfigurationsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMatchmakingConfigurationsResponse is the response type for the
// DescribeMatchmakingConfigurations API operation.
type DescribeMatchmakingConfigurationsResponse struct {
	*DescribeMatchmakingConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMatchmakingConfigurations request.
func (r *DescribeMatchmakingConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
