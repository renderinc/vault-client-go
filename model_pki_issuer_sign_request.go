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

// PkiIssuerSignRequest struct for PkiIssuerSignRequest
type PkiIssuerSignRequest struct {
	// The requested Subject Alternative Names, if any, in a comma-delimited list. If email protection is enabled for the role, this may contain email addresses.
	AltNames *string `json:"alt_names,omitempty"`
	// The requested common name; if you want more than one, specify the alternative names in the alt_names map. If email protection is enabled in the role, this may be an email address.
	CommonName *string `json:"common_name,omitempty"`
	// PEM-format CSR to be signed.
	Csr *string `json:"csr,omitempty"`
	// If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included).
	ExcludeCnFromSans *bool `json:"exclude_cn_from_sans,omitempty"`
	// Format for returned data. Can be \"pem\", \"der\", or \"pem_bundle\". If \"pem_bundle\", any private key and issuing cert will be appended to the certificate pem. If \"der\", the value will be base64 encoded. Defaults to \"pem\".
	Format *string `json:"format,omitempty"`
	// The requested IP SANs, if any, in a comma-delimited list
	IpSans []string `json:"ip_sans,omitempty"`
	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ
	NotAfter *string `json:"not_after,omitempty"`
	// Requested other SANs, in an array with the format <oid>;UTF8:<utf8 string value> for each entry.
	OtherSans []string `json:"other_sans,omitempty"`
	// Format for the returned private key. Generally the default will be controlled by the \"format\" parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \"pkcs8\" to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \"der\".
	PrivateKeyFormat *string `json:"private_key_format,omitempty"`
	// The Subject's requested serial number, if any. See RFC 4519 Section 2.31 'serialNumber' for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate's Serial Number field.
	SerialNumber *string `json:"serial_number,omitempty"`
	// The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the role max TTL.
	Ttl *int32 `json:"ttl,omitempty"`
	// The requested URI SANs, if any, in a comma-delimited list.
	UriSans []string `json:"uri_sans,omitempty"`
}

// NewPkiIssuerSignRequest instantiates a new PkiIssuerSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkiIssuerSignRequest() *PkiIssuerSignRequest {
	this := PkiIssuerSignRequest{}
	var csr string = ""
	this.Csr = &csr
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	return &this
}

// NewPkiIssuerSignRequestWithDefaults instantiates a new PkiIssuerSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignRequestWithDefaults() *PkiIssuerSignRequest {
	this := PkiIssuerSignRequest{}
	var csr string = ""
	this.Csr = &csr
	var excludeCnFromSans bool = false
	this.ExcludeCnFromSans = &excludeCnFromSans
	var format string = "pem"
	this.Format = &format
	var privateKeyFormat string = "der"
	this.PrivateKeyFormat = &privateKeyFormat
	return &this
}

// GetAltNames returns the AltNames field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetAltNames() string {
	if o == nil || o.AltNames == nil {
		var ret string
		return ret
	}
	return *o.AltNames
}

// GetAltNamesOk returns a tuple with the AltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetAltNamesOk() (*string, bool) {
	if o == nil || o.AltNames == nil {
		return nil, false
	}
	return o.AltNames, true
}

// HasAltNames returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasAltNames() bool {
	if o != nil && o.AltNames != nil {
		return true
	}

	return false
}

// SetAltNames gets a reference to the given string and assigns it to the AltNames field.
func (o *PkiIssuerSignRequest) SetAltNames(v string) {
	o.AltNames = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkiIssuerSignRequest) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetCsr() string {
	if o == nil || o.Csr == nil {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetCsrOk() (*string, bool) {
	if o == nil || o.Csr == nil {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasCsr() bool {
	if o != nil && o.Csr != nil {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *PkiIssuerSignRequest) SetCsr(v string) {
	o.Csr = &v
}

// GetExcludeCnFromSans returns the ExcludeCnFromSans field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetExcludeCnFromSans() bool {
	if o == nil || o.ExcludeCnFromSans == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCnFromSans
}

// GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetExcludeCnFromSansOk() (*bool, bool) {
	if o == nil || o.ExcludeCnFromSans == nil {
		return nil, false
	}
	return o.ExcludeCnFromSans, true
}

// HasExcludeCnFromSans returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasExcludeCnFromSans() bool {
	if o != nil && o.ExcludeCnFromSans != nil {
		return true
	}

	return false
}

// SetExcludeCnFromSans gets a reference to the given bool and assigns it to the ExcludeCnFromSans field.
func (o *PkiIssuerSignRequest) SetExcludeCnFromSans(v bool) {
	o.ExcludeCnFromSans = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PkiIssuerSignRequest) SetFormat(v string) {
	o.Format = &v
}

// GetIpSans returns the IpSans field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetIpSans() []string {
	if o == nil || o.IpSans == nil {
		var ret []string
		return ret
	}
	return o.IpSans
}

// GetIpSansOk returns a tuple with the IpSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetIpSansOk() ([]string, bool) {
	if o == nil || o.IpSans == nil {
		return nil, false
	}
	return o.IpSans, true
}

// HasIpSans returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasIpSans() bool {
	if o != nil && o.IpSans != nil {
		return true
	}

	return false
}

// SetIpSans gets a reference to the given []string and assigns it to the IpSans field.
func (o *PkiIssuerSignRequest) SetIpSans(v []string) {
	o.IpSans = v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetNotAfter() string {
	if o == nil || o.NotAfter == nil {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetNotAfterOk() (*string, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *PkiIssuerSignRequest) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetOtherSans returns the OtherSans field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetOtherSans() []string {
	if o == nil || o.OtherSans == nil {
		var ret []string
		return ret
	}
	return o.OtherSans
}

// GetOtherSansOk returns a tuple with the OtherSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetOtherSansOk() ([]string, bool) {
	if o == nil || o.OtherSans == nil {
		return nil, false
	}
	return o.OtherSans, true
}

// HasOtherSans returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasOtherSans() bool {
	if o != nil && o.OtherSans != nil {
		return true
	}

	return false
}

// SetOtherSans gets a reference to the given []string and assigns it to the OtherSans field.
func (o *PkiIssuerSignRequest) SetOtherSans(v []string) {
	o.OtherSans = v
}

// GetPrivateKeyFormat returns the PrivateKeyFormat field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetPrivateKeyFormat() string {
	if o == nil || o.PrivateKeyFormat == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyFormat
}

// GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetPrivateKeyFormatOk() (*string, bool) {
	if o == nil || o.PrivateKeyFormat == nil {
		return nil, false
	}
	return o.PrivateKeyFormat, true
}

// HasPrivateKeyFormat returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasPrivateKeyFormat() bool {
	if o != nil && o.PrivateKeyFormat != nil {
		return true
	}

	return false
}

// SetPrivateKeyFormat gets a reference to the given string and assigns it to the PrivateKeyFormat field.
func (o *PkiIssuerSignRequest) SetPrivateKeyFormat(v string) {
	o.PrivateKeyFormat = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PkiIssuerSignRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *PkiIssuerSignRequest) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUriSans returns the UriSans field value if set, zero value otherwise.
func (o *PkiIssuerSignRequest) GetUriSans() []string {
	if o == nil || o.UriSans == nil {
		var ret []string
		return ret
	}
	return o.UriSans
}

// GetUriSansOk returns a tuple with the UriSans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkiIssuerSignRequest) GetUriSansOk() ([]string, bool) {
	if o == nil || o.UriSans == nil {
		return nil, false
	}
	return o.UriSans, true
}

// HasUriSans returns a boolean if a field has been set.
func (o *PkiIssuerSignRequest) HasUriSans() bool {
	if o != nil && o.UriSans != nil {
		return true
	}

	return false
}

// SetUriSans gets a reference to the given []string and assigns it to the UriSans field.
func (o *PkiIssuerSignRequest) SetUriSans(v []string) {
	o.UriSans = v
}

func (o PkiIssuerSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltNames != nil {
		toSerialize["alt_names"] = o.AltNames
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Csr != nil {
		toSerialize["csr"] = o.Csr
	}
	if o.ExcludeCnFromSans != nil {
		toSerialize["exclude_cn_from_sans"] = o.ExcludeCnFromSans
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.IpSans != nil {
		toSerialize["ip_sans"] = o.IpSans
	}
	if o.NotAfter != nil {
		toSerialize["not_after"] = o.NotAfter
	}
	if o.OtherSans != nil {
		toSerialize["other_sans"] = o.OtherSans
	}
	if o.PrivateKeyFormat != nil {
		toSerialize["private_key_format"] = o.PrivateKeyFormat
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.UriSans != nil {
		toSerialize["uri_sans"] = o.UriSans
	}
	return json.Marshal(toSerialize)
}

type NullablePkiIssuerSignRequest struct {
	value *PkiIssuerSignRequest
	isSet bool
}

func (v NullablePkiIssuerSignRequest) Get() *PkiIssuerSignRequest {
	return v.value
}

func (v *NullablePkiIssuerSignRequest) Set(val *PkiIssuerSignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePkiIssuerSignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePkiIssuerSignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkiIssuerSignRequest(val *PkiIssuerSignRequest) *NullablePkiIssuerSignRequest {
	return &NullablePkiIssuerSignRequest{value: val, isSet: true}
}

func (v NullablePkiIssuerSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkiIssuerSignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


