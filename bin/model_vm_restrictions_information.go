/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VmRestrictionsInformation struct {
	Id string `json:"id"`
	// The organization manages the VM
	ManagedOrg string `json:"managedOrg,omitempty"`
	Integrityconstraint string `json:"integrityconstraint,omitempty"`
	Cpu *Vmcpu `json:"cpu,omitempty"`
	Memory *VmMemory `json:"memory,omitempty"`
	ApplianceView *VmApplianceView `json:"applianceView,omitempty"`
	CddvdList *VmConnectedDeviceList `json:"cddvdList,omitempty"`
	FloopyList *VmConnectedDeviceList `json:"floopyList,omitempty"`
	FirewareType *Number `json:"firewareType,omitempty"`
	GuestIsolation *VmGuestIsolation `json:"guestIsolation,omitempty"`
	Niclist *NicDevices `json:"niclist,omitempty"`
	ParallelPortList *VmConnectedDeviceList `json:"parallelPortList,omitempty"`
	SerialPortList *VmConnectedDeviceList `json:"serialPortList,omitempty"`
	UsbList *VmUsbList `json:"usbList,omitempty"`
	RemoteVNC *VmRemoteVnc `json:"remoteVNC,omitempty"`
}
