/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"encoding/json"
)

// CfRolesRequest struct for CfRolesRequest
type CfRolesRequest struct {
	// Require that the client certificate presented has at least one of these app IDs.
	BoundApplicationIds []string `json:"bound_application_ids,omitempty"`
	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs,omitempty"`
	// Require that the client certificate presented has at least one of these instance IDs.
	BoundInstanceIds []string `json:"bound_instance_ids,omitempty"`
	// Require that the client certificate presented has at least one of these org IDs.
	BoundOrganizationIds []string `json:"bound_organization_ids,omitempty"`
	// Require that the client certificate presented has at least one of these space IDs.
	BoundSpaceIds []string `json:"bound_space_ids,omitempty"`
	// If set to true, disables the default behavior that logging in must be performed from an acceptable IP address described by the certificate presented.
	DisableIpMatching *bool `json:"disable_ip_matching,omitempty"`
	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl *int32 `json:"max_ttl,omitempty"`
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period *int32 `json:"period,omitempty"`
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`
	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl *int32 `json:"token_explicit_max_ttl,omitempty"`
	// The maximum lifetime of the generated token
	TokenMaxTtl *int32 `json:"token_max_ttl,omitempty"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"token_no_default_policy,omitempty"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *int32 `json:"token_num_uses,omitempty"`
	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod *int32 `json:"token_period,omitempty"`
	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`
	// The initial ttl of the token to generate
	TokenTtl *int32 `json:"token_ttl,omitempty"`
	// The type of token to generate, service or batch
	TokenType *string `json:"token_type,omitempty"`
	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewCfRolesRequest instantiates a new CfRolesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfRolesRequest() *CfRolesRequest {
	this := CfRolesRequest{}
	var disableIpMatching bool = false
	this.DisableIpMatching = &disableIpMatching
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// NewCfRolesRequestWithDefaults instantiates a new CfRolesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfRolesRequestWithDefaults() *CfRolesRequest {
	this := CfRolesRequest{}
	var disableIpMatching bool = false
	this.DisableIpMatching = &disableIpMatching
	var tokenType string = "default-service"
	this.TokenType = &tokenType
	return &this
}

// GetBoundApplicationIds returns the BoundApplicationIds field value if set, zero value otherwise.
func (o *CfRolesRequest) GetBoundApplicationIds() []string {
	if o == nil || o.BoundApplicationIds == nil {
		var ret []string
		return ret
	}
	return o.BoundApplicationIds
}

// GetBoundApplicationIdsOk returns a tuple with the BoundApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetBoundApplicationIdsOk() ([]string, bool) {
	if o == nil || o.BoundApplicationIds == nil {
		return nil, false
	}
	return o.BoundApplicationIds, true
}

// HasBoundApplicationIds returns a boolean if a field has been set.
func (o *CfRolesRequest) HasBoundApplicationIds() bool {
	if o != nil && o.BoundApplicationIds != nil {
		return true
	}

	return false
}

// SetBoundApplicationIds gets a reference to the given []string and assigns it to the BoundApplicationIds field.
func (o *CfRolesRequest) SetBoundApplicationIds(v []string) {
	o.BoundApplicationIds = v
}

// GetBoundCidrs returns the BoundCidrs field value if set, zero value otherwise.
// Deprecated
func (o *CfRolesRequest) GetBoundCidrs() []string {
	if o == nil || o.BoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.BoundCidrs
}

// GetBoundCidrsOk returns a tuple with the BoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CfRolesRequest) GetBoundCidrsOk() ([]string, bool) {
	if o == nil || o.BoundCidrs == nil {
		return nil, false
	}
	return o.BoundCidrs, true
}

// HasBoundCidrs returns a boolean if a field has been set.
func (o *CfRolesRequest) HasBoundCidrs() bool {
	if o != nil && o.BoundCidrs != nil {
		return true
	}

	return false
}

// SetBoundCidrs gets a reference to the given []string and assigns it to the BoundCidrs field.
// Deprecated
func (o *CfRolesRequest) SetBoundCidrs(v []string) {
	o.BoundCidrs = v
}

// GetBoundInstanceIds returns the BoundInstanceIds field value if set, zero value otherwise.
func (o *CfRolesRequest) GetBoundInstanceIds() []string {
	if o == nil || o.BoundInstanceIds == nil {
		var ret []string
		return ret
	}
	return o.BoundInstanceIds
}

// GetBoundInstanceIdsOk returns a tuple with the BoundInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetBoundInstanceIdsOk() ([]string, bool) {
	if o == nil || o.BoundInstanceIds == nil {
		return nil, false
	}
	return o.BoundInstanceIds, true
}

// HasBoundInstanceIds returns a boolean if a field has been set.
func (o *CfRolesRequest) HasBoundInstanceIds() bool {
	if o != nil && o.BoundInstanceIds != nil {
		return true
	}

	return false
}

// SetBoundInstanceIds gets a reference to the given []string and assigns it to the BoundInstanceIds field.
func (o *CfRolesRequest) SetBoundInstanceIds(v []string) {
	o.BoundInstanceIds = v
}

// GetBoundOrganizationIds returns the BoundOrganizationIds field value if set, zero value otherwise.
func (o *CfRolesRequest) GetBoundOrganizationIds() []string {
	if o == nil || o.BoundOrganizationIds == nil {
		var ret []string
		return ret
	}
	return o.BoundOrganizationIds
}

// GetBoundOrganizationIdsOk returns a tuple with the BoundOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetBoundOrganizationIdsOk() ([]string, bool) {
	if o == nil || o.BoundOrganizationIds == nil {
		return nil, false
	}
	return o.BoundOrganizationIds, true
}

// HasBoundOrganizationIds returns a boolean if a field has been set.
func (o *CfRolesRequest) HasBoundOrganizationIds() bool {
	if o != nil && o.BoundOrganizationIds != nil {
		return true
	}

	return false
}

// SetBoundOrganizationIds gets a reference to the given []string and assigns it to the BoundOrganizationIds field.
func (o *CfRolesRequest) SetBoundOrganizationIds(v []string) {
	o.BoundOrganizationIds = v
}

// GetBoundSpaceIds returns the BoundSpaceIds field value if set, zero value otherwise.
func (o *CfRolesRequest) GetBoundSpaceIds() []string {
	if o == nil || o.BoundSpaceIds == nil {
		var ret []string
		return ret
	}
	return o.BoundSpaceIds
}

// GetBoundSpaceIdsOk returns a tuple with the BoundSpaceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetBoundSpaceIdsOk() ([]string, bool) {
	if o == nil || o.BoundSpaceIds == nil {
		return nil, false
	}
	return o.BoundSpaceIds, true
}

// HasBoundSpaceIds returns a boolean if a field has been set.
func (o *CfRolesRequest) HasBoundSpaceIds() bool {
	if o != nil && o.BoundSpaceIds != nil {
		return true
	}

	return false
}

// SetBoundSpaceIds gets a reference to the given []string and assigns it to the BoundSpaceIds field.
func (o *CfRolesRequest) SetBoundSpaceIds(v []string) {
	o.BoundSpaceIds = v
}

// GetDisableIpMatching returns the DisableIpMatching field value if set, zero value otherwise.
func (o *CfRolesRequest) GetDisableIpMatching() bool {
	if o == nil || o.DisableIpMatching == nil {
		var ret bool
		return ret
	}
	return *o.DisableIpMatching
}

// GetDisableIpMatchingOk returns a tuple with the DisableIpMatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetDisableIpMatchingOk() (*bool, bool) {
	if o == nil || o.DisableIpMatching == nil {
		return nil, false
	}
	return o.DisableIpMatching, true
}

// HasDisableIpMatching returns a boolean if a field has been set.
func (o *CfRolesRequest) HasDisableIpMatching() bool {
	if o != nil && o.DisableIpMatching != nil {
		return true
	}

	return false
}

// SetDisableIpMatching gets a reference to the given bool and assigns it to the DisableIpMatching field.
func (o *CfRolesRequest) SetDisableIpMatching(v bool) {
	o.DisableIpMatching = &v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
// Deprecated
func (o *CfRolesRequest) GetMaxTtl() int32 {
	if o == nil || o.MaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CfRolesRequest) GetMaxTtlOk() (*int32, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *CfRolesRequest) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int32 and assigns it to the MaxTtl field.
// Deprecated
func (o *CfRolesRequest) SetMaxTtl(v int32) {
	o.MaxTtl = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
// Deprecated
func (o *CfRolesRequest) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CfRolesRequest) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CfRolesRequest) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
// Deprecated
func (o *CfRolesRequest) SetPeriod(v int32) {
	o.Period = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
// Deprecated
func (o *CfRolesRequest) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CfRolesRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CfRolesRequest) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
// Deprecated
func (o *CfRolesRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetTokenBoundCidrs returns the TokenBoundCidrs field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenBoundCidrs() []string {
	if o == nil || o.TokenBoundCidrs == nil {
		var ret []string
		return ret
	}
	return o.TokenBoundCidrs
}

// GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenBoundCidrsOk() ([]string, bool) {
	if o == nil || o.TokenBoundCidrs == nil {
		return nil, false
	}
	return o.TokenBoundCidrs, true
}

// HasTokenBoundCidrs returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenBoundCidrs() bool {
	if o != nil && o.TokenBoundCidrs != nil {
		return true
	}

	return false
}

// SetTokenBoundCidrs gets a reference to the given []string and assigns it to the TokenBoundCidrs field.
func (o *CfRolesRequest) SetTokenBoundCidrs(v []string) {
	o.TokenBoundCidrs = v
}

// GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenExplicitMaxTtl() int32 {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenExplicitMaxTtl
}

// GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenExplicitMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenExplicitMaxTtl == nil {
		return nil, false
	}
	return o.TokenExplicitMaxTtl, true
}

// HasTokenExplicitMaxTtl returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenExplicitMaxTtl() bool {
	if o != nil && o.TokenExplicitMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenExplicitMaxTtl gets a reference to the given int32 and assigns it to the TokenExplicitMaxTtl field.
func (o *CfRolesRequest) SetTokenExplicitMaxTtl(v int32) {
	o.TokenExplicitMaxTtl = &v
}

// GetTokenMaxTtl returns the TokenMaxTtl field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenMaxTtl() int32 {
	if o == nil || o.TokenMaxTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenMaxTtl
}

// GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenMaxTtlOk() (*int32, bool) {
	if o == nil || o.TokenMaxTtl == nil {
		return nil, false
	}
	return o.TokenMaxTtl, true
}

// HasTokenMaxTtl returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenMaxTtl() bool {
	if o != nil && o.TokenMaxTtl != nil {
		return true
	}

	return false
}

// SetTokenMaxTtl gets a reference to the given int32 and assigns it to the TokenMaxTtl field.
func (o *CfRolesRequest) SetTokenMaxTtl(v int32) {
	o.TokenMaxTtl = &v
}

// GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenNoDefaultPolicy() bool {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		var ret bool
		return ret
	}
	return *o.TokenNoDefaultPolicy
}

// GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenNoDefaultPolicyOk() (*bool, bool) {
	if o == nil || o.TokenNoDefaultPolicy == nil {
		return nil, false
	}
	return o.TokenNoDefaultPolicy, true
}

// HasTokenNoDefaultPolicy returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenNoDefaultPolicy() bool {
	if o != nil && o.TokenNoDefaultPolicy != nil {
		return true
	}

	return false
}

// SetTokenNoDefaultPolicy gets a reference to the given bool and assigns it to the TokenNoDefaultPolicy field.
func (o *CfRolesRequest) SetTokenNoDefaultPolicy(v bool) {
	o.TokenNoDefaultPolicy = &v
}

// GetTokenNumUses returns the TokenNumUses field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenNumUses() int32 {
	if o == nil || o.TokenNumUses == nil {
		var ret int32
		return ret
	}
	return *o.TokenNumUses
}

// GetTokenNumUsesOk returns a tuple with the TokenNumUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenNumUsesOk() (*int32, bool) {
	if o == nil || o.TokenNumUses == nil {
		return nil, false
	}
	return o.TokenNumUses, true
}

// HasTokenNumUses returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenNumUses() bool {
	if o != nil && o.TokenNumUses != nil {
		return true
	}

	return false
}

// SetTokenNumUses gets a reference to the given int32 and assigns it to the TokenNumUses field.
func (o *CfRolesRequest) SetTokenNumUses(v int32) {
	o.TokenNumUses = &v
}

// GetTokenPeriod returns the TokenPeriod field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenPeriod() int32 {
	if o == nil || o.TokenPeriod == nil {
		var ret int32
		return ret
	}
	return *o.TokenPeriod
}

// GetTokenPeriodOk returns a tuple with the TokenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenPeriodOk() (*int32, bool) {
	if o == nil || o.TokenPeriod == nil {
		return nil, false
	}
	return o.TokenPeriod, true
}

// HasTokenPeriod returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenPeriod() bool {
	if o != nil && o.TokenPeriod != nil {
		return true
	}

	return false
}

// SetTokenPeriod gets a reference to the given int32 and assigns it to the TokenPeriod field.
func (o *CfRolesRequest) SetTokenPeriod(v int32) {
	o.TokenPeriod = &v
}

// GetTokenPolicies returns the TokenPolicies field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenPolicies() []string {
	if o == nil || o.TokenPolicies == nil {
		var ret []string
		return ret
	}
	return o.TokenPolicies
}

// GetTokenPoliciesOk returns a tuple with the TokenPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenPoliciesOk() ([]string, bool) {
	if o == nil || o.TokenPolicies == nil {
		return nil, false
	}
	return o.TokenPolicies, true
}

// HasTokenPolicies returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenPolicies() bool {
	if o != nil && o.TokenPolicies != nil {
		return true
	}

	return false
}

// SetTokenPolicies gets a reference to the given []string and assigns it to the TokenPolicies field.
func (o *CfRolesRequest) SetTokenPolicies(v []string) {
	o.TokenPolicies = v
}

// GetTokenTtl returns the TokenTtl field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenTtl() int32 {
	if o == nil || o.TokenTtl == nil {
		var ret int32
		return ret
	}
	return *o.TokenTtl
}

// GetTokenTtlOk returns a tuple with the TokenTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenTtlOk() (*int32, bool) {
	if o == nil || o.TokenTtl == nil {
		return nil, false
	}
	return o.TokenTtl, true
}

// HasTokenTtl returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenTtl() bool {
	if o != nil && o.TokenTtl != nil {
		return true
	}

	return false
}

// SetTokenTtl gets a reference to the given int32 and assigns it to the TokenTtl field.
func (o *CfRolesRequest) SetTokenTtl(v int32) {
	o.TokenTtl = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *CfRolesRequest) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfRolesRequest) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *CfRolesRequest) SetTokenType(v string) {
	o.TokenType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
// Deprecated
func (o *CfRolesRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CfRolesRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CfRolesRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
// Deprecated
func (o *CfRolesRequest) SetTtl(v int32) {
	o.Ttl = &v
}

func (o CfRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoundApplicationIds != nil {
		toSerialize["bound_application_ids"] = o.BoundApplicationIds
	}
	if o.BoundCidrs != nil {
		toSerialize["bound_cidrs"] = o.BoundCidrs
	}
	if o.BoundInstanceIds != nil {
		toSerialize["bound_instance_ids"] = o.BoundInstanceIds
	}
	if o.BoundOrganizationIds != nil {
		toSerialize["bound_organization_ids"] = o.BoundOrganizationIds
	}
	if o.BoundSpaceIds != nil {
		toSerialize["bound_space_ids"] = o.BoundSpaceIds
	}
	if o.DisableIpMatching != nil {
		toSerialize["disable_ip_matching"] = o.DisableIpMatching
	}
	if o.MaxTtl != nil {
		toSerialize["max_ttl"] = o.MaxTtl
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.TokenBoundCidrs != nil {
		toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	}
	if o.TokenExplicitMaxTtl != nil {
		toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	}
	if o.TokenMaxTtl != nil {
		toSerialize["token_max_ttl"] = o.TokenMaxTtl
	}
	if o.TokenNoDefaultPolicy != nil {
		toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	}
	if o.TokenNumUses != nil {
		toSerialize["token_num_uses"] = o.TokenNumUses
	}
	if o.TokenPeriod != nil {
		toSerialize["token_period"] = o.TokenPeriod
	}
	if o.TokenPolicies != nil {
		toSerialize["token_policies"] = o.TokenPolicies
	}
	if o.TokenTtl != nil {
		toSerialize["token_ttl"] = o.TokenTtl
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableCfRolesRequest struct {
	value *CfRolesRequest
	isSet bool
}

func (v NullableCfRolesRequest) Get() *CfRolesRequest {
	return v.value
}

func (v *NullableCfRolesRequest) Set(val *CfRolesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCfRolesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCfRolesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfRolesRequest(val *CfRolesRequest) *NullableCfRolesRequest {
	return &NullableCfRolesRequest{value: val, isSet: true}
}

func (v NullableCfRolesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfRolesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


