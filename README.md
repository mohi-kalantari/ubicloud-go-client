# Go API client for openapi

API for managing resources on Ubicloud

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.ubicloud.com/](https://www.ubicloud.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/mohi-kalantari/ubicloud-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ubicloud.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FirewallApi* | [**ActionLocationFirewallAttachSubnet**](docs/FirewallApi.md#actionlocationfirewallattachsubnet) | **Post** /project/{project_id}/location/{location}/firewall/_{firewall_id}/attach-subnet | Attach a subnet to firewall
*FirewallApi* | [**ActionLocationFirewallDetachSubnet**](docs/FirewallApi.md#actionlocationfirewalldetachsubnet) | **Post** /project/{project_id}/location/{location}/firewall/_{firewall_id}/detach-subnet | Detach a subnet from firewall
*FirewallApi* | [**CreateFirewall**](docs/FirewallApi.md#createfirewall) | **Post** /project/{project_id}/firewall | Create a new firewall
*FirewallApi* | [**CreateLocationFirewall**](docs/FirewallApi.md#createlocationfirewall) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name} | Create a new firewall
*FirewallApi* | [**DeleteFirewall**](docs/FirewallApi.md#deletefirewall) | **Delete** /project/{project_id}/firewall/{firewall_name} | Delete a specific firewall
*FirewallApi* | [**DeleteLocationFirewall**](docs/FirewallApi.md#deletelocationfirewall) | **Delete** /project/{project_id}/location/{location}/firewall/{firewall_name} | Delete a specific firewall
*FirewallApi* | [**DeleteLocationFirewallWithId**](docs/FirewallApi.md#deletelocationfirewallwithid) | **Delete** /project/{project_id}/location/{location}/firewall/id/{firewall_id} | Delete a specific firewall
*FirewallApi* | [**GetFirewall**](docs/FirewallApi.md#getfirewall) | **Get** /project/{project_id}/firewall | Return the list of firewalls in the project
*FirewallApi* | [**GetFirewallDetails**](docs/FirewallApi.md#getfirewalldetails) | **Get** /project/{project_id}/firewall/{firewall_name} | Get details of a specific firewall
*FirewallApi* | [**GetLocationFirewall**](docs/FirewallApi.md#getlocationfirewall) | **Get** /project/{project_id}/location/{location}/firewall | Return the list of firewalls in the project and location
*FirewallApi* | [**GetLocationFirewallDetails**](docs/FirewallApi.md#getlocationfirewalldetails) | **Get** /project/{project_id}/location/{location}/firewall/{firewall_name} | Get details of a specific firewall
*FirewallApi* | [**GetLocationFirewallDetailsWithId**](docs/FirewallApi.md#getlocationfirewalldetailswithid) | **Get** /project/{project_id}/location/{location}/firewall/id/{firewall_id} | Get details of a specific firewall
*FirewallRuleApi* | [**CreateFirewallRule**](docs/FirewallRuleApi.md#createfirewallrule) | **Post** /project/{project_id}/firewall/{firewall_name}/firewall-rule | Create a new firewall rule
*FirewallRuleApi* | [**CreateLocationFirewallFirewallRule**](docs/FirewallRuleApi.md#createlocationfirewallfirewallrule) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Create a new firewall rule
*FirewallRuleApi* | [**CreateLocationFirewallFirewallRuleWithId**](docs/FirewallRuleApi.md#createlocationfirewallfirewallrulewithid) | **Post** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Create a new firewall rule
*FirewallRuleApi* | [**CreateLocationFirewallRule**](docs/FirewallRuleApi.md#createlocationfirewallrule) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule | Create a new firewall rule
*FirewallRuleApi* | [**CreateLocationFirewallRuleWithId**](docs/FirewallRuleApi.md#createlocationfirewallrulewithid) | **Post** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule | Create a new firewall rule
*FirewallRuleApi* | [**DeleteFirewallRule**](docs/FirewallRuleApi.md#deletefirewallrule) | **Delete** /project/{project_id}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
*FirewallRuleApi* | [**DeleteLocationFirewallFirewallRule**](docs/FirewallRuleApi.md#deletelocationfirewallfirewallrule) | **Delete** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
*FirewallRuleApi* | [**DeleteLocationFirewallFirewallRuleWithId**](docs/FirewallRuleApi.md#deletelocationfirewallfirewallrulewithid) | **Delete** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
*FirewallRuleApi* | [**DeleteLocationPostgresFirewallRule**](docs/FirewallRuleApi.md#deletelocationpostgresfirewallrule) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
*FirewallRuleApi* | [**DeleteLocationPostgresFirewallRuleWithId**](docs/FirewallRuleApi.md#deletelocationpostgresfirewallrulewithid) | **Delete** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Delete a specific Postgres firewall rule
*FirewallRuleApi* | [**GetFirewallRuleDetails**](docs/FirewallRuleApi.md#getfirewallruledetails) | **Get** /project/{project_id}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule
*FirewallRuleApi* | [**GetLocationFirewallFirewallRuleDetails**](docs/FirewallRuleApi.md#getlocationfirewallfirewallruledetails) | **Get** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule
*FirewallRuleApi* | [**GetLocationFirewallFirewallRuleDetailsWithId**](docs/FirewallRuleApi.md#getlocationfirewallfirewallruledetailswithid) | **Get** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule
*KubernetesClusterApi* | [**CreateLocationKubernetesCluster**](docs/KubernetesClusterApi.md#createlocationkubernetescluster) | **Post** /project/{project_id}/location/{location}/kubernetes-cluster/{kubernetes_cluster_name} | Create Kubernetes Cluster in a specific location of a project
*KubernetesClusterApi* | [**GetKubernetesClusterDetails**](docs/KubernetesClusterApi.md#getkubernetesclusterdetails) | **Get** /project/{project_id}/location/{location}/kubernetes-cluster/{kubernetes_cluster_name} | Get details of a specific Kubernetes Cluster in a location
*KubernetesClusterApi* | [**ListLocationKubernetesClusters**](docs/KubernetesClusterApi.md#listlocationkubernetesclusters) | **Get** /project/{project_id}/location/{location}/kubernetes-cluster | List kubernetes clusters in a specific location of a project
*KubernetesVMApi* | [**CreateLocationKubernetesVM**](docs/KubernetesVMApi.md#createlocationkubernetesvm) | **Post** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Create Kubernetes VM in a specific location of a project
*KubernetesVMApi* | [**GetKubernetesVMDetails**](docs/KubernetesVMApi.md#getkubernetesvmdetails) | **Get** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Get details of a specific Kubernetes VM in a location
*KubernetesVMApi* | [**ListLocationKubernetesVMs**](docs/KubernetesVMApi.md#listlocationkubernetesvms) | **Get** /project/{project_id}/location/{location}/kubernetes-vm | List Kubernetes VMs in a specific location of a project
*LoadBalancerApi* | [**AttachVmLocationLoadBalancer**](docs/LoadBalancerApi.md#attachvmlocationloadbalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name}/attach-vm | Attach a VM to a Load Balancer in a specific location of a project
*LoadBalancerApi* | [**CreateLoadBalancer**](docs/LoadBalancerApi.md#createloadbalancer) | **Post** /project/{project_id}/load-balancer/{load_balancer_name} | Create a new Load Balancer in a project
*LoadBalancerApi* | [**CreateLocationLoadBalancer**](docs/LoadBalancerApi.md#createlocationloadbalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Create a new Load Balancer in a specific location of a project
*LoadBalancerApi* | [**DeleteLoadBalancer**](docs/LoadBalancerApi.md#deleteloadbalancer) | **Delete** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Delete a specific Load Balancer
*LoadBalancerApi* | [**DeleteLoadBalancerWithID**](docs/LoadBalancerApi.md#deleteloadbalancerwithid) | **Delete** /project/{project_id}/location/{location}/load-balancer/id/{load_balancer_id} | Delete a specific Load Balancer with ID
*LoadBalancerApi* | [**DetachVmLocationLoadBalancer**](docs/LoadBalancerApi.md#detachvmlocationloadbalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name}/detach-vm | Detach a VM from a Load Balancer in a specific location of a project
*LoadBalancerApi* | [**GetLoadBalancer**](docs/LoadBalancerApi.md#getloadbalancer) | **Get** /project/{project_id}/load-balancer/{load_balancer_name} | Get details of a specific Load Balancer
*LoadBalancerApi* | [**GetLoadBalancerDetails**](docs/LoadBalancerApi.md#getloadbalancerdetails) | **Get** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Get details of a specific Load Balancer in a location
*LoadBalancerApi* | [**GetLoadBalancerDetailsWithId**](docs/LoadBalancerApi.md#getloadbalancerdetailswithid) | **Get** /project/{project_id}/location/{location}/load-balancer/id/{load_balancer_id} | Get details of a specific Load Balancer in a location with ID
*LoadBalancerApi* | [**ListLoadBalancers**](docs/LoadBalancerApi.md#listloadbalancers) | **Get** /project/{project_id}/load-balancer | List Load Balancers in a specific project
*LoadBalancerApi* | [**ListLocationLoadBalancers**](docs/LoadBalancerApi.md#listlocationloadbalancers) | **Get** /project/{project_id}/location/{location}/load-balancer | List Load Balancers in a specific location of a project
*LoadBalancerApi* | [**PatchLocationLoadBalancer**](docs/LoadBalancerApi.md#patchlocationloadbalancer) | **Patch** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Update a Load Balancer in a specific location of a project
*LoginApi* | [**Login**](docs/LoginApi.md#login) | **Post** /login | Login with user information
*PostgresApi* | [**CreatePostgresDatabase**](docs/PostgresApi.md#createpostgresdatabase) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Create a new Postgres database in a specific location of a project
*PostgresApi* | [**DeletePostgresDatabase**](docs/PostgresApi.md#deletepostgresdatabase) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Delete a specific Postgres database
*PostgresApi* | [**DeletePostgresDatabaseWithID**](docs/PostgresApi.md#deletepostgresdatabasewithid) | **Delete** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id} | Delete a specific Postgres database with ID
*PostgresApi* | [**FailoverPostgresDatabaseWithID**](docs/PostgresApi.md#failoverpostgresdatabasewithid) | **Post** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/failover | Failover a specific Postgres database with ID
*PostgresApi* | [**GetPostgresDatabaseDetails**](docs/PostgresApi.md#getpostgresdatabasedetails) | **Get** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Get details of a specific Postgres database in a location
*PostgresApi* | [**GetPostgresDetailsWithId**](docs/PostgresApi.md#getpostgresdetailswithid) | **Get** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id} | Get details of a specific Postgres database in a location with ID
*PostgresApi* | [**ListLocationPostgresDatabases**](docs/PostgresApi.md#listlocationpostgresdatabases) | **Get** /project/{project_id}/location/{location}/postgres | List Postgres databases in a specific location of a project
*PostgresApi* | [**ListPostgresDatabases**](docs/PostgresApi.md#listpostgresdatabases) | **Get** /project/{project_id}/postgres | List visible Postgres databases
*PostgresApi* | [**ResetSuperuserPassword**](docs/PostgresApi.md#resetsuperuserpassword) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/reset-superuser-password | Reset superuser password of the Postgres database
*PostgresApi* | [**ResetSuperuserPasswordWithID**](docs/PostgresApi.md#resetsuperuserpasswordwithid) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/reset-superuser-password | Reset super-user password of the Postgres database
*PostgresApi* | [**RestorePostgresDatabase**](docs/PostgresApi.md#restorepostgresdatabase) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/restore | Restore a new Postgres database in a specific location of a project
*PostgresApi* | [**RestorePostgresDatabaseWithID**](docs/PostgresApi.md#restorepostgresdatabasewithid) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/restore | Restore a new Postgres database in a specific location of a project with ID
*PostgresFirewallRuleApi* | [**CreateLocationPostgresFirewallRule**](docs/PostgresFirewallRuleApi.md#createlocationpostgresfirewallrule) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/firewall-rule | Create a new postgres firewall rule
*PostgresFirewallRuleApi* | [**CreateLocationPostgresFirewallRuleWithId**](docs/PostgresFirewallRuleApi.md#createlocationpostgresfirewallrulewithid) | **Post** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule | Create a new Postgres firewall rule
*PostgresFirewallRuleApi* | [**CreateLocationPostgresFirewallRuleWithIdWithId**](docs/PostgresFirewallRuleApi.md#createlocationpostgresfirewallrulewithidwithid) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Create a new Postgres firewall rule
*PostgresFirewallRuleApi* | [**GetLocationPostgresFirewallRuleDetailsWithId**](docs/PostgresFirewallRuleApi.md#getlocationpostgresfirewallruledetailswithid) | **Get** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Get details of a Postgres firewall rule
*PostgresFirewallRuleApi* | [**ListLocationPostgresFirewallRules**](docs/PostgresFirewallRuleApi.md#listlocationpostgresfirewallrules) | **Get** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule | List location Postgres firewall rules
*PostgresMetricDestinationApi* | [**CreateLocationPostgresMetricDestination**](docs/PostgresMetricDestinationApi.md#createlocationpostgresmetricdestination) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination | Create a new Postgres Metric Destination
*PostgresMetricDestinationApi* | [**DeleteLocationPostgresMetricDestination**](docs/PostgresMetricDestinationApi.md#deletelocationpostgresmetricdestination) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination/{metric_destination_id} | Delete a specific Metric Destination
*PrivateSubnetApi* | [**CreatePrivateSubnet**](docs/PrivateSubnetApi.md#createprivatesubnet) | **Post** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Create a new Private Subnet in a specific location of a project
*PrivateSubnetApi* | [**DeletePSWithId**](docs/PrivateSubnetApi.md#deletepswithid) | **Delete** /project/{project_id}/location/{location}/private-subnet/id/{private_subnet_id} | Delete a specific Private Subnet with ID
*PrivateSubnetApi* | [**DeletePrivateSubnet**](docs/PrivateSubnetApi.md#deleteprivatesubnet) | **Delete** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Delete a specific Private Subnet
*PrivateSubnetApi* | [**GetPSDetailsWithId**](docs/PrivateSubnetApi.md#getpsdetailswithid) | **Get** /project/{project_id}/location/{location}/private-subnet/id/{private_subnet_id} | Get details of a specific Private Subnet in a location with ID
*PrivateSubnetApi* | [**GetPrivateSubnetDetails**](docs/PrivateSubnetApi.md#getprivatesubnetdetails) | **Get** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Get details of a specific Private Subnet in a location
*PrivateSubnetApi* | [**ListLocationPrivateSubnets**](docs/PrivateSubnetApi.md#listlocationprivatesubnets) | **Get** /project/{project_id}/location/{location}/private-subnet | List Private Subnets in a specific location of a project
*PrivateSubnetApi* | [**ListPSs**](docs/PrivateSubnetApi.md#listpss) | **Get** /project/{project_id}/private-subnet | List visible Private Subnets
*ProjectApi* | [**CreateProject**](docs/ProjectApi.md#createproject) | **Post** /project | Create a new project
*ProjectApi* | [**DeleteProject**](docs/ProjectApi.md#deleteproject) | **Delete** /project/{project_id} | Delete a project
*ProjectApi* | [**GetProject**](docs/ProjectApi.md#getproject) | **Get** /project/{project_id} | Retrieve a project
*ProjectApi* | [**ListProjects**](docs/ProjectApi.md#listprojects) | **Get** /project | List all projects visible to the logged in user.
*VirtualMachineApi* | [**CreateVM**](docs/VirtualMachineApi.md#createvm) | **Post** /project/{project_id}/location/{location}/vm/{vm_name} | Create a new VM in a specific location of a project
*VirtualMachineApi* | [**DeleteVM**](docs/VirtualMachineApi.md#deletevm) | **Delete** /project/{project_id}/location/{location}/vm/{vm_name} | Delete a specific VM
*VirtualMachineApi* | [**DeleteVMWithId**](docs/VirtualMachineApi.md#deletevmwithid) | **Delete** /project/{project_id}/location/{location}/vm/id/{vm_id} | Delete a specific VM with ID
*VirtualMachineApi* | [**GetVMDetails**](docs/VirtualMachineApi.md#getvmdetails) | **Get** /project/{project_id}/location/{location}/vm/{vm_name} | Get details of a specific VM in a location
*VirtualMachineApi* | [**GetVMDetailsWithId**](docs/VirtualMachineApi.md#getvmdetailswithid) | **Get** /project/{project_id}/location/{location}/vm/id/{vm_id} | Get details of a specific VM in a location with ID
*VirtualMachineApi* | [**ListLocationVMs**](docs/VirtualMachineApi.md#listlocationvms) | **Get** /project/{project_id}/location/{location}/vm | List VMs in a specific location of a project
*VirtualMachineApi* | [**ListProjectVMs**](docs/VirtualMachineApi.md#listprojectvms) | **Get** /project/{project_id}/vm | List all VMs created under the given project ID and visible to logged in user


## Documentation For Models

 - [ActionLocationFirewallAttachSubnetRequest](docs/ActionLocationFirewallAttachSubnetRequest.md)
 - [ActionLocationFirewallDetachSubnetRequest](docs/ActionLocationFirewallDetachSubnetRequest.md)
 - [AttachVmLocationLoadBalancerRequest](docs/AttachVmLocationLoadBalancerRequest.md)
 - [CreateFirewallRequest](docs/CreateFirewallRequest.md)
 - [CreateFirewallRuleRequest](docs/CreateFirewallRuleRequest.md)
 - [CreateLoadBalancerRequest](docs/CreateLoadBalancerRequest.md)
 - [CreateLocationFirewallRequest](docs/CreateLocationFirewallRequest.md)
 - [CreateLocationKubernetesClusterRequest](docs/CreateLocationKubernetesClusterRequest.md)
 - [CreateLocationKubernetesVMRequest](docs/CreateLocationKubernetesVMRequest.md)
 - [CreateLocationPostgresFirewallRuleWithIdRequest](docs/CreateLocationPostgresFirewallRuleWithIdRequest.md)
 - [CreateLocationPostgresFirewallRuleWithIdWithIdRequest](docs/CreateLocationPostgresFirewallRuleWithIdWithIdRequest.md)
 - [CreateLocationPostgresMetricDestinationRequest](docs/CreateLocationPostgresMetricDestinationRequest.md)
 - [CreatePostgresDatabaseRequest](docs/CreatePostgresDatabaseRequest.md)
 - [CreatePrivateSubnetRequest](docs/CreatePrivateSubnetRequest.md)
 - [CreateProjectRequest](docs/CreateProjectRequest.md)
 - [CreateVMRequest](docs/CreateVMRequest.md)
 - [DetachVmLocationLoadBalancerRequest](docs/DetachVmLocationLoadBalancerRequest.md)
 - [Error](docs/Error.md)
 - [ErrorError](docs/ErrorError.md)
 - [Firewall](docs/Firewall.md)
 - [FirewallDetailed](docs/FirewallDetailed.md)
 - [FirewallDetailedAllOf](docs/FirewallDetailedAllOf.md)
 - [FirewallRule](docs/FirewallRule.md)
 - [GetFirewall200Response](docs/GetFirewall200Response.md)
 - [KubernetesCluster](docs/KubernetesCluster.md)
 - [ListLoadBalancers200Response](docs/ListLoadBalancers200Response.md)
 - [ListLocationKubernetesClusters200Response](docs/ListLocationKubernetesClusters200Response.md)
 - [ListLocationPostgresDatabases200Response](docs/ListLocationPostgresDatabases200Response.md)
 - [ListLocationPostgresFirewallRules200Response](docs/ListLocationPostgresFirewallRules200Response.md)
 - [ListLocationPrivateSubnets200Response](docs/ListLocationPrivateSubnets200Response.md)
 - [ListLocationVMs200Response](docs/ListLocationVMs200Response.md)
 - [ListProjects200Response](docs/ListProjects200Response.md)
 - [LoadBalancer](docs/LoadBalancer.md)
 - [LoadBalancerDetailed](docs/LoadBalancerDetailed.md)
 - [LoadBalancerDetailedAllOf](docs/LoadBalancerDetailedAllOf.md)
 - [Login200Response](docs/Login200Response.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [Nic](docs/Nic.md)
 - [PatchLocationLoadBalancerRequest](docs/PatchLocationLoadBalancerRequest.md)
 - [Postgres](docs/Postgres.md)
 - [PostgresDetailed](docs/PostgresDetailed.md)
 - [PostgresDetailedAllOf](docs/PostgresDetailedAllOf.md)
 - [PostgresFirewallRule](docs/PostgresFirewallRule.md)
 - [PrivateSubnet](docs/PrivateSubnet.md)
 - [Project](docs/Project.md)
 - [ResetSuperuserPasswordWithIDRequest](docs/ResetSuperuserPasswordWithIDRequest.md)
 - [RestorePostgresDatabaseWithIDRequest](docs/RestorePostgresDatabaseWithIDRequest.md)
 - [Vm](docs/Vm.md)
 - [VmDetailed](docs/VmDetailed.md)
 - [VmDetailedAllOf](docs/VmDetailedAllOf.md)


## Documentation For Authorization



### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@ubicloud.com

