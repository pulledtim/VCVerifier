/*
 * vcverifier
 *
 * Backend component to verify credentials
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TokenResponse struct {

	TokenType string `json:"token_type,omitempty"`

	ExpiresIn float32 `json:"expires_in,omitempty"`

	AccessToken string `json:"access_token,omitempty"`
}
