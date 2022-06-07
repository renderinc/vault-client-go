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

// IdentityEntityMergeRequest struct for IdentityEntityMergeRequest
type IdentityEntityMergeRequest struct {
	// Setting this will follow the 'mine' strategy for merging MFA secrets. If there are secrets of the same type both in entities that are merged from and in entity into which all others are getting merged, secrets in the destination will be unaltered. If not set, this API will throw an error containing all the conflicts.
	Force *bool `json:"force,omitempty"`
	// Entity IDs which needs to get merged
	FromEntityIds []string `json:"from_entity_ids,omitempty"`
	// Entity ID into which all the other entities need to get merged
	ToEntityId *string `json:"to_entity_id,omitempty"`
}

// NewIdentityEntityMergeRequest instantiates a new IdentityEntityMergeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEntityMergeRequest() *IdentityEntityMergeRequest {
	this := IdentityEntityMergeRequest{}
	return &this
}

// NewIdentityEntityMergeRequestWithDefaults instantiates a new IdentityEntityMergeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEntityMergeRequestWithDefaults() *IdentityEntityMergeRequest {
	this := IdentityEntityMergeRequest{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *IdentityEntityMergeRequest) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityMergeRequest) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *IdentityEntityMergeRequest) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *IdentityEntityMergeRequest) SetForce(v bool) {
	o.Force = &v
}

// GetFromEntityIds returns the FromEntityIds field value if set, zero value otherwise.
func (o *IdentityEntityMergeRequest) GetFromEntityIds() []string {
	if o == nil || o.FromEntityIds == nil {
		var ret []string
		return ret
	}
	return o.FromEntityIds
}

// GetFromEntityIdsOk returns a tuple with the FromEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityMergeRequest) GetFromEntityIdsOk() ([]string, bool) {
	if o == nil || o.FromEntityIds == nil {
		return nil, false
	}
	return o.FromEntityIds, true
}

// HasFromEntityIds returns a boolean if a field has been set.
func (o *IdentityEntityMergeRequest) HasFromEntityIds() bool {
	if o != nil && o.FromEntityIds != nil {
		return true
	}

	return false
}

// SetFromEntityIds gets a reference to the given []string and assigns it to the FromEntityIds field.
func (o *IdentityEntityMergeRequest) SetFromEntityIds(v []string) {
	o.FromEntityIds = v
}

// GetToEntityId returns the ToEntityId field value if set, zero value otherwise.
func (o *IdentityEntityMergeRequest) GetToEntityId() string {
	if o == nil || o.ToEntityId == nil {
		var ret string
		return ret
	}
	return *o.ToEntityId
}

// GetToEntityIdOk returns a tuple with the ToEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEntityMergeRequest) GetToEntityIdOk() (*string, bool) {
	if o == nil || o.ToEntityId == nil {
		return nil, false
	}
	return o.ToEntityId, true
}

// HasToEntityId returns a boolean if a field has been set.
func (o *IdentityEntityMergeRequest) HasToEntityId() bool {
	if o != nil && o.ToEntityId != nil {
		return true
	}

	return false
}

// SetToEntityId gets a reference to the given string and assigns it to the ToEntityId field.
func (o *IdentityEntityMergeRequest) SetToEntityId(v string) {
	o.ToEntityId = &v
}

func (o IdentityEntityMergeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if o.FromEntityIds != nil {
		toSerialize["from_entity_ids"] = o.FromEntityIds
	}
	if o.ToEntityId != nil {
		toSerialize["to_entity_id"] = o.ToEntityId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityEntityMergeRequest struct {
	value *IdentityEntityMergeRequest
	isSet bool
}

func (v NullableIdentityEntityMergeRequest) Get() *IdentityEntityMergeRequest {
	return v.value
}

func (v *NullableIdentityEntityMergeRequest) Set(val *IdentityEntityMergeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEntityMergeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEntityMergeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEntityMergeRequest(val *IdentityEntityMergeRequest) *NullableIdentityEntityMergeRequest {
	return &NullableIdentityEntityMergeRequest{value: val, isSet: true}
}

func (v NullableIdentityEntityMergeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEntityMergeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


