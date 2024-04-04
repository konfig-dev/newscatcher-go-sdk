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

// SubscriptionResponse struct for SubscriptionResponse
type SubscriptionResponse struct {
	Active bool `json:"active"`
	CallsPerSeconds *int32 `json:"calls_per_seconds,omitempty"`
	PlanName string `json:"plan_name"`
	UsageAssignedCalls *int32 `json:"usage_assigned_calls,omitempty"`
	UsageRemainingCalls *int32 `json:"usage_remaining_calls,omitempty"`
	HistoricalDays *int32 `json:"historical_days,omitempty"`
}

// NewSubscriptionResponse instantiates a new SubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionResponse(active bool, planName string) *SubscriptionResponse {
	this := SubscriptionResponse{}
	this.Active = active
	this.PlanName = planName
	return &this
}

// NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionResponseWithDefaults() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *SubscriptionResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SubscriptionResponse) SetActive(v bool) {
	o.Active = v
}

// GetCallsPerSeconds returns the CallsPerSeconds field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetCallsPerSeconds() int32 {
	if o == nil || isNil(o.CallsPerSeconds) {
		var ret int32
		return ret
	}
	return *o.CallsPerSeconds
}

// GetCallsPerSecondsOk returns a tuple with the CallsPerSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetCallsPerSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.CallsPerSeconds) {
    return nil, false
	}
	return o.CallsPerSeconds, true
}

// HasCallsPerSeconds returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCallsPerSeconds() bool {
	if o != nil && !isNil(o.CallsPerSeconds) {
		return true
	}

	return false
}

// SetCallsPerSeconds gets a reference to the given int32 and assigns it to the CallsPerSeconds field.
func (o *SubscriptionResponse) SetCallsPerSeconds(v int32) {
	o.CallsPerSeconds = &v
}

// GetPlanName returns the PlanName field value
func (o *SubscriptionResponse) GetPlanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetPlanNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlanName, true
}

// SetPlanName sets field value
func (o *SubscriptionResponse) SetPlanName(v string) {
	o.PlanName = v
}

// GetUsageAssignedCalls returns the UsageAssignedCalls field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetUsageAssignedCalls() int32 {
	if o == nil || isNil(o.UsageAssignedCalls) {
		var ret int32
		return ret
	}
	return *o.UsageAssignedCalls
}

// GetUsageAssignedCallsOk returns a tuple with the UsageAssignedCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetUsageAssignedCallsOk() (*int32, bool) {
	if o == nil || isNil(o.UsageAssignedCalls) {
    return nil, false
	}
	return o.UsageAssignedCalls, true
}

// HasUsageAssignedCalls returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasUsageAssignedCalls() bool {
	if o != nil && !isNil(o.UsageAssignedCalls) {
		return true
	}

	return false
}

// SetUsageAssignedCalls gets a reference to the given int32 and assigns it to the UsageAssignedCalls field.
func (o *SubscriptionResponse) SetUsageAssignedCalls(v int32) {
	o.UsageAssignedCalls = &v
}

// GetUsageRemainingCalls returns the UsageRemainingCalls field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetUsageRemainingCalls() int32 {
	if o == nil || isNil(o.UsageRemainingCalls) {
		var ret int32
		return ret
	}
	return *o.UsageRemainingCalls
}

// GetUsageRemainingCallsOk returns a tuple with the UsageRemainingCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetUsageRemainingCallsOk() (*int32, bool) {
	if o == nil || isNil(o.UsageRemainingCalls) {
    return nil, false
	}
	return o.UsageRemainingCalls, true
}

// HasUsageRemainingCalls returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasUsageRemainingCalls() bool {
	if o != nil && !isNil(o.UsageRemainingCalls) {
		return true
	}

	return false
}

// SetUsageRemainingCalls gets a reference to the given int32 and assigns it to the UsageRemainingCalls field.
func (o *SubscriptionResponse) SetUsageRemainingCalls(v int32) {
	o.UsageRemainingCalls = &v
}

// GetHistoricalDays returns the HistoricalDays field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetHistoricalDays() int32 {
	if o == nil || isNil(o.HistoricalDays) {
		var ret int32
		return ret
	}
	return *o.HistoricalDays
}

// GetHistoricalDaysOk returns a tuple with the HistoricalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetHistoricalDaysOk() (*int32, bool) {
	if o == nil || isNil(o.HistoricalDays) {
    return nil, false
	}
	return o.HistoricalDays, true
}

// HasHistoricalDays returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasHistoricalDays() bool {
	if o != nil && !isNil(o.HistoricalDays) {
		return true
	}

	return false
}

// SetHistoricalDays gets a reference to the given int32 and assigns it to the HistoricalDays field.
func (o *SubscriptionResponse) SetHistoricalDays(v int32) {
	o.HistoricalDays = &v
}

func (o SubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.CallsPerSeconds) {
		toSerialize["calls_per_seconds"] = o.CallsPerSeconds
	}
	if true {
		toSerialize["plan_name"] = o.PlanName
	}
	if !isNil(o.UsageAssignedCalls) {
		toSerialize["usage_assigned_calls"] = o.UsageAssignedCalls
	}
	if !isNil(o.UsageRemainingCalls) {
		toSerialize["usage_remaining_calls"] = o.UsageRemainingCalls
	}
	if !isNil(o.HistoricalDays) {
		toSerialize["historical_days"] = o.HistoricalDays
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionResponse struct {
	value *SubscriptionResponse
	isSet bool
}

func (v NullableSubscriptionResponse) Get() *SubscriptionResponse {
	return v.value
}

func (v *NullableSubscriptionResponse) Set(val *SubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponse(val *SubscriptionResponse) *NullableSubscriptionResponse {
	return &NullableSubscriptionResponse{value: val, isSet: true}
}

func (v NullableSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


