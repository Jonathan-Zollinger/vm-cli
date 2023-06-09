/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VmConnectedDevice struct {
	Index *Number `json:"index,omitempty"`
	StartConnected string `json:"startConnected,omitempty"`
	ConnectionStatus *Number `json:"connectionStatus,omitempty"`
	DevicePath string `json:"devicePath,omitempty"`
}
