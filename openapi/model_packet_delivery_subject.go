/*
 * vcverifier
 *
 * Backend component to verify credentials
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PacketDeliverySubject struct {

	Id string `json:"id,omitempty"`

	FamilyName string `json:"familyName,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Roles []SubjectRole `json:"roles,omitempty"`

	Email string `json:"email,omitempty"`
}
