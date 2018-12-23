package recoveryservicesbackup

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ProtectedItemOperationResultsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectedItemOperationResultsClient struct {
	ManagementClient
}

// NewProtectedItemOperationResultsClient creates an instance of the ProtectedItemOperationResultsClient client.
func NewProtectedItemOperationResultsClient(subscriptionID string) ProtectedItemOperationResultsClient {
	return NewProtectedItemOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectedItemOperationResultsClientWithBaseURI creates an instance of the ProtectedItemOperationResultsClient
// client.
func NewProtectedItemOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) ProtectedItemOperationResultsClient {
	return ProtectedItemOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get fetches the result of any operation on the backup item.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is fabric name associated with the backup item. containerName is
// container name associated with the backup item. protectedItemName is backup item name whose details are to be
// fetched. operationID is operationID which represents the operation whose result needs to be fetched.
func (client ProtectedItemOperationResultsClient) Get(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result ProtectedItemResource, err error) {
	req, err := client.GetPreparer(vaultName, resourceGroupName, fabricName, containerName, protectedItemName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemOperationResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemOperationResultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemOperationResultsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectedItemOperationResultsClient) GetPreparer(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"operationId":       autorest.Encode("path", operationID),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemOperationResultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectedItemOperationResultsClient) GetResponder(resp *http.Response) (result ProtectedItemResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
