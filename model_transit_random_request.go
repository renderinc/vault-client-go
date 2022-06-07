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

// TransitRandomRequest struct for TransitRandomRequest
type TransitRandomRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes *int32 `json:"bytes,omitempty"`
	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format *string `json:"format,omitempty"`
	// Which system to source random data from, ether \"platform\", \"seal\", or \"all\".
	Source *string `json:"source,omitempty"`
	// The number of bytes to generate (POST URL parameter)
	Urlbytes *string `json:"urlbytes,omitempty"`
}

// NewTransitRandomRequest instantiates a new TransitRandomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransitRandomRequest() *TransitRandomRequest {
	this := TransitRandomRequest{}
	var bytes int32 = 32
	this.Bytes = &bytes
	var format string = "base64"
	this.Format = &format
	var source string = "platform"
	this.Source = &source
	return &this
}

// NewTransitRandomRequestWithDefaults instantiates a new TransitRandomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitRandomRequestWithDefaults() *TransitRandomRequest {
	this := TransitRandomRequest{}
	var bytes int32 = 32
	this.Bytes = &bytes
	var format string = "base64"
	this.Format = &format
	var source string = "platform"
	this.Source = &source
	return &this
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *TransitRandomRequest) GetBytes() int32 {
	if o == nil || o.Bytes == nil {
		var ret int32
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitRandomRequest) GetBytesOk() (*int32, bool) {
	if o == nil || o.Bytes == nil {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *TransitRandomRequest) HasBytes() bool {
	if o != nil && o.Bytes != nil {
		return true
	}

	return false
}

// SetBytes gets a reference to the given int32 and assigns it to the Bytes field.
func (o *TransitRandomRequest) SetBytes(v int32) {
	o.Bytes = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *TransitRandomRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitRandomRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *TransitRandomRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *TransitRandomRequest) SetFormat(v string) {
	o.Format = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TransitRandomRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitRandomRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TransitRandomRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TransitRandomRequest) SetSource(v string) {
	o.Source = &v
}

// GetUrlbytes returns the Urlbytes field value if set, zero value otherwise.
func (o *TransitRandomRequest) GetUrlbytes() string {
	if o == nil || o.Urlbytes == nil {
		var ret string
		return ret
	}
	return *o.Urlbytes
}

// GetUrlbytesOk returns a tuple with the Urlbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitRandomRequest) GetUrlbytesOk() (*string, bool) {
	if o == nil || o.Urlbytes == nil {
		return nil, false
	}
	return o.Urlbytes, true
}

// HasUrlbytes returns a boolean if a field has been set.
func (o *TransitRandomRequest) HasUrlbytes() bool {
	if o != nil && o.Urlbytes != nil {
		return true
	}

	return false
}

// SetUrlbytes gets a reference to the given string and assigns it to the Urlbytes field.
func (o *TransitRandomRequest) SetUrlbytes(v string) {
	o.Urlbytes = &v
}

func (o TransitRandomRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bytes != nil {
		toSerialize["bytes"] = o.Bytes
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Urlbytes != nil {
		toSerialize["urlbytes"] = o.Urlbytes
	}
	return json.Marshal(toSerialize)
}

type NullableTransitRandomRequest struct {
	value *TransitRandomRequest
	isSet bool
}

func (v NullableTransitRandomRequest) Get() *TransitRandomRequest {
	return v.value
}

func (v *NullableTransitRandomRequest) Set(val *TransitRandomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitRandomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitRandomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitRandomRequest(val *TransitRandomRequest) *NullableTransitRandomRequest {
	return &NullableTransitRandomRequest{value: val, isSet: true}
}

func (v NullableTransitRandomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitRandomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


