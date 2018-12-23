package advisor

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
	"github.com/satori/go.uuid"
	"net/http"
)

// RecommendationsClient is the REST APIs for Azure Advisor
type RecommendationsClient struct {
	ManagementClient
}

// NewRecommendationsClient creates an instance of the RecommendationsClient client.
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return NewRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendationsClientWithBaseURI creates an instance of the RecommendationsClient client.
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return RecommendationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Generate initiates the recommendation generation or computation process for a subscription. This operation is
// asynchronous. The generated recommendations are stored in a cache in the Advisor service.
func (client RecommendationsClient) Generate() (result autorest.Response, err error) {
	req, err := client.GeneratePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", resp, "Failure responding to request")
	}

	return
}

// GeneratePreparer prepares the Generate request.
func (client RecommendationsClient) GeneratePreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-12-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GenerateSender sends the Generate request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GenerateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateResponder handles the response to the Generate request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GenerateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get obtains details of a cached recommendation.
//
// resourceURI is the fully qualified Azure Resource Manager identifier of the resource to which the recommendation
// applies. recommendationID is the recommendation ID.
func (client RecommendationsClient) Get(resourceURI string, recommendationID string) (result ResourceRecommendationBase, err error) {
	req, err := client.GetPreparer(resourceURI, recommendationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecommendationsClient) GetPreparer(resourceURI string, recommendationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendationId": autorest.Encode("path", recommendationID),
		"resourceUri":      autorest.Encode("path", resourceURI),
	}

	const APIVersion = "2016-07-12-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetResponder(resp *http.Response) (result ResourceRecommendationBase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGenerateRecommendationsStatus retrieves the status of the recommendation computation or generation process.
// Invoke this API after calling the generation recommendation. The URI of this API is returned in the Location field
// of the response header.
//
// operationID is the operation ID, which can be found from the Location field in the generate recommendation response
// header.
func (client RecommendationsClient) GetGenerateRecommendationsStatus(operationID uuid.UUID) (result autorest.Response, err error) {
	req, err := client.GetGenerateRecommendationsStatusPreparer(operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateRecommendationsStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGenerateRecommendationsStatusSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateRecommendationsStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetGenerateRecommendationsStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateRecommendationsStatus", resp, "Failure responding to request")
	}

	return
}

// GetGenerateRecommendationsStatusPreparer prepares the GetGenerateRecommendationsStatus request.
func (client RecommendationsClient) GetGenerateRecommendationsStatusPreparer(operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-12-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetGenerateRecommendationsStatusSender sends the GetGenerateRecommendationsStatus request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetGenerateRecommendationsStatusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetGenerateRecommendationsStatusResponder handles the response to the GetGenerateRecommendationsStatus request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetGenerateRecommendationsStatusResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List obtains cached recommendations for a subscription. The recommendations are generated or computed by invoking
// generateRecommendations.
//
// filter is the filter to apply to the recommendations. top is the number of recommendations per page if a paged
// version of this API is being used. skipToken is the page-continuation token to use with a paged version of this API.
func (client RecommendationsClient) List(filter string, top *int32, skipToken string) (result ResourceRecommendationBaseListResult, err error) {
	req, err := client.ListPreparer(filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecommendationsClient) ListPreparer(filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-12-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ListResponder(resp *http.Response) (result ResourceRecommendationBaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client RecommendationsClient) ListNextResults(lastResults ResourceRecommendationBaseListResult) (result ResourceRecommendationBaseListResult, err error) {
	req, err := lastResults.ResourceRecommendationBaseListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client RecommendationsClient) ListComplete(filter string, top *int32, skipToken string, cancel <-chan struct{}) (<-chan ResourceRecommendationBase, <-chan error) {
	resultChan := make(chan ResourceRecommendationBase)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(filter, top, skipToken)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
