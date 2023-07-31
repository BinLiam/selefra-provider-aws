# Table: aws_ec2_vpc_endpoint_service_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service_name | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| payer_responsibility | string | X | √ |  | 
| private_dns_name_configuration | json | X | √ |  | 
| service_id | string | X | √ |  | 
| acceptance_required | bool | X | √ |  | 
| service_state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| gateway_load_balancer_arns | string_array | X | √ |  | 
| supported_ip_address_types | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| manages_vpc_endpoints | bool | X | √ |  | 
| network_load_balancer_arns | string_array | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| service_type | json | X | √ |  | 
| region | string | X | √ |  | 


