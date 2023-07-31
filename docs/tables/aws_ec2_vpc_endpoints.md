# Table: aws_ec2_vpc_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_id | string | X | √ |  | 
| policy_document | string | X | √ |  | 
| service_name | string | X | √ |  | 
| vpc_endpoint_type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| dns_entries | json | X | √ |  | 
| groups | json | X | √ |  | 
| last_error | json | X | √ |  | 
| private_dns_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| creation_timestamp | timestamp | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| vpc_endpoint_id | string | X | √ |  | 
| dns_options | json | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| route_table_ids | string_array | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


