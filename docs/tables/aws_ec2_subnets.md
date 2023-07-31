# Table: aws_ec2_subnets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| assign_ipv6_address_on_creation | bool | X | √ |  | 
| available_ip_address_count | big_int | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| enable_dns64 | bool | X | √ |  | 
| subnet_arn | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| cidr_block | string | X | √ |  | 
| map_customer_owned_ip_on_launch | bool | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| default_for_az | bool | X | √ |  | 
| enable_lni_at_device_index | big_int | X | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 
| ipv6_native | bool | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| map_public_ip_on_launch | bool | X | √ |  | 
| private_dns_name_options_on_launch | json | X | √ |  | 
| vpc_id | string | X | √ |  | 


