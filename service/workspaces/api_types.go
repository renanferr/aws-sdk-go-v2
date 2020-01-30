// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes a modification to the configuration of Bring Your Own License (BYOL)
// for the specified account.
type AccountModification struct {
	_ struct{} `type:"structure"`

	// The IP address range, specified as an IPv4 CIDR block, for the management
	// network interface used for the account.
	DedicatedTenancyManagementCidrRange *string `type:"string"`

	// The status of BYOL (whether BYOL is being enabled or disabled).
	DedicatedTenancySupport DedicatedTenancySupportResultEnum `type:"string" enum:"true"`

	// The error code that is returned if the configuration of BYOL cannot be modified.
	ErrorCode *string `type:"string"`

	// The text of the error message that is returned if the configuration of BYOL
	// cannot be modified.
	ErrorMessage *string `type:"string"`

	// The state of the modification to the configuration of BYOL.
	ModificationState DedicatedTenancyModificationStateEnum `type:"string" enum:"true"`

	// The timestamp when the modification of the BYOL configuration was started.
	StartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s AccountModification) String() string {
	return awsutil.Prettify(s)
}

// Describes the compute type.
type ComputeType struct {
	_ struct{} `type:"structure"`

	// The compute type.
	Name Compute `type:"string" enum:"true"`
}

// String returns the string representation
func (s ComputeType) String() string {
	return awsutil.Prettify(s)
}

// Describes the default values that are used to create WorkSpaces. For more
// information, see Update Directory Details for Your WorkSpaces (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html).
type DefaultWorkspaceCreationProperties struct {
	_ struct{} `type:"structure"`

	// The identifier of any security groups to apply to WorkSpaces when they are
	// created.
	CustomSecurityGroupId *string `min:"11" type:"string"`

	// The organizational unit (OU) in the directory for the WorkSpace machine accounts.
	DefaultOu *string `type:"string"`

	// Specifies whether to automatically assign an Elastic public IP address to
	// WorkSpaces in this directory by default. If enabled, the Elastic public IP
	// address allows outbound internet access from your WorkSpaces when you’re
	// using an internet gateway in the Amazon VPC in which your WorkSpaces are
	// located. If you're using a Network Address Translation (NAT) gateway for
	// outbound internet access from your VPC, or if your WorkSpaces are in public
	// subnets and you manually assign them Elastic IP addresses, you should disable
	// this setting. This setting applies to new WorkSpaces that you launch or to
	// existing WorkSpaces that you rebuild. For more information, see Configure
	// a VPC for Amazon WorkSpaces (https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-vpc.html).
	EnableInternetAccess *bool `type:"boolean"`

	// Specifies whether maintenance mode is enabled for WorkSpaces. For more information,
	// see WorkSpace Maintenance (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html).
	EnableMaintenanceMode *bool `type:"boolean"`

	// Specifies whether the directory is enabled for Amazon WorkDocs.
	EnableWorkDocs *bool `type:"boolean"`

	// Specifies whether WorkSpace users are local administrators on their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool `type:"boolean"`
}

// String returns the string representation
func (s DefaultWorkspaceCreationProperties) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace that cannot be created.
type FailedCreateWorkspaceRequest struct {
	_ struct{} `type:"structure"`

	// The error code that is returned if the WorkSpace cannot be created.
	ErrorCode *string `type:"string"`

	// The text of the error message that is returned if the WorkSpace cannot be
	// created.
	ErrorMessage *string `type:"string"`

	// Information about the WorkSpace.
	WorkspaceRequest *WorkspaceRequest `type:"structure"`
}

// String returns the string representation
func (s FailedCreateWorkspaceRequest) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace that could not be rebooted. (RebootWorkspaces), rebuilt
// (RebuildWorkspaces), restored (RestoreWorkspace), terminated (TerminateWorkspaces),
// started (StartWorkspaces), or stopped (StopWorkspaces).
type FailedWorkspaceChangeRequest struct {
	_ struct{} `type:"structure"`

	// The error code that is returned if the WorkSpace cannot be rebooted.
	ErrorCode *string `type:"string"`

	// The text of the error message that is returned if the WorkSpace cannot be
	// rebooted.
	ErrorMessage *string `type:"string"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s FailedWorkspaceChangeRequest) String() string {
	return awsutil.Prettify(s)
}

// Describes a rule for an IP access control group.
type IpRuleItem struct {
	_ struct{} `type:"structure"`

	// The IP address range, in CIDR notation.
	IpRule *string `locationName:"ipRule" type:"string"`

	// The description.
	RuleDesc *string `locationName:"ruleDesc" type:"string"`
}

// String returns the string representation
func (s IpRuleItem) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace modification.
type ModificationState struct {
	_ struct{} `type:"structure"`

	// The resource.
	Resource ModificationResourceEnum `type:"string" enum:"true"`

	// The modification state.
	State ModificationStateEnum `type:"string" enum:"true"`
}

// String returns the string representation
func (s ModificationState) String() string {
	return awsutil.Prettify(s)
}

// The operating system that the image is running.
type OperatingSystem struct {
	_ struct{} `type:"structure"`

	// The operating system.
	Type OperatingSystemType `type:"string" enum:"true"`
}

// String returns the string representation
func (s OperatingSystem) String() string {
	return awsutil.Prettify(s)
}

// Describes an Amazon WorkSpaces client.
type Properties struct {
	_ struct{} `type:"structure"`

	// Specifies whether users can cache their credentials on the Amazon WorkSpaces
	// client. When enabled, users can choose to reconnect to their WorkSpaces without
	// re-entering their credentials.
	ReconnectEnabled ReconnectEnum `type:"string" enum:"true"`
}

// String returns the string representation
func (s Properties) String() string {
	return awsutil.Prettify(s)
}

// Information about the Amazon WorkSpaces client.
type PropertiesResult struct {
	_ struct{} `type:"structure"`

	// Information about the Amazon WorkSpaces client.
	ClientProperties *Properties `type:"structure"`

	// The resource identifier, in the form of a directory ID.
	ResourceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PropertiesResult) String() string {
	return awsutil.Prettify(s)
}

// Describes the information used to reboot a WorkSpace.
type RebootRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebootRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootRequest"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the information used to rebuild a WorkSpace.
type RebuildRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebuildRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebuildRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebuildRequest"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the root volume for a WorkSpace bundle.
type RootStorage struct {
	_ struct{} `type:"structure"`

	// The size of the root volume.
	Capacity *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RootStorage) String() string {
	return awsutil.Prettify(s)
}

// Describes the self-service permissions for a directory. For more information,
// see Enable Self-Service WorkSpace Management Capabilities for Your Users
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/enable-user-self-service-workspace-management.html).
type SelfservicePermissions struct {
	_ struct{} `type:"structure"`

	// Specifies whether users can change the compute type (bundle) for their WorkSpace.
	ChangeComputeType ReconnectEnum `type:"string" enum:"true"`

	// Specifies whether users can increase the volume size of the drives on their
	// WorkSpace.
	IncreaseVolumeSize ReconnectEnum `type:"string" enum:"true"`

	// Specifies whether users can rebuild the operating system of a WorkSpace to
	// its original state.
	RebuildWorkspace ReconnectEnum `type:"string" enum:"true"`

	// Specifies whether users can restart their WorkSpace.
	RestartWorkspace ReconnectEnum `type:"string" enum:"true"`

	// Specifies whether users can switch the running mode of their WorkSpace.
	SwitchRunningMode ReconnectEnum `type:"string" enum:"true"`
}

// String returns the string representation
func (s SelfservicePermissions) String() string {
	return awsutil.Prettify(s)
}

// Describes a snapshot.
type Snapshot struct {
	_ struct{} `type:"structure"`

	// The time when the snapshot was created.
	SnapshotTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s Snapshot) String() string {
	return awsutil.Prettify(s)
}

// Information used to start a WorkSpace.
type StartRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s StartRequest) String() string {
	return awsutil.Prettify(s)
}

// Describes the information used to stop a WorkSpace.
type StopRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s StopRequest) String() string {
	return awsutil.Prettify(s)
}

// Describes a tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The value of the tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the information used to terminate a WorkSpace.
type TerminateRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace.
	//
	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateRequest"}

	if s.WorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the user storage for a WorkSpace bundle.
type UserStorage struct {
	_ struct{} `type:"structure"`

	// The size of the user storage.
	Capacity *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UserStorage) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace.
type Workspace struct {
	_ struct{} `type:"structure"`

	// The identifier of the bundle used to create the WorkSpace.
	BundleId *string `type:"string"`

	// The name of the WorkSpace, as seen by the operating system.
	ComputerName *string `type:"string"`

	// The identifier of the AWS Directory Service directory for the WorkSpace.
	DirectoryId *string `min:"10" type:"string"`

	// The error code that is returned if the WorkSpace cannot be created.
	ErrorCode *string `type:"string"`

	// The text of the error message that is returned if the WorkSpace cannot be
	// created.
	ErrorMessage *string `type:"string"`

	// The IP address of the WorkSpace.
	IpAddress *string `type:"string"`

	// The modification states of the WorkSpace.
	ModificationStates []ModificationState `type:"list"`

	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool `type:"boolean"`

	// The operational state of the WorkSpace.
	State WorkspaceState `type:"string" enum:"true"`

	// The identifier of the subnet for the WorkSpace.
	SubnetId *string `min:"15" type:"string"`

	// The user for the WorkSpace.
	UserName *string `min:"1" type:"string"`

	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool `type:"boolean"`

	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored
	// on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey *string `type:"string"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`

	// The properties of the WorkSpace.
	WorkspaceProperties *WorkspaceProperties `type:"structure"`
}

// String returns the string representation
func (s Workspace) String() string {
	return awsutil.Prettify(s)
}

// The device types and operating systems that can be used to access a WorkSpace.
// For more information, see Amazon WorkSpaces Client Network Requirements (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-network-requirements.html).
type WorkspaceAccessProperties struct {
	_ struct{} `type:"structure"`

	// Indicates whether users can use Android devices to access their WorkSpaces.
	DeviceTypeAndroid AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can use Chromebooks to access their WorkSpaces.
	DeviceTypeChromeOs AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can use iOS devices to access their WorkSpaces.
	DeviceTypeIos AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can use macOS clients to access their WorkSpaces.
	// To restrict WorkSpaces access to trusted devices (also known as managed devices)
	// with valid certificates, specify a value of TRUST. For more information,
	// see Restrict WorkSpaces Access to Trusted Devices (https://docs.aws.amazon.com/workspaces/latest/adminguide/trusted-devices.html).
	DeviceTypeOsx AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can access their WorkSpaces through a web browser.
	DeviceTypeWeb AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can use Windows clients to access their WorkSpaces.
	// To restrict WorkSpaces access to trusted devices (also known as managed devices)
	// with valid certificates, specify a value of TRUST. For more information,
	// see Restrict WorkSpaces Access to Trusted Devices (https://docs.aws.amazon.com/workspaces/latest/adminguide/trusted-devices.html).
	DeviceTypeWindows AccessPropertyValue `type:"string" enum:"true"`

	// Indicates whether users can use zero client devices to access their WorkSpaces.
	DeviceTypeZeroClient AccessPropertyValue `type:"string" enum:"true"`
}

// String returns the string representation
func (s WorkspaceAccessProperties) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace bundle.
type WorkspaceBundle struct {
	_ struct{} `type:"structure"`

	// The bundle identifier.
	BundleId *string `type:"string"`

	// The compute type. For more information, see Amazon WorkSpaces Bundles (http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles).
	ComputeType *ComputeType `type:"structure"`

	// A description.
	Description *string `type:"string"`

	// The image identifier of the bundle.
	ImageId *string `type:"string"`

	// The last time that the bundle was updated.
	LastUpdatedTime *time.Time `type:"timestamp"`

	// The name of the bundle.
	Name *string `min:"1" type:"string"`

	// The owner of the bundle. This is the account identifier of the owner, or
	// AMAZON if the bundle is provided by AWS.
	Owner *string `type:"string"`

	// The size of the root volume.
	RootStorage *RootStorage `type:"structure"`

	// The size of the user storage.
	UserStorage *UserStorage `type:"structure"`
}

// String returns the string representation
func (s WorkspaceBundle) String() string {
	return awsutil.Prettify(s)
}

// Describes the connection status of a WorkSpace.
type WorkspaceConnectionStatus struct {
	_ struct{} `type:"structure"`

	// The connection state of the WorkSpace. The connection state is unknown if
	// the WorkSpace is stopped.
	ConnectionState ConnectionState `type:"string" enum:"true"`

	// The timestamp of the connection status check.
	ConnectionStateCheckTimestamp *time.Time `type:"timestamp"`

	// The timestamp of the last known user connection.
	LastKnownUserConnectionTimestamp *time.Time `type:"timestamp"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s WorkspaceConnectionStatus) String() string {
	return awsutil.Prettify(s)
}

// Describes the default properties that are used for creating WorkSpaces. For
// more information, see Update Directory Details for Your WorkSpaces (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html).
type WorkspaceCreationProperties struct {
	_ struct{} `type:"structure"`

	// The identifier of your custom security group.
	CustomSecurityGroupId *string `min:"11" type:"string"`

	// The default organizational unit (OU) for your WorkSpace directories.
	DefaultOu *string `type:"string"`

	// Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess *bool `type:"boolean"`

	// Indicates whether maintenance mode is enabled for your WorkSpaces. For more
	// information, see WorkSpace Maintenance (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html).
	EnableMaintenanceMode *bool `type:"boolean"`

	// Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool `type:"boolean"`
}

// String returns the string representation
func (s WorkspaceCreationProperties) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *WorkspaceCreationProperties) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "WorkspaceCreationProperties"}
	if s.CustomSecurityGroupId != nil && len(*s.CustomSecurityGroupId) < 11 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomSecurityGroupId", 11))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a directory that is used with Amazon WorkSpaces.
type WorkspaceDirectory struct {
	_ struct{} `type:"structure"`

	// The directory alias.
	Alias *string `type:"string"`

	// The user name for the service account.
	CustomerUserName *string `min:"1" type:"string"`

	// The directory identifier.
	DirectoryId *string `min:"10" type:"string"`

	// The name of the directory.
	DirectoryName *string `type:"string"`

	// The directory type.
	DirectoryType WorkspaceDirectoryType `type:"string" enum:"true"`

	// The IP addresses of the DNS servers for the directory.
	DnsIpAddresses []string `type:"list"`

	// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces
	// to make calls to other services, such as Amazon EC2, on your behalf.
	IamRoleId *string `type:"string"`

	// The identifiers of the IP access control groups associated with the directory.
	IpGroupIds []string `locationName:"ipGroupIds" type:"list"`

	// The registration code for the directory. This is the code that users enter
	// in their Amazon WorkSpaces client application to connect to the directory.
	RegistrationCode *string `min:"1" type:"string"`

	// The default self-service permissions for WorkSpaces in the directory.
	SelfservicePermissions *SelfservicePermissions `type:"structure"`

	// The state of the directory's registration with Amazon WorkSpaces.
	State WorkspaceDirectoryState `type:"string" enum:"true"`

	// The identifiers of the subnets used with the directory.
	SubnetIds []string `type:"list"`

	// Specifies whether the directory is dedicated or shared. To use Bring Your
	// Own License (BYOL), this value must be set to DEDICATED. For more information,
	// see Bring Your Own Windows Desktop Images (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
	Tenancy Tenancy `type:"string" enum:"true"`

	// The devices and operating systems that users can use to access WorkSpaces.
	WorkspaceAccessProperties *WorkspaceAccessProperties `type:"structure"`

	// The default creation properties for all WorkSpaces in the directory.
	WorkspaceCreationProperties *DefaultWorkspaceCreationProperties `type:"structure"`

	// The identifier of the security group that is assigned to new WorkSpaces.
	WorkspaceSecurityGroupId *string `min:"11" type:"string"`
}

// String returns the string representation
func (s WorkspaceDirectory) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace image.
type WorkspaceImage struct {
	_ struct{} `type:"structure"`

	// The description of the image.
	Description *string `min:"1" type:"string"`

	// The error code that is returned for the image.
	ErrorCode *string `type:"string"`

	// The text of the error message that is returned for the image.
	ErrorMessage *string `type:"string"`

	// The identifier of the image.
	ImageId *string `type:"string"`

	// The name of the image.
	Name *string `min:"1" type:"string"`

	// The operating system that the image is running.
	OperatingSystem *OperatingSystem `type:"structure"`

	// Specifies whether the image is running on dedicated hardware. When Bring
	// Your Own License (BYOL) is enabled, this value is set to DEDICATED. For more
	// information, see Bring Your Own Windows Desktop Images (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
	RequiredTenancy WorkspaceImageRequiredTenancy `type:"string" enum:"true"`

	// The status of the image.
	State WorkspaceImageState `type:"string" enum:"true"`
}

// String returns the string representation
func (s WorkspaceImage) String() string {
	return awsutil.Prettify(s)
}

// Describes a WorkSpace.
type WorkspaceProperties struct {
	_ struct{} `type:"structure"`

	// The compute type. For more information, see Amazon WorkSpaces Bundles (http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles).
	ComputeTypeName Compute `type:"string" enum:"true"`

	// The size of the root volume.
	RootVolumeSizeGib *int64 `type:"integer"`

	// The running mode. For more information, see Manage the WorkSpace Running
	// Mode (https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html).
	RunningMode RunningMode `type:"string" enum:"true"`

	// The time after a user logs off when WorkSpaces are automatically stopped.
	// Configured in 60-minute intervals.
	RunningModeAutoStopTimeoutInMinutes *int64 `type:"integer"`

	// The size of the user storage.
	UserVolumeSizeGib *int64 `type:"integer"`
}

// String returns the string representation
func (s WorkspaceProperties) String() string {
	return awsutil.Prettify(s)
}

// Describes the information used to create a WorkSpace.
type WorkspaceRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the bundle for the WorkSpace. You can use DescribeWorkspaceBundles
	// to list the available bundles.
	//
	// BundleId is a required field
	BundleId *string `type:"string" required:"true"`

	// The identifier of the AWS Directory Service directory for the WorkSpace.
	// You can use DescribeWorkspaceDirectories to list the available directories.
	//
	// DirectoryId is a required field
	DirectoryId *string `min:"10" type:"string" required:"true"`

	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool `type:"boolean"`

	// The tags for the WorkSpace.
	Tags []Tag `type:"list"`

	// The user name of the user for the WorkSpace. This user name must exist in
	// the AWS Directory Service directory for the WorkSpace.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`

	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool `type:"boolean"`

	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored
	// on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey *string `type:"string"`

	// The WorkSpace properties.
	WorkspaceProperties *WorkspaceProperties `type:"structure"`
}

// String returns the string representation
func (s WorkspaceRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *WorkspaceRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "WorkspaceRequest"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.DirectoryId != nil && len(*s.DirectoryId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("DirectoryId", 10))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes an IP access control group.
type WorkspacesIpGroup struct {
	_ struct{} `type:"structure"`

	// The description of the group.
	GroupDesc *string `locationName:"groupDesc" type:"string"`

	// The identifier of the group.
	GroupId *string `locationName:"groupId" type:"string"`

	// The name of the group.
	GroupName *string `locationName:"groupName" type:"string"`

	// The rules.
	UserRules []IpRuleItem `locationName:"userRules" type:"list"`
}

// String returns the string representation
func (s WorkspacesIpGroup) String() string {
	return awsutil.Prettify(s)
}
