# Table: aws_elbv2_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| scheme | string | X | √ |  | 
| type | string | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_time | timestamp | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| availability_zones | json | X | √ |  | 
| canonical_hosted_zone_id | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| state | json | X | √ |  | 
| dns_name | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


