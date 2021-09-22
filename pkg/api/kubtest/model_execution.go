/*
 * Kubtest API
 *
 * Kubtest provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: kubtest@kubshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kubtest

// API server script execution
type Execution struct {
	// execution id
	Id string `json:"id,omitempty"`
	// unique script name (CRD Script name)
	ScriptName string `json:"scriptName,omitempty"`
	// script type e.g. postman/collection
	ScriptType string `json:"scriptType,omitempty"`
	// execution name
	// script metadata content
	ScriptContent string      `json:"scriptContent,omitempty"`
	Repository    *Repository `json:"repository,omitempty"`
	Name          string      `json:"name,omitempty"`
	// execution envs passed to executor
	Envs map[string]string `json:"envs,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
	Result *Result           `json:"execution,omitempty"`
}
