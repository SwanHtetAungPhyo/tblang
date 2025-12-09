package provider

// This file serves as the main entry point for the AWS client package.
// All functionality has been split into separate files for better organization:
//
// Client Core:
// - client_types.go: Type definitions for AWSClient and result types
// - client_constructor.go: Client constructor
// - client_helpers.go: Helper functions (buildTags, etc.)
//
// Resource-specific clients:
// - client_vpc.go: VPC operations (Create, Delete, EnableDNS)
// - client_subnet.go: Subnet operations (Create, Delete, ConfigurePublicIP)
// - client_security_group.go: Security Group operations (Create, Delete, Authorize rules)
// - client_ec2.go: EC2 instance operations (Create, Terminate)
// - client_internet_gateway.go: Internet Gateway operations (Create, Delete, Attach/Detach)
// - client_route_table.go: Route Table operations (Create, Delete, CreateRoute)
// - client_eip.go: Elastic IP operations (Allocate, Release)
// - client_nat_gateway.go: NAT Gateway operations (Create, Delete)
// - client_datasources.go: Data source operations (Describe AMI, VPC, Subnet, AZs, CallerIdentity)
