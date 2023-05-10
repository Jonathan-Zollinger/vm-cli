# API documentation for vmware workstation pro 17 

vmrest 1.3.0 build-21139696

## Overview
This API client greatly generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  Shout out to [swagger-spec](https://github.com/swagger-api/swagger-spec) for a smooth experience and making this tool possible.

- API version: 1.3.0
- Package version: 0.0.1
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8697*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HostNetworksManagementApi* | [**CreateNetwork**](HostNetworksManagementApi.md#createnetwork) | **Post** /vmnets | Creates a virtual network
*HostNetworksManagementApi* | [**DeletePortforward**](HostNetworksManagementApi.md#deleteportforward) | **Delete** /vmnet/{vmnet}/portforward/{protocol}/{port} | Deletes port forwarding
*HostNetworksManagementApi* | [**GetAllNetworks**](HostNetworksManagementApi.md#getallnetworks) | **Get** /vmnet | Returns all virtual networks
*HostNetworksManagementApi* | [**GetMACToIPs**](HostNetworksManagementApi.md#getmactoips) | **Get** /vmnet/{vmnet}/mactoip | Returns all MAC-to-IP settings for DHCP service
*HostNetworksManagementApi* | [**GetPortforwards**](HostNetworksManagementApi.md#getportforwards) | **Get** /vmnet/{vmnet}/portforward | Returns all port forwardings
*HostNetworksManagementApi* | [**UpdateMacToIP**](HostNetworksManagementApi.md#updatemactoip) | **Put** /vmnet/{vmnet}/mactoip/{mac} | Updates the MAC-to-IP binding
*HostNetworksManagementApi* | [**UpdatePortforward**](HostNetworksManagementApi.md#updateportforward) | **Put** /vmnet/{vmnet}/portforward/{protocol}/{port} | Updates port forwarding
*VMManagementApi* | [**ConfigVMParams**](VMManagementApi.md#configvmparams) | **Put** /vms/{id}/configparams | update the vm config params
*VMManagementApi* | [**CreateVM**](VMManagementApi.md#createvm) | **Post** /vms | Creates a copy of the VM
*VMManagementApi* | [**DeleteVM**](VMManagementApi.md#deletevm) | **Delete** /vms/{id} | Deletes a VM
*VMManagementApi* | [**GetAllVMs**](VMManagementApi.md#getallvms) | **Get** /vms | Returns a list of VM IDs and paths for all VMs
*VMManagementApi* | [**GetVM**](VMManagementApi.md#getvm) | **Get** /vms/{id} | Returns the VM setting information of a VM
*VMManagementApi* | [**GetVMParams**](VMManagementApi.md#getvmparams) | **Get** /vms/{id}/params/{name} | Get the VM config params
*VMManagementApi* | [**GetVMRestrictions**](VMManagementApi.md#getvmrestrictions) | **Get** /vms/{id}/restrictions | Returns the restrictions information of the VM
*VMManagementApi* | [**RegisterVM**](VMManagementApi.md#registervm) | **Post** /vms/registration | Register VM to VM Library
*VMManagementApi* | [**UpdateVM**](VMManagementApi.md#updatevm) | **Put** /vms/{id} | Updates the VM settings
*VMNetworkAdaptersManagementApi* | [**CreateNICDevice**](VMNetworkAdaptersManagementApi.md#createnicdevice) | **Post** /vms/{id}/nic | Creates a network adapter in the VM
*VMNetworkAdaptersManagementApi* | [**DeleteNICDevice**](VMNetworkAdaptersManagementApi.md#deletenicdevice) | **Delete** /vms/{id}/nic/{index} | Deletes a VM network adapter
*VMNetworkAdaptersManagementApi* | [**GetAllNICDevices**](VMNetworkAdaptersManagementApi.md#getallnicdevices) | **Get** /vms/{id}/nic | Returns all network adapters in the VM
*VMNetworkAdaptersManagementApi* | [**GetIPAddress**](VMNetworkAdaptersManagementApi.md#getipaddress) | **Get** /vms/{id}/ip | Returns the IP address of a VM
*VMNetworkAdaptersManagementApi* | [**GetNicInfo**](VMNetworkAdaptersManagementApi.md#getnicinfo) | **Get** /vms/{id}/nicips | Returns the IP stack configuration of all NICs of a VM
*VMNetworkAdaptersManagementApi* | [**UpdateNICDevice**](VMNetworkAdaptersManagementApi.md#updatenicdevice) | **Put** /vms/{id}/nic/{index} | Updates a network adapter in the VM
*VMPowerManagementApi* | [**ChangePowerState**](VMPowerManagementApi.md#changepowerstate) | **Put** /vms/{id}/power | Changes the VM power state
*VMPowerManagementApi* | [**GetPowerState**](VMPowerManagementApi.md#getpowerstate) | **Get** /vms/{id}/power | Returns the power state of the VM
*VMSharedFoldersManagementApi* | [**CreateSharedFolder**](VMSharedFoldersManagementApi.md#createsharedfolder) | **Post** /vms/{id}/sharedfolders | Mounts a new shared folder in the VM
*VMSharedFoldersManagementApi* | [**DeleteSharedFolder**](VMSharedFoldersManagementApi.md#deletesharedfolder) | **Delete** /vms/{id}/sharedfolders/{folder id} | Deletes a shared folder
*VMSharedFoldersManagementApi* | [**GetAllSharedFolders**](VMSharedFoldersManagementApi.md#getallsharedfolders) | **Get** /vms/{id}/sharedfolders | Returns all shared folders mounted in the VM
*VMSharedFoldersManagementApi* | [**UpdataSharedFolder**](VMSharedFoldersManagementApi.md#updatasharedfolder) | **Put** /vms/{id}/sharedfolders/{folder id} | Updates a shared folder mounted in the VM


## Documentation For Models

 - [ConfigVmParamsParameter](ConfigVmParamsParameter.md)
 - [CreateVmnetParameter](CreateVmnetParameter.md)
 - [DaemonState](DaemonState.md)
 - [DhcpConfig](DhcpConfig.md)
 - [DnsConfig](DnsConfig.md)
 - [ErrorModel](ErrorModel.md)
 - [InlineResponse200](InlineResponse200.md)
 - [IpAddress](IpAddress.md)
 - [IpNetAddress](IpNetAddress.md)
 - [MacAddress](MacAddress.md)
 - [MacToIpParameter](MacToIpParameter.md)
 - [MactoIp](MactoIp.md)
 - [MactoIps](MactoIps.md)
 - [Network](Network.md)
 - [Networks](Networks.md)
 - [NicDevice](NicDevice.md)
 - [NicDeviceParameter](NicDeviceParameter.md)
 - [NicDevices](NicDevices.md)
 - [NicIndex](NicIndex.md)
 - [NicIpStack](NicIpStack.md)
 - [NicIpStackAll](NicIpStackAll.md)
 - [NicNumber](NicNumber.md)
 - [Number](Number.md)
 - [Port](Port.md)
 - [Portforward](Portforward.md)
 - [PortforwardGuest](PortforwardGuest.md)
 - [PortforwardParameter](PortforwardParameter.md)
 - [Portforwards](Portforwards.md)
 - [RouteEntry](RouteEntry.md)
 - [SharedFolder](SharedFolder.md)
 - [SharedFolderParameter](SharedFolderParameter.md)
 - [SharedFolders](SharedFolders.md)
 - [VmApplianceView](VmApplianceView.md)
 - [VmCloneParameter](VmCloneParameter.md)
 - [VmConnectedDevice](VmConnectedDevice.md)
 - [VmConnectedDeviceList](VmConnectedDeviceList.md)
 - [VmGuestIsolation](VmGuestIsolation.md)
 - [VmInformation](VmInformation.md)
 - [VmMemory](VmMemory.md)
 - [VmParameter](VmParameter.md)
 - [VmPowerOperation](VmPowerOperation.md)
 - [VmPowerState](VmPowerState.md)
 - [VmProcessors](VmProcessors.md)
 - [VmRegisterParameter](VmRegisterParameter.md)
 - [VmRemoteVnc](VmRemoteVnc.md)
 - [VmRestrictionsInformation](VmRestrictionsInformation.md)
 - [VmRrgistrationInformation](VmRrgistrationInformation.md)
 - [VmUsbDevice](VmUsbDevice.md)
 - [VmUsbList](VmUsbList.md)
 - [Vmcpu](Vmcpu.md)
 - [Vmid](Vmid.md)
 - [WinsConfig](WinsConfig.md)


## Documentation For Authorization

## BasicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
<!-- #TODO(Jonathan) convert to powershell -->