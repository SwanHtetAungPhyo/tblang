# TBLang Loop Guide

## For Loops

TBLang supports `for` loops to iterate over collections and create multiple resources dynamically.

### Syntax

```tblang
for iterator in collection {
    // statements
}
```

### Basic Examples

#### Loop over Array

```tblang
declare availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"];

for az in availability_zones {
    declare subnet = subnet("subnet-${az}", {
        availability_zone: az
        cidr_block: "10.0.${index}.0/24"
    });
}
```

#### Loop over Object Array

```tblang
declare subnet_configs = [
    {
        name: "public-subnet-1"
        cidr: "10.0.1.0/24"
        type: "public"
    },
    {
        name: "private-subnet-1"
        cidr: "10.0.10.0/24"
        type: "private"
    }
];

for config in subnet_configs {
    declare subnet = subnet(config.name, {
        cidr_block: config.cidr
        map_public_ip: config.type == "public"
        tags: {
            Name: config.name
            Type: config.type
        }
    });
}
```

### Real-World Examples

#### Create Multiple Subnets

```tblang
cloud_vendor "aws" {
    region = "us-east-1"
    profile = "default"
}

declare main_vpc = vpc("main-vpc", {
    cidr_block: "10.0.0.0/16"
});

declare subnets = [
    { name: "web-subnet", cidr: "10.0.1.0/24", az: "us-east-1a" },
    { name: "app-subnet", cidr: "10.0.2.0/24", az: "us-east-1b" },
    { name: "db-subnet", cidr: "10.0.3.0/24", az: "us-east-1c" }
];

for subnet_config in subnets {
    declare subnet = subnet(subnet_config.name, {
        vpc_id: main_vpc
        cidr_block: subnet_config.cidr
        availability_zone: subnet_config.az
        tags: {
            Name: subnet_config.name
        }
    });
}
```

#### Create Security Groups for Multiple Environments

```tblang
declare environments = ["dev", "staging", "prod"];

for env in environments {
    declare sg = security_group("${env}-sg", {
        vpc_id: main_vpc
        name: "${env}-security-group"
        description: "Security group for ${env}"
        ingress_rules: [
            {
                protocol: "tcp"
                from_port: 80
                to_port: 80
                cidr_blocks: ["0.0.0.0/0"]
            }
        ]
        tags: {
            Name: "${env}-sg"
            Environment: env
        }
    });
}
```

#### Create Multiple VPCs

```tblang
declare regions = [
    { name: "us-east", region: "us-east-1", cidr: "10.0.0.0/16" },
    { name: "us-west", region: "us-west-2", cidr: "10.1.0.0/16" },
    { name: "eu-west", region: "eu-west-1", cidr: "10.2.0.0/16" }
];

for region_config in regions {
    declare vpc = vpc("${region_config.name}-vpc", {
        cidr_block: region_config.cidr
        tags: {
            Name: "${region_config.name}-vpc"
            Region: region_config.region
        }
    });
}
```

### Loop Features

#### Access Iterator Value

```tblang
for item in collection {
    // 'item' contains the current element
    declare resource = resource_type(item.name, {
        property: item.value
    });
}
```

#### Nested Loops

```tblang
declare environments = ["dev", "prod"];
declare availability_zones = ["us-east-1a", "us-east-1b"];

for env in environments {
    for az in availability_zones {
        declare subnet = subnet("${env}-${az}", {
            availability_zone: az
            tags: {
                Environment: env
                AZ: az
            }
        });
    }
}
```

### Best Practices

1. **Use Descriptive Iterator Names**
   ```tblang
   // Good
   for subnet_config in subnet_configs { }
   
   // Avoid
   for x in list { }
   ```

2. **Keep Loop Bodies Simple**
   ```tblang
   // Good - one resource per loop
   for config in configs {
       declare resource = resource_type(config.name, config);
   }
   ```

3. **Use Loops for Repetitive Resources**
   ```tblang
   // Instead of:
   declare subnet1 = subnet("subnet-1", {...});
   declare subnet2 = subnet("subnet-2", {...});
   declare subnet3 = subnet("subnet-3", {...});
   
   // Use:
   for config in subnet_configs {
       declare subnet = subnet(config.name, config);
   }
   ```

4. **Combine with Variables**
   ```tblang
   declare base_config = {
       enable_dns: true
       enable_monitoring: true
   };
   
   for name in resource_names {
       declare resource = resource_type(name, base_config);
   }
   ```

### Limitations

- Loops are evaluated at compile time
- Cannot use loop variables outside the loop scope
- Loop collections must be defined before use

### Complete Example

See [loop-example.tbl](tblang-demo/loop-example.tbl) for a complete working example.

### Future Enhancements

Coming soon:
- Range loops: `for i in range(0, 10) { }`
- Index access: `for index, item in collection { }`
- Conditional loops: `for item in collection if condition { }`
- Map loops: `for key, value in map { }`
