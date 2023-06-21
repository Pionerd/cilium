// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified attribute of the specified Amazon FPGA Image (AFI).
func (c *Client) ModifyFpgaImageAttribute(ctx context.Context, params *ModifyFpgaImageAttributeInput, optFns ...func(*Options)) (*ModifyFpgaImageAttributeOutput, error) {
	if params == nil {
		params = &ModifyFpgaImageAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyFpgaImageAttribute", params, optFns, c.addOperationModifyFpgaImageAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyFpgaImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyFpgaImageAttributeInput struct {

	// The ID of the AFI.
	//
	// This member is required.
	FpgaImageId *string

	// The name of the attribute.
	Attribute types.FpgaImageAttributeName

	// A description for the AFI.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The load permission for the AFI.
	LoadPermission *types.LoadPermissionModifications

	// A name for the AFI.
	Name *string

	// The operation type.
	OperationType types.OperationType

	// The product codes. After you add a product code to an AFI, it can't be removed.
	// This parameter is valid only when modifying the productCodes attribute.
	ProductCodes []string

	// The user groups. This parameter is valid only when modifying the loadPermission
	// attribute.
	UserGroups []string

	// The Amazon Web Services account IDs. This parameter is valid only when
	// modifying the loadPermission attribute.
	UserIds []string

	noSmithyDocumentSerde
}

type ModifyFpgaImageAttributeOutput struct {

	// Information about the attribute.
	FpgaImageAttribute *types.FpgaImageAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyFpgaImageAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyFpgaImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyFpgaImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyFpgaImageAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyFpgaImageAttribute(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opModifyFpgaImageAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyFpgaImageAttribute",
	}
}
