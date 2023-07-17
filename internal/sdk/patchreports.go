// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"jamf/internal/sdk/pkg/models/operations"
	"jamf/internal/sdk/pkg/models/shared"
	"jamf/internal/sdk/pkg/utils"
	"net/http"
)

type patchreports struct {
	sdkConfiguration sdkConfiguration
}

func newPatchreports(sdkConfig sdkConfiguration) *patchreports {
	return &patchreports{
		sdkConfiguration: sdkConfig,
	}
}

// PatchreportsPatchsoftwaretitleidByIDGet - Finds patch reports by patch software title ID. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}/patch-report".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchreports) PatchreportsPatchsoftwaretitleidByIDGet(ctx context.Context, request operations.PatchreportsPatchsoftwaretitleidByIDGetRequest) (*operations.PatchreportsPatchsoftwaretitleidByIDGetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchreports/patchsoftwaretitleid/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/xml;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchreportsPatchsoftwaretitleidByIDGetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PatchReport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PatchReport = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}

// PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGet - Display computers for a specific version of a patch report. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}/patch-report".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchreports) PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGet(ctx context.Context, request operations.PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetRequest) (*operations.PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchreports/patchsoftwaretitleid/{id}/version/{version}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/xml;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PatchReport
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PatchReport = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}