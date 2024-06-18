// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an AMI. When you're creating an instance-store backed AMI,
// registering the AMI is the final step in the creation process. For more
// information about creating AMIs, see [Create your own AMI]in the Amazon Elastic Compute Cloud User
// Guide.
//
// For Amazon EBS-backed instances, CreateImage creates and registers the AMI in a single
// request, so you don't have to register the AMI yourself. We recommend that you
// always use CreateImageunless you have a specific reason to use RegisterImage.
//
// If needed, you can deregister an AMI at any time. Any modifications you make to
// an AMI backed by an instance store volume invalidates its registration. If you
// make changes to an image, deregister the previous image and register the new
// image.
//
// # Register a snapshot of a root device volume
//
// You can use RegisterImage to create an Amazon EBS-backed Linux AMI from a
// snapshot of a root device volume. You specify the snapshot using a block device
// mapping. You can't set the encryption state of the volume using the block device
// mapping. If the snapshot is encrypted, or encryption by default is enabled, the
// root volume of an instance launched from the AMI is encrypted.
//
// For more information, see [Create a Linux AMI from a snapshot] and [Use encryption with Amazon EBS-backed AMIs] in the Amazon Elastic Compute Cloud User Guide.
//
// # Amazon Web Services Marketplace product codes
//
// If any snapshots have Amazon Web Services Marketplace product codes, they are
// copied to the new AMI.
//
// Windows and some Linux distributions, such as Red Hat Enterprise Linux (RHEL)
// and SUSE Linux Enterprise Server (SLES), use the Amazon EC2 billing product code
// associated with an AMI to verify the subscription status for package updates. To
// create a new AMI for operating systems that require a billing product code,
// instead of registering the AMI, do the following to preserve the billing product
// code association:
//
//   - Launch an instance from an existing AMI with that billing product code.
//
//   - Customize the instance.
//
//   - Create an AMI from the instance using CreateImage.
//
// If you purchase a Reserved Instance to apply to an On-Demand Instance that was
// launched from an AMI with a billing product code, make sure that the Reserved
// Instance has the matching billing product code. If you purchase a Reserved
// Instance without the matching billing product code, the Reserved Instance will
// not be applied to the On-Demand Instance. For information about how to obtain
// the platform details and billing information of an AMI, see [Understand AMI billing information]in the Amazon EC2
// User Guide.
//
// [Understand AMI billing information]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html
// [Create a Linux AMI from a snapshot]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html#creating-launching-ami-from-snapshot
// [Use encryption with Amazon EBS-backed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIEncryption.html
// [Create your own AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami.html
func (c *Client) RegisterImage(ctx context.Context, params *RegisterImageInput, optFns ...func(*Options)) (*RegisterImageOutput, error) {
	if params == nil {
		params = &RegisterImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterImage", params, optFns, c.addOperationRegisterImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for RegisterImage.
type RegisterImageInput struct {

	// A name for your AMI.
	//
	// Constraints: 3-128 alphanumeric characters, parentheses (()), square brackets
	// ([]), spaces ( ), periods (.), slashes (/), dashes (-), single quotes ('),
	// at-signs (@), or underscores(_)
	//
	// This member is required.
	Name *string

	// The architecture of the AMI.
	//
	// Default: For Amazon EBS-backed AMIs, i386 . For instance store-backed AMIs, the
	// architecture specified in the manifest file.
	Architecture types.ArchitectureValues

	// The billing product codes. Your account must be authorized to specify billing
	// product codes.
	//
	// If your account is not authorized to specify billing product codes, you can
	// publish AMIs that include billable software and list them on the Amazon Web
	// Services Marketplace. You must first register as a seller on the Amazon Web
	// Services Marketplace. For more information, see [Getting started as a seller]and [AMI-based products] in the Amazon Web Services
	// Marketplace Seller Guide.
	//
	// [Getting started as a seller]: https://docs.aws.amazon.com/marketplace/latest/userguide/user-guide-for-sellers.html
	// [AMI-based products]: https://docs.aws.amazon.com/marketplace/latest/userguide/ami-products.html
	BillingProducts []string

	// The block device mapping entries.
	//
	// If you specify an Amazon EBS volume using the ID of an Amazon EBS snapshot, you
	// can't specify the encryption state of the volume.
	//
	// If you create an AMI on an Outpost, then all backing snapshots must be on the
	// same Outpost or in the Region of that Outpost. AMIs on an Outpost that include
	// local snapshots can be used to launch instances on the same Outpost only. For
	// more information, [Amazon EBS local snapshots on Outposts]in the Amazon EBS User Guide.
	//
	// [Amazon EBS local snapshots on Outposts]: https://docs.aws.amazon.com/ebs/latest/userguide/snapshots-outposts.html#ami
	BlockDeviceMappings []types.BlockDeviceMapping

	// The boot mode of the AMI. A value of uefi-preferred indicates that the AMI
	// supports both UEFI and Legacy BIOS.
	//
	// The operating system contained in the AMI must be configured to support the
	// specified boot mode.
	//
	// For more information, see [Boot modes] in the Amazon EC2 User Guide.
	//
	// [Boot modes]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html
	BootMode types.BootModeValues

	// A description for your AMI.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Set to true to enable enhanced networking with ENA for the AMI and any
	// instances that you launch from the AMI.
	//
	// This option is supported only for HVM AMIs. Specifying this option with a PV
	// AMI can make instances launched from the AMI unreachable.
	EnaSupport *bool

	// The full path to your AMI manifest in Amazon S3 storage. The specified bucket
	// must have the aws-exec-read canned access control list (ACL) to ensure that it
	// can be accessed by Amazon EC2. For more information, see [Canned ACLs]in the Amazon S3
	// Service Developer Guide.
	//
	// [Canned ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
	ImageLocation *string

	// Set to v2.0 to indicate that IMDSv2 is specified in the AMI. Instances launched
	// from this AMI will have HttpTokens automatically set to required so that, by
	// default, the instance requires that IMDSv2 is used when requesting instance
	// metadata. In addition, HttpPutResponseHopLimit is set to 2 . For more
	// information, see [Configure the AMI]in the Amazon EC2 User Guide.
	//
	// If you set the value to v2.0 , make sure that your AMI software can support
	// IMDSv2.
	//
	// [Configure the AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration
	ImdsSupport types.ImdsSupportValues

	// The ID of the kernel.
	KernelId *string

	// The ID of the RAM disk.
	RamdiskId *string

	// The device name of the root device volume (for example, /dev/sda1 ).
	RootDeviceName *string

	// Set to simple to enable enhanced networking with the Intel 82599 Virtual
	// Function interface for the AMI and any instances that you launch from the AMI.
	//
	// There is no way to disable sriovNetSupport at this time.
	//
	// This option is supported only for HVM AMIs. Specifying this option with a PV
	// AMI can make instances launched from the AMI unreachable.
	SriovNetSupport *string

	// The tags to apply to the AMI.
	//
	// To tag the AMI, the value for ResourceType must be image . If you specify
	// another value for ResourceType , the request fails.
	//
	// To tag an AMI after it has been registered, see [CreateTags].
	//
	// [CreateTags]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html
	TagSpecifications []types.TagSpecification

	// Set to v2.0 to enable Trusted Platform Module (TPM) support. For more
	// information, see [NitroTPM]in the Amazon EC2 User Guide.
	//
	// [NitroTPM]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html
	TpmSupport types.TpmSupportValues

	// Base64 representation of the non-volatile UEFI variable store. To retrieve the
	// UEFI data, use the [GetInstanceUefiData]command. You can inspect and modify the UEFI data by using
	// the [python-uefivars tool]on GitHub. For more information, see [UEFI Secure Boot] in the Amazon EC2 User Guide.
	//
	// [UEFI Secure Boot]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html
	// [GetInstanceUefiData]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceUefiData
	// [python-uefivars tool]: https://github.com/awslabs/python-uefivars
	UefiData *string

	// The type of virtualization ( hvm | paravirtual ).
	//
	// Default: paravirtual
	VirtualizationType *string

	noSmithyDocumentSerde
}

// Contains the output of RegisterImage.
type RegisterImageOutput struct {

	// The ID of the newly registered AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpRegisterImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterImage"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterImage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRegisterImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterImage",
	}
}
