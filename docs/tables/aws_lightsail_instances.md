# Table: aws_lightsail_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ipv6_addresses | string_array | X | √ |  | 
| location | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | json | X | √ |  | 
| username | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| ip_address_type | string | X | √ |  | 
| metadata_options | json | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| ssh_key_name | string | X | √ |  | 
| bundle_id | string | X | √ |  | 
| hardware | json | X | √ |  | 
| networking | json | X | √ |  | 
| support_code | string | X | √ |  | 
| tags | json | X | √ |  | 
| access_details | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| blueprint_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| is_static_ip | bool | X | √ |  | 
| name | string | X | √ |  | 
| public_ip_address | string | X | √ |  | 
| region | string | X | √ |  | 
| add_ons | json | X | √ |  | 
| blueprint_name | string | X | √ |  | 


