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

// UserpassLoginRequest struct for UserpassLoginRequest
type UserpassLoginRequest struct {
	// Password for this user.
	Password *string `json:"password,omitempty"`
}

// NewUserpassLoginRequest instantiates a new UserpassLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserpassLoginRequest() *UserpassLoginRequest {
	this := UserpassLoginRequest{}
	return &this
}

// NewUserpassLoginRequestWithDefaults instantiates a new UserpassLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassLoginRequestWithDefaults() *UserpassLoginRequest {
	this := UserpassLoginRequest{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserpassLoginRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserpassLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserpassLoginRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserpassLoginRequest) SetPassword(v string) {
	o.Password = &v
}

func (o UserpassLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserpassLoginRequest struct {
	value *UserpassLoginRequest
	isSet bool
}

func (v NullableUserpassLoginRequest) Get() *UserpassLoginRequest {
	return v.value
}

func (v *NullableUserpassLoginRequest) Set(val *UserpassLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserpassLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserpassLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserpassLoginRequest(val *UserpassLoginRequest) *NullableUserpassLoginRequest {
	return &NullableUserpassLoginRequest{value: val, isSet: true}
}

func (v NullableUserpassLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserpassLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


