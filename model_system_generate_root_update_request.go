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

// SystemGenerateRootUpdateRequest struct for SystemGenerateRootUpdateRequest
type SystemGenerateRootUpdateRequest struct {
	// Specifies a single unseal key share.
	Key *string `json:"key,omitempty"`
	// Specifies the nonce of the attempt.
	Nonce *string `json:"nonce,omitempty"`
}

// NewSystemGenerateRootUpdateRequest instantiates a new SystemGenerateRootUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemGenerateRootUpdateRequest() *SystemGenerateRootUpdateRequest {
	this := SystemGenerateRootUpdateRequest{}
	return &this
}

// NewSystemGenerateRootUpdateRequestWithDefaults instantiates a new SystemGenerateRootUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemGenerateRootUpdateRequestWithDefaults() *SystemGenerateRootUpdateRequest {
	this := SystemGenerateRootUpdateRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SystemGenerateRootUpdateRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemGenerateRootUpdateRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SystemGenerateRootUpdateRequest) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SystemGenerateRootUpdateRequest) SetKey(v string) {
	o.Key = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *SystemGenerateRootUpdateRequest) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemGenerateRootUpdateRequest) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *SystemGenerateRootUpdateRequest) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *SystemGenerateRootUpdateRequest) SetNonce(v string) {
	o.Nonce = &v
}

func (o SystemGenerateRootUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	return json.Marshal(toSerialize)
}

type NullableSystemGenerateRootUpdateRequest struct {
	value *SystemGenerateRootUpdateRequest
	isSet bool
}

func (v NullableSystemGenerateRootUpdateRequest) Get() *SystemGenerateRootUpdateRequest {
	return v.value
}

func (v *NullableSystemGenerateRootUpdateRequest) Set(val *SystemGenerateRootUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemGenerateRootUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemGenerateRootUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemGenerateRootUpdateRequest(val *SystemGenerateRootUpdateRequest) *NullableSystemGenerateRootUpdateRequest {
	return &NullableSystemGenerateRootUpdateRequest{value: val, isSet: true}
}

func (v NullableSystemGenerateRootUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemGenerateRootUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


