# Table: aws_ec2_network_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| network_interface_id | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| groups | json | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| ipv6_native | bool | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| status | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| association | json | X | √ |  | 
| mac_address | string | X | √ |  | 
| requester_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| private_ip_addresses | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| ipv4_prefixes | json | X | √ |  | 
| ipv6_addresses | json | X | √ |  | 
| ipv6_prefixes | json | X | √ |  | 
| deny_all_igw_traffic | bool | X | √ |  | 
| interface_type | string | X | √ |  | 
| tag_set | json | X | √ |  | 
| attachment | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| vpc_id | string | X | √ |  | 


