// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type AcceptMatchInput struct {
	_ struct{} `type:"structure"`

	// Player response to the proposed match.
	//
	// AcceptanceType is a required field
	AcceptanceType AcceptanceType `type:"string" required:"true" enum:"true"`

	// A unique identifier for a player delivering the response. This parameter
	// can include one or multiple player IDs.
	//
	// PlayerIds is a required field
	PlayerIds []string `type:"list" required:"true"`

	// A unique identifier for a matchmaking ticket. The ticket must be in status
	// REQUIRES_ACCEPTANCE; otherwise this request will fail.
	//
	// TicketId is a required field
	TicketId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptMatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptMatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptMatchInput"}
	if len(s.AcceptanceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AcceptanceType"))
	}

	if s.PlayerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlayerIds"))
	}

	if s.TicketId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TicketId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptMatchOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptMatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptMatch = "AcceptMatch"

// AcceptMatchRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Registers a player's acceptance or rejection of a proposed FlexMatch match.
// A matchmaking configuration may require player acceptance; if so, then matches
// built with that configuration cannot be completed unless all players accept
// the proposed match within a specified time limit.
//
// When FlexMatch builds a match, all the matchmaking tickets involved in the
// proposed match are placed into status REQUIRES_ACCEPTANCE. This is a trigger
// for your game to get acceptance from all players in the ticket. Acceptances
// are only valid for tickets when they are in this status; all other acceptances
// result in an error.
//
// To register acceptance, specify the ticket ID, a response, and one or more
// players. Once all players have registered acceptance, the matchmaking tickets
// advance to status PLACING, where a new game session is created for the match.
//
// If any player rejects the match, or if acceptances are not received before
// a specified timeout, the proposed match is dropped. The matchmaking tickets
// are then handled in one of two ways: For tickets where one or more players
// rejected the match, the ticket status is returned to SEARCHING to find a
// new match. For tickets where one or more players failed to respond, the ticket
// status is set to CANCELLED, and processing is terminated. A new matchmaking
// request for these players can be submitted as needed.
//
// Learn more
//
//  Add FlexMatch to a Game Client (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-client.html)
//
//  FlexMatch Events Reference (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-events.html)
//
// Related operations
//
//    * StartMatchmaking
//
//    * DescribeMatchmaking
//
//    * StopMatchmaking
//
//    * AcceptMatch
//
//    * StartMatchBackfill
//
//    // Example sending a request using AcceptMatchRequest.
//    req := client.AcceptMatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/AcceptMatch
func (c *Client) AcceptMatchRequest(input *AcceptMatchInput) AcceptMatchRequest {
	op := &aws.Operation{
		Name:       opAcceptMatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptMatchInput{}
	}

	req := c.newRequest(op, input, &AcceptMatchOutput{})
	return AcceptMatchRequest{Request: req, Input: input, Copy: c.AcceptMatchRequest}
}

// AcceptMatchRequest is the request type for the
// AcceptMatch API operation.
type AcceptMatchRequest struct {
	*aws.Request
	Input *AcceptMatchInput
	Copy  func(*AcceptMatchInput) AcceptMatchRequest
}

// Send marshals and sends the AcceptMatch API request.
func (r AcceptMatchRequest) Send(ctx context.Context) (*AcceptMatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptMatchResponse{
		AcceptMatchOutput: r.Request.Data.(*AcceptMatchOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptMatchResponse is the response type for the
// AcceptMatch API operation.
type AcceptMatchResponse struct {
	*AcceptMatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptMatch request.
func (r *AcceptMatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
