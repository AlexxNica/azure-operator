package redis

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// FirewallRuleClient is the REST API for Azure Redis Cache Service.
type FirewallRuleClient struct {
	BaseClient
}

// NewFirewallRuleClient creates an instance of the FirewallRuleClient client.
func NewFirewallRuleClient(subscriptionID string) FirewallRuleClient {
	return NewFirewallRuleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallRuleClientWithBaseURI creates an instance of the FirewallRuleClient client.
func NewFirewallRuleClientWithBaseURI(baseURI string, subscriptionID string) FirewallRuleClient {
	return FirewallRuleClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a redis cache firewall rule
//
// resourceGroupName is the name of the resource group. cacheName is the name of the Redis cache. ruleName is the name
// of the firewall rule. parameters is parameters supplied to the create or update redis firewall rule operation.
func (client FirewallRuleClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule) (result FirewallRule, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FirewallRuleProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIP", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.FirewallRuleProperties.EndIP", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "redis.FirewallRuleClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, cacheName, ruleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FirewallRuleClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a single firewall rule in a specified redis cache.
//
// resourceGroupName is the name of the resource group. cacheName is the name of the Redis cache. ruleName is the name
// of the firewall rule.
func (client FirewallRuleClient) Delete(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, cacheName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FirewallRuleClient) DeletePreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a single firewall rule in a specified redis cache.
//
// resourceGroupName is the name of the resource group. cacheName is the name of the Redis cache. ruleName is the name
// of the firewall rule.
func (client FirewallRuleClient) Get(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result FirewallRule, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, cacheName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallRuleClient) GetPreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) GetResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
