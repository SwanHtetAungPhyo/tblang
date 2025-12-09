package provider

// This file serves as the main entry point for the AWS provider package.
// All functionality has been split into separate files for better organization:
//
// Provider Core:
// - types.go: Type definitions for AWSProvider
// - constructor.go: Provider constructor
// - configure.go: Provider configuration
// - schema.go: Schema definitions entry point
// - schema_resources.go: Resource schema definitions
// - schema_datasources.go: Data source schema definitions
// - interface.go: Provider interface methods (Plan, Read, Import, Validate)
// - apply.go: ApplyResourceChange dispatcher
// - helpers.go: Helper functions
//
// Resources (apply and destroy):
// - resource_vpc.go: VPC resource operations
// - resource_subnet.go: Subnet resource operations
// - resource_security_group.go: Security Group resource operations
// - resource_ec2.go: EC2 instance resource operations
// - resource_internet_gateway.go: Internet Gateway resource operations
// - resource_route_table.go: Route Table resource operations
// - resource_eip.go: Elastic IP resource operations
// - resource_nat_gateway.go: NAT Gateway resource operations
//
// Data Sources:
// - datasource_ami.go: AMI data source
// - datasource_vpc.go: VPC data source
// - datasource_subnet.go: Subnet data source
// - datasource_availability_zones.go: Availability Zones data source
// - datasource_caller_identity.go: Caller Identity data source
