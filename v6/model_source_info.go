/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: 3.2.16
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"encoding/json"
)

// SourceInfo struct for SourceInfo
type SourceInfo struct {
	NameSource *string `json:"name_source,omitempty"`
	DomainUrl string `json:"domain_url"`
	Logo *string `json:"logo,omitempty"`
	AdditionalInfo *AdditionalSourceInfo `json:"additional_info,omitempty"`
}

// NewSourceInfo instantiates a new SourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceInfo(domainUrl string) *SourceInfo {
	this := SourceInfo{}
	this.DomainUrl = domainUrl
	return &this
}

// NewSourceInfoWithDefaults instantiates a new SourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceInfoWithDefaults() *SourceInfo {
	this := SourceInfo{}
	return &this
}

// GetNameSource returns the NameSource field value if set, zero value otherwise.
func (o *SourceInfo) GetNameSource() string {
	if o == nil || isNil(o.NameSource) {
		var ret string
		return ret
	}
	return *o.NameSource
}

// GetNameSourceOk returns a tuple with the NameSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceInfo) GetNameSourceOk() (*string, bool) {
	if o == nil || isNil(o.NameSource) {
    return nil, false
	}
	return o.NameSource, true
}

// HasNameSource returns a boolean if a field has been set.
func (o *SourceInfo) HasNameSource() bool {
	if o != nil && !isNil(o.NameSource) {
		return true
	}

	return false
}

// SetNameSource gets a reference to the given string and assigns it to the NameSource field.
func (o *SourceInfo) SetNameSource(v string) {
	o.NameSource = &v
}

// GetDomainUrl returns the DomainUrl field value
func (o *SourceInfo) GetDomainUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainUrl
}

// GetDomainUrlOk returns a tuple with the DomainUrl field value
// and a boolean to check if the value has been set.
func (o *SourceInfo) GetDomainUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DomainUrl, true
}

// SetDomainUrl sets field value
func (o *SourceInfo) SetDomainUrl(v string) {
	o.DomainUrl = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *SourceInfo) GetLogo() string {
	if o == nil || isNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceInfo) GetLogoOk() (*string, bool) {
	if o == nil || isNil(o.Logo) {
    return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *SourceInfo) HasLogo() bool {
	if o != nil && !isNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *SourceInfo) SetLogo(v string) {
	o.Logo = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *SourceInfo) GetAdditionalInfo() AdditionalSourceInfo {
	if o == nil || isNil(o.AdditionalInfo) {
		var ret AdditionalSourceInfo
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceInfo) GetAdditionalInfoOk() (*AdditionalSourceInfo, bool) {
	if o == nil || isNil(o.AdditionalInfo) {
    return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *SourceInfo) HasAdditionalInfo() bool {
	if o != nil && !isNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given AdditionalSourceInfo and assigns it to the AdditionalInfo field.
func (o *SourceInfo) SetAdditionalInfo(v AdditionalSourceInfo) {
	o.AdditionalInfo = &v
}

func (o SourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NameSource) {
		toSerialize["name_source"] = o.NameSource
	}
	if true {
		toSerialize["domain_url"] = o.DomainUrl
	}
	if !isNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !isNil(o.AdditionalInfo) {
		toSerialize["additional_info"] = o.AdditionalInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSourceInfo struct {
	value *SourceInfo
	isSet bool
}

func (v NullableSourceInfo) Get() *SourceInfo {
	return v.value
}

func (v *NullableSourceInfo) Set(val *SourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceInfo(val *SourceInfo) *NullableSourceInfo {
	return &NullableSourceInfo{value: val, isSet: true}
}

func (v NullableSourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

