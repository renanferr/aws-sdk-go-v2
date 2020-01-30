// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Amazon Detective is currently in preview.
//
// An AWS account that is the master of or a member of a behavior graph.
type Account struct {
	_ struct{} `type:"structure"`

	// The account identifier of the AWS account.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The AWS account root user email address for the AWS account.
	//
	// EmailAddress is a required field
	EmailAddress *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Account) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Account) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Account"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}
	if s.EmailAddress != nil && len(*s.EmailAddress) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailAddress", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Account) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Amazon Detective is currently in preview.
//
// A behavior graph in Detective.
type Graph struct {
	_ struct{} `type:"structure"`

	// The ARN of the behavior graph.
	Arn *string `type:"string"`

	// The date and time that the behavior graph was created. The value is in milliseconds
	// since the epoch.
	CreatedTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s Graph) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Graph) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedTime != nil {
		v := *s.CreatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	return nil
}

// Amazon Detective is currently in preview.
//
// Details about a member account that was invited to contribute to a behavior
// graph.
type MemberDetail struct {
	_ struct{} `type:"structure"`

	// The AWS account identifier for the member account.
	AccountId *string `min:"12" type:"string"`

	// The AWS account root user email address for the member account.
	EmailAddress *string `min:"1" type:"string"`

	// The ARN of the behavior graph that the member account was invited to.
	GraphArn *string `type:"string"`

	// The date and time that Detective sent the invitation to the member account.
	// The value is in milliseconds since the epoch.
	InvitedTime *time.Time `type:"timestamp"`

	// The AWS account identifier of the master account for the behavior graph.
	MasterId *string `min:"12" type:"string"`

	// The current membership status of the member account. The status can have
	// one of the following values:
	//
	//    * INVITED - Indicates that the member was sent an invitation but has not
	//    yet responded.
	//
	//    * VERIFICATION_IN_PROGRESS - Indicates that Detective is verifying that
	//    the account identifier and email address provided for the member account
	//    match. If they do match, then Detective sends the invitation. If the email
	//    address and account identifier don't match, then the member cannot be
	//    added to the behavior graph.
	//
	//    * VERIFICATION_FAILED - Indicates that the account and email address provided
	//    for the member account do not match, and Detective did not send an invitation
	//    to the account.
	//
	//    * ENABLED - Indicates that the member account accepted the invitation
	//    to contribute to the behavior graph.
	//
	// Member accounts that declined an invitation or that were removed from the
	// behavior graph are not included.
	Status MemberStatus `type:"string" enum:"true"`

	// The date and time that the member account was last updated. The value is
	// in milliseconds since the epoch.
	UpdatedTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s MemberDetail) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MemberDetail) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InvitedTime != nil {
		v := *s.InvitedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InvitedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.MasterId != nil {
		v := *s.MasterId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MasterId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.UpdatedTime != nil {
		v := *s.UpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	return nil
}

// Amazon Detective is currently in preview.
//
// A member account that was included in a request but for which the request
// could not be processed.
type UnprocessedAccount struct {
	_ struct{} `type:"structure"`

	// The AWS account identifier of the member account that was not processed.
	AccountId *string `min:"12" type:"string"`

	// The reason that the member account request could not be processed.
	Reason *string `type:"string"`
}

// String returns the string representation
func (s UnprocessedAccount) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UnprocessedAccount) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Reason != nil {
		v := *s.Reason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Reason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
