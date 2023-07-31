# Table: aws_ec2_vpc_endpoint_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| acceptance_required | bool | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| supported_ip_address_types | string_array | X | √ |  | 
| vpc_endpoint_policy_supported | bool | X | √ |  | 
| manages_vpc_endpoints | bool | X | √ |  | 
| owner | string | X | √ |  | 
| service_type | json | X | √ |  | 
| account_id | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| private_dns_name_verification_state | string | X | √ |  | 
| service_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| payer_responsibility | string | X | √ |  | 
| private_dns_names | json | X | √ |  | 
| service_id | string | X | √ |  | 


