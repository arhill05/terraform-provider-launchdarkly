/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.24
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type PatchComment struct {
	Comment string `json:"comment,omitempty"`
	Patch []PatchOperation `json:"patch,omitempty"`
}
