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

// Cluster Cluster DTO class.
type Cluster struct {
	ClusterId string `json:"cluster_id"`
	ClusterSize int32 `json:"cluster_size"`
	Articles []map[string]interface{} `json:"articles"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(clusterId string, clusterSize int32, articles []map[string]interface{}) *Cluster {
	this := Cluster{}
	this.ClusterId = clusterId
	this.ClusterSize = clusterSize
	this.Articles = articles
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *Cluster) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetClusterIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *Cluster) SetClusterId(v string) {
	o.ClusterId = v
}

// GetClusterSize returns the ClusterSize field value
func (o *Cluster) GetClusterSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterSize
}

// GetClusterSizeOk returns a tuple with the ClusterSize field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetClusterSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClusterSize, true
}

// SetClusterSize sets field value
func (o *Cluster) SetClusterSize(v int32) {
	o.ClusterSize = v
}

// GetArticles returns the Articles field value
func (o *Cluster) GetArticles() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetArticlesOk() ([]map[string]interface{}, bool) {
	if o == nil {
    return nil, false
	}
	return o.Articles, true
}

// SetArticles sets field value
func (o *Cluster) SetArticles(v []map[string]interface{}) {
	o.Articles = v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["cluster_size"] = o.ClusterSize
	}
	if true {
		toSerialize["articles"] = o.Articles
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


