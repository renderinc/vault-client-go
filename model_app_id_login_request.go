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

// AppIdLoginRequest struct for AppIdLoginRequest
type AppIdLoginRequest struct {
	// The unique user ID
	UserId *string `json:"user_id,omitempty"`
}

// NewAppIdLoginRequest instantiates a new AppIdLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppIdLoginRequest() *AppIdLoginRequest {
	this := AppIdLoginRequest{}
	return &this
}

// NewAppIdLoginRequestWithDefaults instantiates a new AppIdLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppIdLoginRequestWithDefaults() *AppIdLoginRequest {
	this := AppIdLoginRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AppIdLoginRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppIdLoginRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AppIdLoginRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AppIdLoginRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o AppIdLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableAppIdLoginRequest struct {
	value *AppIdLoginRequest
	isSet bool
}

func (v NullableAppIdLoginRequest) Get() *AppIdLoginRequest {
	return v.value
}

func (v *NullableAppIdLoginRequest) Set(val *AppIdLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppIdLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppIdLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppIdLoginRequest(val *AppIdLoginRequest) *NullableAppIdLoginRequest {
	return &NullableAppIdLoginRequest{value: val, isSet: true}
}

func (v NullableAppIdLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppIdLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


