// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/vault-client-go/schema"
)

// Secrets is a simple wrapper around the client for Secrets requests
type Secrets struct {
	client *Client
}

// AWSConfigReadLease Configure the default lease information for generated credentials.
func (a *Secrets) AWSConfigReadLease(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigReadRootIAMCredentials Configure the root credentials that are used to manage IAM.
func (a *Secrets) AWSConfigReadRootIAMCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigRotateRootIAMCredentials
func (a *Secrets) AWSConfigRotateRootIAMCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteLease Configure the default lease information for generated credentials.
func (a *Secrets) AWSConfigWriteLease(ctx context.Context, request schema.AWSConfigWriteLeaseRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSConfigWriteRootIAMCredentials Configure the root credentials that are used to manage IAM.
func (a *Secrets) AWSConfigWriteRootIAMCredentials(ctx context.Context, request schema.AWSConfigWriteRootIAMCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/config/root"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSDeleteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (a *Secrets) AWSDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSListRoles List the existing roles in this backend
func (a *Secrets) AWSListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadCredentials Generate AWS credentials from a specific Vault role.
func (a *Secrets) AWSReadCredentials(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (a *Secrets) AWSReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSReadSecurityTokenService Generate AWS credentials from a specific Vault role.
// name: Name of the role
func (a *Secrets) AWSReadSecurityTokenService(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteCredentials Generate AWS credentials from a specific Vault role.
func (a *Secrets) AWSWriteCredentials(ctx context.Context, request schema.AWSWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/creds"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteRole Read, write and reference IAM policies that access keys can be made for.
// name: Name of the policy
func (a *Secrets) AWSWriteRole(ctx context.Context, name string, request schema.AWSWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AWSWriteSecurityTokenService Generate AWS credentials from a specific Vault role.
// name: Name of the role
func (a *Secrets) AWSWriteSecurityTokenService(ctx context.Context, name string, request schema.AWSWriteSecurityTokenServiceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{aws_mount_path}/sts/{name}"
	requestPath = strings.Replace(requestPath, "{"+"aws_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("aws")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryCheckInLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryCheckInManageLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (a *Secrets) ActiveDirectoryCheckOutLibrary(ctx context.Context, name string, request schema.ActiveDirectoryCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryDeleteConfig Configure the AD server to connect to, along with password options.
func (a *Secrets) ActiveDirectoryDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryDeleteLibrary Delete a library set.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryDeleteRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (a *Secrets) ActiveDirectoryDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryListLibraries
func (a *Secrets) ActiveDirectoryListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryListRoles List the name of each role currently stored.
func (a *Secrets) ActiveDirectoryListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryReadConfig Configure the AD server to connect to, along with password options.
func (a *Secrets) ActiveDirectoryReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryReadCredentials
// name: Name of the role
func (a *Secrets) ActiveDirectoryReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryReadLibrary Read a library set.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryReadRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (a *Secrets) ActiveDirectoryReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryRotateRole
// name: Name of the static role
func (a *Secrets) ActiveDirectoryRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryRotateRoot
func (a *Secrets) ActiveDirectoryRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryWriteConfig Configure the AD server to connect to, along with password options.
func (a *Secrets) ActiveDirectoryWriteConfig(ctx context.Context, request schema.ActiveDirectoryWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryWriteLibrary Update a library set.
// name: Name of the set.
func (a *Secrets) ActiveDirectoryWriteLibrary(ctx context.Context, name string, request schema.ActiveDirectoryWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ActiveDirectoryWriteRole Manage roles to build links between Vault and Active Directory service accounts.
// name: Name of the role
func (a *Secrets) ActiveDirectoryWriteRole(ctx context.Context, name string, request schema.ActiveDirectoryWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ad_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudDeleteConfig Configure the access key and secret to use for RAM and STS calls.
func (a *Secrets) AliCloudDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudDeleteRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
func (a *Secrets) AliCloudDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudListRoles List the existing roles in this backend.
func (a *Secrets) AliCloudListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadConfig Configure the access key and secret to use for RAM and STS calls.
func (a *Secrets) AliCloudReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadCredentials Generate an API key or STS credential using the given role's configuration.'
// name: The name of the role.
func (a *Secrets) AliCloudReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudReadRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
func (a *Secrets) AliCloudReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudWriteConfig Configure the access key and secret to use for RAM and STS calls.
func (a *Secrets) AliCloudWriteConfig(ctx context.Context, request schema.AliCloudWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AliCloudWriteRole Read, write and reference policies and roles that API keys or STS credentials can be made for.
// name: The name of the role.
func (a *Secrets) AliCloudWriteRole(ctx context.Context, name string, request schema.AliCloudWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{alicloud_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"alicloud_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("alicloud")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteConfig
func (a *Secrets) AzureDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureDeleteRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
func (a *Secrets) AzureDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureListRoles List existing roles.
func (a *Secrets) AzureListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadConfig
func (a *Secrets) AzureReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadCredentials
// role: Name of the Vault role
func (a *Secrets) AzureReadCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureReadRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
func (a *Secrets) AzureReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureRotateRoot
func (a *Secrets) AzureRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteConfig
func (a *Secrets) AzureWriteConfig(ctx context.Context, request schema.AzureWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// AzureWriteRole Manage the Vault roles used to generate Azure credentials.
// name: Name of the role.
func (a *Secrets) AzureWriteRole(ctx context.Context, name string, request schema.AzureWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{azure_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"azure_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("azure")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulDeleteRole
// name: Name of the role.
func (a *Secrets) ConsulDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulListRoles
func (a *Secrets) ConsulListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulReadAccessConfig
func (a *Secrets) ConsulReadAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulReadCredentials
// role: Name of the role.
func (a *Secrets) ConsulReadCredentials(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulReadRole
// name: Name of the role.
func (a *Secrets) ConsulReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulWriteAccessConfig
func (a *Secrets) ConsulWriteAccessConfig(ctx context.Context, request schema.ConsulWriteAccessConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// ConsulWriteRole
// name: Name of the role.
func (a *Secrets) ConsulWriteRole(ctx context.Context, name string, request schema.ConsulWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{consul_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"consul_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("consul")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeDelete Deletes the secret at the specified location.
// path: Specifies the path of the secret.
func (a *Secrets) CubbyholeDelete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeRead Retrieve the secret at the specified location.
// path: Specifies the path of the secret.
// list: Return a list if &#x60;true&#x60;
func (a *Secrets) CubbyholeRead(ctx context.Context, path string, list string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", url.QueryEscape(list))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// CubbyholeWrite Store a secret at the specified location.
// path: Specifies the path of the secret.
func (a *Secrets) CubbyholeWrite(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{cubbyhole_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"cubbyhole_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("cubbyhole")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteRoleset
// name: Required. Name of the role.
func (a *Secrets) GoogleCloudDeleteRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudDeleteStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
func (a *Secrets) GoogleCloudDeleteStaticAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSDecrypt Decrypt a ciphertext value using a named key
// key: Name of the key in Vault to use for decryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (a *Secrets) GoogleCloudKMSDecrypt(ctx context.Context, key string, request schema.GoogleCloudKMSDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/decrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSDeleteConfig Configure the GCP KMS secrets engine
func (a *Secrets) GoogleCloudKMSDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSDeleteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSDeleteKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSDeregisterKey Deregister an existing key in Vault
// key: Name of the key to deregister in Vault. If the key exists in Google Cloud KMS, it will be left untouched.
func (a *Secrets) GoogleCloudKMSDeregisterKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/deregister/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSEncrypt Encrypt a plaintext value using a named key
// key: Name of the key in Vault to use for encryption. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (a *Secrets) GoogleCloudKMSEncrypt(ctx context.Context, key string, request schema.GoogleCloudKMSEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/encrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSListKeys List named keys
func (a *Secrets) GoogleCloudKMSListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSReadConfig Configure the GCP KMS secrets engine
func (a *Secrets) GoogleCloudKMSReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSReadKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSReadKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSReadKeyConfig Configure the key in Vault
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSReadKeyConfig(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSReadPubkey Retrieve the public key associated with the named key
// key: Name of the key for which to get the public key. This key must already exist in Vault and Google Cloud KMS.
func (a *Secrets) GoogleCloudKMSReadPubkey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/pubkey/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSReencrypt Re-encrypt existing ciphertext data to a new version
// key: Name of the key to use for encryption. This key must already exist in Vault and Google Cloud KMS.
func (a *Secrets) GoogleCloudKMSReencrypt(ctx context.Context, key string, request schema.GoogleCloudKMSReencryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/reencrypt/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSRegisterKey Register an existing crypto key in Google Cloud KMS
// key: Name of the key to register in Vault. This will be the named used to refer to the underlying crypto key when encrypting or decrypting data.
func (a *Secrets) GoogleCloudKMSRegisterKey(ctx context.Context, key string, request schema.GoogleCloudKMSRegisterKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/register/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSRotateKey Rotate a crypto key to a new primary version
// key: Name of the key to rotate. This key must already be registered with Vault and point to a valid Google Cloud KMS crypto key.
func (a *Secrets) GoogleCloudKMSRotateKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/rotate/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSSign Signs a message or digest using a named key
// key: Name of the key in Vault to use for signing. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (a *Secrets) GoogleCloudKMSSign(ctx context.Context, key string, request schema.GoogleCloudKMSSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/sign/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSTrimKey Delete old crypto key versions from Google Cloud KMS
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSTrimKey(ctx context.Context, key string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/trim/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSVerify Verify a signature using a named key
// key: Name of the key in Vault to use for verification. This key must already exist in Vault and must map back to a Google Cloud KMS key.
func (a *Secrets) GoogleCloudKMSVerify(ctx context.Context, key string, request schema.GoogleCloudKMSVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/verify/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSWriteConfig Configure the GCP KMS secrets engine
func (a *Secrets) GoogleCloudKMSWriteConfig(ctx context.Context, request schema.GoogleCloudKMSWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSWriteKey Interact with crypto keys in Vault and Google Cloud KMS
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSWriteKey(ctx context.Context, key string, request schema.GoogleCloudKMSWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudKMSWriteKeyConfig Configure the key in Vault
// key: Name of the key in Vault.
func (a *Secrets) GoogleCloudKMSWriteKeyConfig(ctx context.Context, key string, request schema.GoogleCloudKMSWriteKeyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcpkms_mount_path}/keys/config/{key}"
	requestPath = strings.Replace(requestPath, "{"+"gcpkms_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcpkms")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key"+"}", url.PathEscape(key), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListRolesets
func (a *Secrets) GoogleCloudListRolesets(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/rolesets"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudListStaticAccounts
func (a *Secrets) GoogleCloudListStaticAccounts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-accounts"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadConfig
func (a *Secrets) GoogleCloudReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadKey
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudReadKey(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/key/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRoleset
// name: Required. Name of the role.
func (a *Secrets) GoogleCloudReadRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRolesetKey
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudReadRolesetKey(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadRolesetToken
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudReadRolesetToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
func (a *Secrets) GoogleCloudReadStaticAccount(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadStaticAccountKey
// name: Required. Name of the static account.
func (a *Secrets) GoogleCloudReadStaticAccountKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadStaticAccountToken
// name: Required. Name of the static account.
func (a *Secrets) GoogleCloudReadStaticAccountToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudReadToken
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudReadToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/token/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRoleset
// name: Name of the role.
func (a *Secrets) GoogleCloudRotateRoleset(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRolesetKey
// name: Name of the role.
func (a *Secrets) GoogleCloudRotateRolesetKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateRoot
func (a *Secrets) GoogleCloudRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudRotateStaticAccountKey
// name: Name of the account.
func (a *Secrets) GoogleCloudRotateStaticAccountKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/rotate-key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteConfig
func (a *Secrets) GoogleCloudWriteConfig(ctx context.Context, request schema.GoogleCloudWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteKey
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudWriteKey(ctx context.Context, roleset string, request schema.GoogleCloudWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/key/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRoleset
// name: Required. Name of the role.
func (a *Secrets) GoogleCloudWriteRoleset(ctx context.Context, name string, request schema.GoogleCloudWriteRolesetRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRolesetKey
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudWriteRolesetKey(ctx context.Context, roleset string, request schema.GoogleCloudWriteRolesetKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteRolesetToken
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudWriteRolesetToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/roleset/{roleset}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteStaticAccount
// name: Required. Name to refer to this static account in Vault. Cannot be updated.
func (a *Secrets) GoogleCloudWriteStaticAccount(ctx context.Context, name string, request schema.GoogleCloudWriteStaticAccountRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteStaticAccountKey
// name: Required. Name of the static account.
func (a *Secrets) GoogleCloudWriteStaticAccountKey(ctx context.Context, name string, request schema.GoogleCloudWriteStaticAccountKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/key"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteStaticAccountToken
// name: Required. Name of the static account.
func (a *Secrets) GoogleCloudWriteStaticAccountToken(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/static-account/{name}/token"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// GoogleCloudWriteToken
// roleset: Required. Name of the role set.
func (a *Secrets) GoogleCloudWriteToken(ctx context.Context, roleset string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{gcp_mount_path}/token/{roleset}"
	requestPath = strings.Replace(requestPath, "{"+"gcp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("gcp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"roleset"+"}", url.PathEscape(roleset), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv1Delete Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
func (a *Secrets) KVv1Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv1Read Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
// list: Return a list if &#x60;true&#x60;
func (a *Secrets) KVv1Read(ctx context.Context, path string, list string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", url.QueryEscape(list))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv1Write Pass-through secret storage to the storage backend, allowing you to read/write arbitrary data into secret storage.
// path: Location of the secret.
func (a *Secrets) KVv1Write(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kv_mount_path}/{path}"
	requestPath = strings.Replace(requestPath, "{"+"kv_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kv")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2Delete Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (a *Secrets) KVv2Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2DeleteMetadata Configures settings for the KV store
// path: Location of the secret.
func (a *Secrets) KVv2DeleteMetadata(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2DeleteVersions Marks one or more versions as deleted in the KV store.
// path: Location of the secret.
func (a *Secrets) KVv2DeleteVersions(ctx context.Context, path string, request schema.KVv2DeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/delete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2DestroyVersions Permanently removes one or more versions in the KV store
// path: Location of the secret.
func (a *Secrets) KVv2DestroyVersions(ctx context.Context, path string, request schema.KVv2DestroyVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/destroy/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2Read Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (a *Secrets) KVv2Read(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2ReadConfig Read the backend level settings.
func (a *Secrets) KVv2ReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2ReadMetadata Configures settings for the KV store
// path: Location of the secret.
// list: Return a list if &#x60;true&#x60;
func (a *Secrets) KVv2ReadMetadata(ctx context.Context, path string, list string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", url.QueryEscape(list))

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2ReadSubkeys Read the structure of a secret entry from the Key-Value store with the values removed.
// path: Location of the secret.
func (a *Secrets) KVv2ReadSubkeys(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/subkeys/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2UndeleteVersions Undeletes one or more versions from the KV store.
// path: Location of the secret.
func (a *Secrets) KVv2UndeleteVersions(ctx context.Context, path string, request schema.KVv2UndeleteVersionsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/undelete/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2Write Write, Patch, Read, and Delete data in the Key-Value Store.
// path: Location of the secret.
func (a *Secrets) KVv2Write(ctx context.Context, path string, request schema.KVv2WriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/data/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2WriteConfig Configure backend level settings that are applied to every key in the key-value store.
func (a *Secrets) KVv2WriteConfig(ctx context.Context, request schema.KVv2WriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KVv2WriteMetadata Configures settings for the KV store
// path: Location of the secret.
func (a *Secrets) KVv2WriteMetadata(ctx context.Context, path string, request schema.KVv2WriteMetadataRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{secret_mount_path}/metadata/{path}"
	requestPath = strings.Replace(requestPath, "{"+"secret_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("secret")), -1)
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(path), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteConfig
func (a *Secrets) KubernetesDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesDeleteRole
// name: Name of the role
func (a *Secrets) KubernetesDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesListRoles
func (a *Secrets) KubernetesListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadConfig
func (a *Secrets) KubernetesReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesReadRole
// name: Name of the role
func (a *Secrets) KubernetesReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteConfig
func (a *Secrets) KubernetesWriteConfig(ctx context.Context, request schema.KubernetesWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteCredentials
// name: Name of the Vault role
func (a *Secrets) KubernetesWriteCredentials(ctx context.Context, name string, request schema.KubernetesWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// KubernetesWriteRole
// name: Name of the role
func (a *Secrets) KubernetesWriteRole(ctx context.Context, name string, request schema.KubernetesWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{kubernetes_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"kubernetes_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("kubernetes")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) LDAPCheckInLibrary(ctx context.Context, name string, request schema.LDAPCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) LDAPCheckInManageLibrary(ctx context.Context, name string, request schema.LDAPCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (a *Secrets) LDAPCheckOutLibrary(ctx context.Context, name string, request schema.LDAPCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteConfig
func (a *Secrets) LDAPDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteLibrary Delete a library set.
// name: Name of the set.
func (a *Secrets) LDAPDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteRole
// name: Name of the role (lowercase)
func (a *Secrets) LDAPDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPDeleteStaticRole
// name: Name of the role
func (a *Secrets) LDAPDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPListLibraries
func (a *Secrets) LDAPListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPListRoles
func (a *Secrets) LDAPListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPListStaticRoles
func (a *Secrets) LDAPListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadConfig
func (a *Secrets) LDAPReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadCredentials
// name: Name of the dynamic role.
func (a *Secrets) LDAPReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadLibrary Read a library set.
// name: Name of the set.
func (a *Secrets) LDAPReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (a *Secrets) LDAPReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadRole
// name: Name of the role (lowercase)
func (a *Secrets) LDAPReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadStaticCredentials
// name: Name of the static role.
func (a *Secrets) LDAPReadStaticCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-cred/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPReadStaticRole
// name: Name of the role
func (a *Secrets) LDAPReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPRotateRole
// name: Name of the static role
func (a *Secrets) LDAPRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPRotateRoot
func (a *Secrets) LDAPRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteConfig
func (a *Secrets) LDAPWriteConfig(ctx context.Context, request schema.LDAPWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteLibrary Update a library set.
// name: Name of the set.
func (a *Secrets) LDAPWriteLibrary(ctx context.Context, name string, request schema.LDAPWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteRole
// name: Name of the role (lowercase)
func (a *Secrets) LDAPWriteRole(ctx context.Context, name string, request schema.LDAPWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// LDAPWriteStaticRole
// name: Name of the role
func (a *Secrets) LDAPWriteStaticRole(ctx context.Context, name string, request schema.LDAPWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"ldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasDeleteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (a *Secrets) MongoDBAtlasDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasListRoles List the existing roles in this backend
func (a *Secrets) MongoDBAtlasListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasReadConfig Configure the  credentials that are used to manage Database Users.
func (a *Secrets) MongoDBAtlasReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasReadCredentials Generate MongoDB Atlas Programmatic API from a specific Vault role.
// name: Name of the role
func (a *Secrets) MongoDBAtlasReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasReadRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (a *Secrets) MongoDBAtlasReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasWriteConfig Configure the  credentials that are used to manage Database Users.
func (a *Secrets) MongoDBAtlasWriteConfig(ctx context.Context, request schema.MongoDBAtlasWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasWriteCredentials Generate MongoDB Atlas Programmatic API from a specific Vault role.
// name: Name of the role
func (a *Secrets) MongoDBAtlasWriteCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// MongoDBAtlasWriteRole Manage the roles used to generate MongoDB Atlas Programmatic API Keys.
// name: Name of the Roles
func (a *Secrets) MongoDBAtlasWriteRole(ctx context.Context, name string, request schema.MongoDBAtlasWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{mongodbatlas_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"mongodbatlas_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("mongodbatlas")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteAccessConfig
func (a *Secrets) NomadDeleteAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteLeaseConfig Configure the lease parameters for generated tokens
func (a *Secrets) NomadDeleteLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadDeleteRole
// name: Name of the role
func (a *Secrets) NomadDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadListRoles
func (a *Secrets) NomadListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadAccessConfig
func (a *Secrets) NomadReadAccessConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadCredentials
// name: Name of the role
func (a *Secrets) NomadReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadLeaseConfig Configure the lease parameters for generated tokens
func (a *Secrets) NomadReadLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadReadRole
// name: Name of the role
func (a *Secrets) NomadReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadWriteAccessConfig
func (a *Secrets) NomadWriteAccessConfig(ctx context.Context, request schema.NomadWriteAccessConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/access"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadWriteLeaseConfig Configure the lease parameters for generated tokens
func (a *Secrets) NomadWriteLeaseConfig(ctx context.Context, request schema.NomadWriteLeaseConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// NomadWriteRole
// name: Name of the role
func (a *Secrets) NomadWriteRole(ctx context.Context, name string, request schema.NomadWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{nomad_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"nomad_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("nomad")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPCheckInLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) OpenLDAPCheckInLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckInLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPCheckInManageLibrary Check service accounts in to the library.
// name: Name of the set.
func (a *Secrets) OpenLDAPCheckInManageLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckInManageLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/manage/{name}/check-in"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPCheckOutLibrary Check a service account out from the library.
// name: Name of the set
func (a *Secrets) OpenLDAPCheckOutLibrary(ctx context.Context, name string, request schema.OpenLDAPCheckOutLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/check-out"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPDeleteConfig
func (a *Secrets) OpenLDAPDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPDeleteLibrary Delete a library set.
// name: Name of the set.
func (a *Secrets) OpenLDAPDeleteLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPDeleteRole
// name: Name of the role (lowercase)
func (a *Secrets) OpenLDAPDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPDeleteStaticRole
// name: Name of the role
func (a *Secrets) OpenLDAPDeleteStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPListLibraries
func (a *Secrets) OpenLDAPListLibraries(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPListRoles
func (a *Secrets) OpenLDAPListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPListStaticRoles
func (a *Secrets) OpenLDAPListStaticRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadConfig
func (a *Secrets) OpenLDAPReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadCredentials
// name: Name of the dynamic role.
func (a *Secrets) OpenLDAPReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadLibrary Read a library set.
// name: Name of the set.
func (a *Secrets) OpenLDAPReadLibrary(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadLibraryStatus Check the status of the service accounts in a library set.
// name: Name of the set.
func (a *Secrets) OpenLDAPReadLibraryStatus(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}/status"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadRole
// name: Name of the role (lowercase)
func (a *Secrets) OpenLDAPReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadStaticCredentials
// name: Name of the static role.
func (a *Secrets) OpenLDAPReadStaticCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-cred/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPReadStaticRole
// name: Name of the role
func (a *Secrets) OpenLDAPReadStaticRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPRotateRole
// name: Name of the static role
func (a *Secrets) OpenLDAPRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPRotateRoot
func (a *Secrets) OpenLDAPRotateRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/rotate-root"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPWriteConfig
func (a *Secrets) OpenLDAPWriteConfig(ctx context.Context, request schema.OpenLDAPWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPWriteLibrary Update a library set.
// name: Name of the set.
func (a *Secrets) OpenLDAPWriteLibrary(ctx context.Context, name string, request schema.OpenLDAPWriteLibraryRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/library/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPWriteRole
// name: Name of the role (lowercase)
func (a *Secrets) OpenLDAPWriteRole(ctx context.Context, name string, request schema.OpenLDAPWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// OpenLDAPWriteStaticRole
// name: Name of the role
func (a *Secrets) OpenLDAPWriteStaticRole(ctx context.Context, name string, request schema.OpenLDAPWriteStaticRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{openldap_mount_path}/static-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"openldap_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("openldap")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIBundleWrite
func (a *Secrets) PKIBundleWrite(ctx context.Context, request schema.PKIBundleWriteRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/bundle"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIDeleteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (a *Secrets) PKIDeleteKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIDeleteRole
// name: Name of the role
func (a *Secrets) PKIDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIDeleteRoot
func (a *Secrets) PKIDeleteRoot(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (a *Secrets) PKIGenerateRoot(ctx context.Context, exported string, request schema.PKIGenerateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIImportKeys
func (a *Secrets) PKIImportKeys(ctx context.Context, request schema.PKIImportKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys/import"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerIssueRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (a *Secrets) PKIIssuerIssueRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerIssueRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerResignCRLs
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerResignCRLs(ctx context.Context, issuerRef string, request schema.PKIIssuerResignCRLsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/resign-crls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerRevoke
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerRevoke(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignIntermediate
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerSignIntermediate(ctx context.Context, issuerRef string, request schema.PKIIssuerSignIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignRevocationList
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerSignRevocationList(ctx context.Context, issuerRef string, request schema.PKIIssuerSignRevocationListRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-revocation-list"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (a *Secrets) PKIIssuerSignRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerSignRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignSelfIssued
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerSignSelfIssued(ctx context.Context, issuerRef string, request schema.PKIIssuerSignSelfIssuedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignVerbatim
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PKIIssuerSignVerbatim(ctx context.Context, issuerRef string, request schema.PKIIssuerSignVerbatimRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuerSignVerbatimRole
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
// role: The desired role with configuration for this request
func (a *Secrets) PKIIssuerSignVerbatimRole(ctx context.Context, issuerRef string, role string, request schema.PKIIssuerSignVerbatimRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuer/{issuer_ref}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuersGenerateIntermediate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (a *Secrets) PKIIssuersGenerateIntermediate(ctx context.Context, exported string, request schema.PKIIssuersGenerateIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/intermediate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuersGenerateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (a *Secrets) PKIIssuersGenerateRoot(ctx context.Context, exported string, request schema.PKIIssuersGenerateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers/generate/root/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIIssuersList
func (a *Secrets) PKIIssuersList(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIListCerts
func (a *Secrets) PKIListCerts(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIListCertsRevoked
func (a *Secrets) PKIListCertsRevoked(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/certs/revoked"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIListKeys
func (a *Secrets) PKIListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIListRoles
func (a *Secrets) PKIListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadAutoTidyConfig
func (a *Secrets) PKIReadAutoTidyConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCA
func (a *Secrets) PKIReadCA(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCAChain
func (a *Secrets) PKIReadCAChain(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCAPem
func (a *Secrets) PKIReadCAPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ca/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCRL
func (a *Secrets) PKIReadCRL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCRLConfig
func (a *Secrets) PKIReadCRLConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCRLRotate
func (a *Secrets) PKIReadCRLRotate(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCRLRotateDelta
func (a *Secrets) PKIReadCRLRotateDelta(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/crl/rotate-delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCert
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (a *Secrets) PKIReadCert(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCertCAChain
func (a *Secrets) PKIReadCertCAChain(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/ca_chain"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCertRaw
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (a *Secrets) PKIReadCertRaw(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadCertRawPem
// serial: Certificate serial number, in colon- or hyphen-separated octal
func (a *Secrets) PKIReadCertRawPem(ctx context.Context, serial string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert/{serial}/raw/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"serial"+"}", url.PathEscape(serial), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadClusterConfig
func (a *Secrets) PKIReadClusterConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadDeltaCRL
func (a *Secrets) PKIReadDeltaCRL(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/delta-crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadIssuersConfig
func (a *Secrets) PKIReadIssuersConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (a *Secrets) PKIReadKey(ctx context.Context, keyRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadKeysConfig
func (a *Secrets) PKIReadKeysConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadOCSPReq
// req: base-64 encoded ocsp request
func (a *Secrets) PKIReadOCSPReq(ctx context.Context, req string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp/{req}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"req"+"}", url.PathEscape(req), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadRole
// name: Name of the role
func (a *Secrets) PKIReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReadURLConfig
func (a *Secrets) PKIReadURLConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIReplaceRoot
func (a *Secrets) PKIReplaceRoot(ctx context.Context, request schema.PKIReplaceRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/replace"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIRevoke
func (a *Secrets) PKIRevoke(ctx context.Context, request schema.PKIRevokeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIRevokeWithKey
func (a *Secrets) PKIRevokeWithKey(ctx context.Context, request schema.PKIRevokeWithKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/revoke-with-key"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIRootSignIntermediate
func (a *Secrets) PKIRootSignIntermediate(ctx context.Context, request schema.PKIRootSignIntermediateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-intermediate"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIRootSignSelfIssued
func (a *Secrets) PKIRootSignSelfIssued(ctx context.Context, request schema.PKIRootSignSelfIssuedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/sign-self-issued"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIRotateRoot
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (a *Secrets) PKIRotateRoot(ctx context.Context, exported string, request schema.PKIRotateRootRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/root/rotate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKISignRole
// role: The desired role with configuration for this request
func (a *Secrets) PKISignRole(ctx context.Context, role string, request schema.PKISignRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKISignVerbatim
func (a *Secrets) PKISignVerbatim(ctx context.Context, request schema.PKISignVerbatimRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKISignVerbatimRole
// role: The desired role with configuration for this request
func (a *Secrets) PKISignVerbatimRole(ctx context.Context, role string, request schema.PKISignVerbatimRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/sign-verbatim/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKITidy
func (a *Secrets) PKITidy(ctx context.Context, request schema.PKITidyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKITidyCancel
func (a *Secrets) PKITidyCancel(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-cancel"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKITidyStatus
func (a *Secrets) PKITidyStatus(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/tidy-status"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteAutoTidyConfig
func (a *Secrets) PKIWriteAutoTidyConfig(ctx context.Context, request schema.PKIWriteAutoTidyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/auto-tidy"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteCAConfig
func (a *Secrets) PKIWriteCAConfig(ctx context.Context, request schema.PKIWriteCAConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteCRLConfig
func (a *Secrets) PKIWriteCRLConfig(ctx context.Context, request schema.PKIWriteCRLConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/crl"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteCerts
func (a *Secrets) PKIWriteCerts(ctx context.Context, request schema.PKIWriteCertsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/cert"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteClusterConfig
func (a *Secrets) PKIWriteClusterConfig(ctx context.Context, request schema.PKIWriteClusterConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/cluster"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteIntermediateCrossSign
func (a *Secrets) PKIWriteIntermediateCrossSign(ctx context.Context, request schema.PKIWriteIntermediateCrossSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/cross-sign"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteIntermediateGenerate
// exported: Must be \&quot;internal\&quot;, \&quot;exported\&quot; or \&quot;kms\&quot;. If set to \&quot;exported\&quot;, the generated private key will be returned. This is your *only* chance to retrieve the private key!
func (a *Secrets) PKIWriteIntermediateGenerate(ctx context.Context, exported string, request schema.PKIWriteIntermediateGenerateRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/generate/{exported}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"exported"+"}", url.PathEscape(exported), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteIntermediateSetSigned
func (a *Secrets) PKIWriteIntermediateSetSigned(ctx context.Context, request schema.PKIWriteIntermediateSetSignedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/intermediate/set-signed"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteInternalExported
func (a *Secrets) PKIWriteInternalExported(ctx context.Context, request schema.PKIWriteInternalExportedRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/internal|exported"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteIssueRole
// role: The desired role with configuration for this request
func (a *Secrets) PKIWriteIssueRole(ctx context.Context, role string, request schema.PKIWriteIssueRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteIssuersConfig
func (a *Secrets) PKIWriteIssuersConfig(ctx context.Context, request schema.PKIWriteIssuersConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/issuers"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteKMS
func (a *Secrets) PKIWriteKMS(ctx context.Context, request schema.PKIWriteKMSRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/kms"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteKey
// keyRef: Reference to key; either \&quot;default\&quot; for the configured default key, an identifier of a key, or the name assigned to the key.
func (a *Secrets) PKIWriteKey(ctx context.Context, keyRef string, request schema.PKIWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/key/{key_ref}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_ref"+"}", url.PathEscape(keyRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteKeysConfig
func (a *Secrets) PKIWriteKeysConfig(ctx context.Context, request schema.PKIWriteKeysConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteOCSP
func (a *Secrets) PKIWriteOCSP(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/ocsp"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteRole
// name: Name of the role
func (a *Secrets) PKIWriteRole(ctx context.Context, name string, request schema.PKIWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PKIWriteURLConfig
func (a *Secrets) PKIWriteURLConfig(ctx context.Context, request schema.PKIWriteURLConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/config/urls"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PkiDeleteIssuerRefDerPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiDeleteJson
func (a *Secrets) PkiDeleteJson(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadDelta
func (a *Secrets) PkiReadDelta(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//delta"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadDeltaPem
func (a *Secrets) PkiReadDeltaPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadDer
func (a *Secrets) PkiReadDer(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//der"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuerRefCrlPemDerDeltaPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PkiReadIssuerRefCrlPemDerDeltaPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/crl/pem|/der|/delta/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PkiReadIssuerRefDerPem(ctx context.Context, issuerRef string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadJson
func (a *Secrets) PkiReadJson(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiReadPem
func (a *Secrets) PkiReadPem(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiWriteIssuerRefDerPem
// issuerRef: Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer.
func (a *Secrets) PkiWriteIssuerRefDerPem(ctx context.Context, issuerRef string, request schema.PkiWriteIssuerRefDerPemRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}/{issuer_ref}/der|/pem"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)
	requestPath = strings.Replace(requestPath, "{"+"issuer_ref"+"}", url.PathEscape(issuerRef), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// PkiWriteJson
func (a *Secrets) PkiWriteJson(ctx context.Context, request schema.PkiWriteJsonRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{pki_mount_path}//json"
	requestPath = strings.Replace(requestPath, "{"+"pki_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("pki")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQDeleteRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (a *Secrets) RabbitMQDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQListRoles Manage the roles that can be created with this backend.
func (a *Secrets) RabbitMQListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQReadCredentials Request RabbitMQ credentials for a certain role.
// name: Name of the role.
func (a *Secrets) RabbitMQReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQReadLeaseConfig Configure the lease parameters for generated credentials
func (a *Secrets) RabbitMQReadLeaseConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQReadRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (a *Secrets) RabbitMQReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQWriteConnectionConfig Configure the connection URI, username, and password to talk to RabbitMQ management HTTP API.
func (a *Secrets) RabbitMQWriteConnectionConfig(ctx context.Context, request schema.RabbitMQWriteConnectionConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/connection"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQWriteLeaseConfig Configure the lease parameters for generated credentials
func (a *Secrets) RabbitMQWriteLeaseConfig(ctx context.Context, request schema.RabbitMQWriteLeaseConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/config/lease"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// RabbitMQWriteRole Manage the roles that can be created with this backend.
// name: Name of the role.
func (a *Secrets) RabbitMQWriteRole(ctx context.Context, name string, request schema.RabbitMQWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{rabbitmq_mount_path}/roles/{name}"
	requestPath = strings.Replace(requestPath, "{"+"rabbitmq_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("rabbitmq")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHDeleteCAConfig Set the SSH private key used for signing certificates.
func (a *Secrets) SSHDeleteCAConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHDeleteKeys Register a shared private key with Vault.
// keyName: [Required] Name of the key
func (a *Secrets) SSHDeleteKeys(ctx context.Context, keyName string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/keys/{key_name}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_name"+"}", url.PathEscape(keyName), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHDeleteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (a *Secrets) SSHDeleteRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHDeleteZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (a *Secrets) SSHDeleteZeroAddressConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHListRoles Manage the 'roles' that can be created with this backend.
func (a *Secrets) SSHListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHLookup List all the roles associated with the given IP address.
func (a *Secrets) SSHLookup(ctx context.Context, request schema.SSHLookupRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/lookup"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHReadCAConfig Set the SSH private key used for signing certificates.
func (a *Secrets) SSHReadCAConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHReadPublicKey Retrieve the public key.
func (a *Secrets) SSHReadPublicKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/public_key"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHReadRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (a *Secrets) SSHReadRole(ctx context.Context, role string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHReadZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (a *Secrets) SSHReadZeroAddressConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHSign Request signing an SSH key using a certain role with the provided details.
// role: The desired role with configuration for this request.
func (a *Secrets) SSHSign(ctx context.Context, role string, request schema.SSHSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/sign/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHVerify Validate the OTP provided by Vault SSH Agent.
func (a *Secrets) SSHVerify(ctx context.Context, request schema.SSHVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/verify"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteCAConfig Set the SSH private key used for signing certificates.
func (a *Secrets) SSHWriteCAConfig(ctx context.Context, request schema.SSHWriteCAConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/ca"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteCredentials Creates a credential for establishing SSH connection with the remote host.
// role: [Required] Name of the role
func (a *Secrets) SSHWriteCredentials(ctx context.Context, role string, request schema.SSHWriteCredentialsRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/creds/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteIssue
// role: The desired role with configuration for this request.
func (a *Secrets) SSHWriteIssue(ctx context.Context, role string, request schema.SSHWriteIssueRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/issue/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteKeys Register a shared private key with Vault.
// keyName: [Required] Name of the key
func (a *Secrets) SSHWriteKeys(ctx context.Context, keyName string, request schema.SSHWriteKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/keys/{key_name}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"key_name"+"}", url.PathEscape(keyName), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteRole Manage the 'roles' that can be created with this backend.
// role: [Required for all types] Name of the role being created.
func (a *Secrets) SSHWriteRole(ctx context.Context, role string, request schema.SSHWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/roles/{role}"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)
	requestPath = strings.Replace(requestPath, "{"+"role"+"}", url.PathEscape(role), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// SSHWriteZeroAddressConfig Assign zero address as default CIDR block for select roles.
func (a *Secrets) SSHWriteZeroAddressConfig(ctx context.Context, request schema.SSHWriteZeroAddressConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{ssh_mount_path}/config/zeroaddress"
	requestPath = strings.Replace(requestPath, "{"+"ssh_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("ssh")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPDeleteKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (a *Secrets) TOTPDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPListKeys Manage the keys that can be created with this backend.
func (a *Secrets) TOTPListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPReadCode Request time-based one-time use password or validate a password for a certain key .
// name: Name of the key.
func (a *Secrets) TOTPReadCode(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPReadKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (a *Secrets) TOTPReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPWriteCode Request time-based one-time use password or validate a password for a certain key .
// name: Name of the key.
func (a *Secrets) TOTPWriteCode(ctx context.Context, name string, request schema.TOTPWriteCodeRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/code/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TOTPWriteKey Manage the keys that can be created with this backend.
// name: Name of the key.
func (a *Secrets) TOTPWriteKey(ctx context.Context, name string, request schema.TOTPWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{totp_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"totp_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("totp")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformDeleteConfig
func (a *Secrets) TerraformDeleteConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformDeleteRole
// name: Name of the role
func (a *Secrets) TerraformDeleteRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformListRoles
func (a *Secrets) TerraformListRoles(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformReadConfig
func (a *Secrets) TerraformReadConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformReadCredentials Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
// name: Name of the role
func (a *Secrets) TerraformReadCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformReadRole
// name: Name of the role
func (a *Secrets) TerraformReadRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformRotateRole
// name: Name of the team or organization role
func (a *Secrets) TerraformRotateRole(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/rotate-role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformWriteConfig
func (a *Secrets) TerraformWriteConfig(ctx context.Context, request schema.TerraformWriteConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/config"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformWriteCredentials Generate a Terraform Cloud or Enterprise API token from a specific Vault role.
// name: Name of the role
func (a *Secrets) TerraformWriteCredentials(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/creds/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TerraformWriteRole
// name: Name of the role
func (a *Secrets) TerraformWriteRole(ctx context.Context, name string, request schema.TerraformWriteRoleRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{terraform_mount_path}/role/{name}"
	requestPath = strings.Replace(requestPath, "{"+"terraform_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("terraform")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitBackup Backup the named key
// name: Name of the key
func (a *Secrets) TransitBackup(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/backup/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitDecrypt Decrypt a ciphertext value using a named key
// name: Name of the key
func (a *Secrets) TransitDecrypt(ctx context.Context, name string, request schema.TransitDecryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/decrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitDeleteKey Managed named encryption keys
// name: Name of the key
func (a *Secrets) TransitDeleteKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodDelete,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitEncrypt Encrypt a plaintext value or a batch of plaintext blocks using a named key
// name: Name of the key
func (a *Secrets) TransitEncrypt(ctx context.Context, name string, request schema.TransitEncryptRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/encrypt/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitExport Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
func (a *Secrets) TransitExport(ctx context.Context, name string, type_ string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitExportVersion Export named encryption or signing key
// name: Name of the key
// type_: Type of key to export (encryption-key, signing-key, hmac-key)
// version: Version of the key
func (a *Secrets) TransitExportVersion(ctx context.Context, name string, type_ string, version string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/export/{type}/{name}/{version}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"type"+"}", url.PathEscape(type_), -1)
	requestPath = strings.Replace(requestPath, "{"+"version"+"}", url.PathEscape(version), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateDataKey Generate a data key
// name: The backend key used for encrypting the data key
// plaintext: \&quot;plaintext\&quot; will return the key in both plaintext and ciphertext; \&quot;wrapped\&quot; will return the ciphertext only.
func (a *Secrets) TransitGenerateDataKey(ctx context.Context, name string, plaintext string, request schema.TransitGenerateDataKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/datakey/{plaintext}/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"plaintext"+"}", url.PathEscape(plaintext), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateHMAC Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
func (a *Secrets) TransitGenerateHMAC(ctx context.Context, name string, request schema.TransitGenerateHMACRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateHMACWithAlgorithm Generate an HMAC for input data using the named key
// name: The key to use for the HMAC function
// urlalgorithm: Algorithm to use (POST URL parameter)
func (a *Secrets) TransitGenerateHMACWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitGenerateHMACWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hmac/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandom Generate random bytes
func (a *Secrets) TransitGenerateRandom(ctx context.Context, request schema.TransitGenerateRandomRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandomSource Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
func (a *Secrets) TransitGenerateRandomSource(ctx context.Context, source string, request schema.TransitGenerateRandomSourceRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitGenerateRandomSourceBytes Generate random bytes
// source: Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;.
// urlbytes: The number of bytes to generate (POST URL parameter)
func (a *Secrets) TransitGenerateRandomSourceBytes(ctx context.Context, source string, urlbytes string, request schema.TransitGenerateRandomSourceBytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{source}/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"source"+"}", url.PathEscape(source), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitHash Generate a hash sum for input data
func (a *Secrets) TransitHash(ctx context.Context, request schema.TransitHashRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitHashWithAlgorithm Generate a hash sum for input data
// urlalgorithm: Algorithm to use (POST URL parameter)
func (a *Secrets) TransitHashWithAlgorithm(ctx context.Context, urlalgorithm string, request schema.TransitHashWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/hash/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitImportKey Imports an externally-generated key into a new transit key
// name: The name of the key
func (a *Secrets) TransitImportKey(ctx context.Context, name string, request schema.TransitImportKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitImportKeyVersion Imports an externally-generated key into an existing imported key
// name: The name of the key
func (a *Secrets) TransitImportKeyVersion(ctx context.Context, name string, request schema.TransitImportKeyVersionRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/import_version"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitListKeys Managed named encryption keys
func (a *Secrets) TransitListKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)
	requestQueryParameters.Set("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadCacheConfig Returns the size of the active cache
func (a *Secrets) TransitReadCacheConfig(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadConfigKeys Configuration common across all keys
func (a *Secrets) TransitReadConfigKeys(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadKey Managed named encryption keys
// name: Name of the key
func (a *Secrets) TransitReadKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitReadWrappingKey Returns the public key to use for wrapping imported keys
func (a *Secrets) TransitReadWrappingKey(ctx context.Context, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/wrapping_key"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodGet,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRestore Restore the named key
func (a *Secrets) TransitRestore(ctx context.Context, request schema.TransitRestoreRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRestoreKey Restore the named key
// name: If set, this will be the name of the restored key.
func (a *Secrets) TransitRestoreKey(ctx context.Context, name string, request schema.TransitRestoreKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/restore/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRewrap Rewrap ciphertext
// name: Name of the key
func (a *Secrets) TransitRewrap(ctx context.Context, name string, request schema.TransitRewrapRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/rewrap/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitRotateKey Rotate named encryption key
// name: Name of the key
func (a *Secrets) TransitRotateKey(ctx context.Context, name string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/rotate"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitSign Generate a signature for input data using the named key
// name: The key to use
func (a *Secrets) TransitSign(ctx context.Context, name string, request schema.TransitSignRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitSignWithAlgorithm Generate a signature for input data using the named key
// name: The key to use
// urlalgorithm: Hash algorithm to use (POST URL parameter)
func (a *Secrets) TransitSignWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitSignWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/sign/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitTrimKey Trim key versions of a named key
// name: Name of the key
func (a *Secrets) TransitTrimKey(ctx context.Context, name string, request schema.TransitTrimKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/trim"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitVerify Verify a signature or HMAC for input data created using the named key
// name: The key to use
func (a *Secrets) TransitVerify(ctx context.Context, name string, request schema.TransitVerifyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitVerifyWithAlgorithm Verify a signature or HMAC for input data created using the named key
// name: The key to use
// urlalgorithm: Hash algorithm to use (POST URL parameter)
func (a *Secrets) TransitVerifyWithAlgorithm(ctx context.Context, name string, urlalgorithm string, request schema.TransitVerifyWithAlgorithmRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/verify/{name}/{urlalgorithm}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlalgorithm"+"}", url.PathEscape(urlalgorithm), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitWriteCacheConfig Configures a new cache of the specified size
func (a *Secrets) TransitWriteCacheConfig(ctx context.Context, request schema.TransitWriteCacheConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/cache-config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitWriteConfigKeys Configuration common across all keys
func (a *Secrets) TransitWriteConfigKeys(ctx context.Context, request schema.TransitWriteConfigKeysRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/config/keys"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitWriteKey Managed named encryption keys
// name: Name of the key
func (a *Secrets) TransitWriteKey(ctx context.Context, name string, request schema.TransitWriteKeyRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitWriteKeyConfig Configure a named encryption key
// name: Name of the key
func (a *Secrets) TransitWriteKeyConfig(ctx context.Context, name string, request schema.TransitWriteKeyConfigRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/keys/{name}/config"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"name"+"}", url.PathEscape(name), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}

// TransitWriteRandomUrlbytes Generate random bytes
// urlbytes: The number of bytes to generate (POST URL parameter)
func (a *Secrets) TransitWriteRandomUrlbytes(ctx context.Context, urlbytes string, request schema.TransitWriteRandomUrlbytesRequest, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestPath := "/v1/{transit_mount_path}/random/{urlbytes}"
	requestPath = strings.Replace(requestPath, "{"+"transit_mount_path"+"}", url.PathEscape(requestModifiers.mountPathOr("transit")), -1)
	requestPath = strings.Replace(requestPath, "{"+"urlbytes"+"}", url.PathEscape(urlbytes), -1)

	requestQueryParameters := make(url.Values)

	return sendStructuredRequestParseResponse[map[string]interface{}](
		ctx,
		a.client,
		http.MethodPost,
		requestPath,
		request,
		requestQueryParameters,
		requestModifiers,
	)
}
