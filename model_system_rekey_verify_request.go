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

// SystemRekeyVerifyRequest struct for SystemRekeyVerifyRequest
type SystemRekeyVerifyRequest struct {
	// Specifies a single unseal share key from the new set of shares.
	Key *string `json:"key,omitempty"`
	// Specifies the nonce of the rekey verification operation.
	Nonce *string `json:"nonce,omitempty"`
}

// NewSystemRekeyVerifyRequest instantiates a new SystemRekeyVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRekeyVerifyRequest() *SystemRekeyVerifyRequest {
	this := SystemRekeyVerifyRequest{}
	return &this
}

// NewSystemRekeyVerifyRequestWithDefaults instantiates a new SystemRekeyVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRekeyVerifyRequestWithDefaults() *SystemRekeyVerifyRequest {
	this := SystemRekeyVerifyRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SystemRekeyVerifyRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRekeyVerifyRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SystemRekeyVerifyRequest) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SystemRekeyVerifyRequest) SetKey(v string) {
	o.Key = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *SystemRekeyVerifyRequest) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRekeyVerifyRequest) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *SystemRekeyVerifyRequest) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *SystemRekeyVerifyRequest) SetNonce(v string) {
	o.Nonce = &v
}

func (o SystemRekeyVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	return json.Marshal(toSerialize)
}

type NullableSystemRekeyVerifyRequest struct {
	value *SystemRekeyVerifyRequest
	isSet bool
}

func (v NullableSystemRekeyVerifyRequest) Get() *SystemRekeyVerifyRequest {
	return v.value
}

func (v *NullableSystemRekeyVerifyRequest) Set(val *SystemRekeyVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemRekeyVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemRekeyVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemRekeyVerifyRequest(val *SystemRekeyVerifyRequest) *NullableSystemRekeyVerifyRequest {
	return &NullableSystemRekeyVerifyRequest{value: val, isSet: true}
}

func (v NullableSystemRekeyVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemRekeyVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


