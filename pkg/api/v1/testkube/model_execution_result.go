/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// execution result returned from executor
type ExecutionResult struct {
	Status *ExecutionStatus `json:"status"`
	// RAW Test execution output, depends of reporter used in particular tool
	Output string `json:"output,omitempty"`
	// output type depends of reporter used in particular tool
	OutputType string `json:"outputType,omitempty"`
	// error message when status is error, separate to output as output can be partial in case of error
	ErrorMessage string `json:"errorMessage,omitempty"`
	// execution steps (for collection of requests)
	Steps   []ExecutionStepResult   `json:"steps,omitempty"`
	Reports *ExecutionResultReports `json:"reports,omitempty"`
}
