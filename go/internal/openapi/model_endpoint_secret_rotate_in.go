/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointSecretRotateIn struct for EndpointSecretRotateIn
type EndpointSecretRotateIn struct {
	// The endpoint's verification secret. If `null` is passed, a secret is automatically generated.
	Key *string `json:"key,omitempty"`
}

// NewEndpointSecretRotateIn instantiates a new EndpointSecretRotateIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointSecretRotateIn() *EndpointSecretRotateIn {
	this := EndpointSecretRotateIn{}
	return &this
}

// NewEndpointSecretRotateInWithDefaults instantiates a new EndpointSecretRotateIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointSecretRotateInWithDefaults() *EndpointSecretRotateIn {
	this := EndpointSecretRotateIn{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EndpointSecretRotateIn) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSecretRotateIn) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EndpointSecretRotateIn) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EndpointSecretRotateIn) SetKey(v string) {
	o.Key = &v
}

func (o EndpointSecretRotateIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointSecretRotateIn struct {
	value *EndpointSecretRotateIn
	isSet bool
}

func (v NullableEndpointSecretRotateIn) Get() *EndpointSecretRotateIn {
	return v.value
}

func (v *NullableEndpointSecretRotateIn) Set(val *EndpointSecretRotateIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointSecretRotateIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointSecretRotateIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointSecretRotateIn(val *EndpointSecretRotateIn) *NullableEndpointSecretRotateIn {
	return &NullableEndpointSecretRotateIn{value: val, isSet: true}
}

func (v NullableEndpointSecretRotateIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointSecretRotateIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


