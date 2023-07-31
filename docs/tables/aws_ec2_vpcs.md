# Table: aws_ec2_vpcs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| cidr_block | string | X | √ |  | 
| dhcp_options_id | string | X | √ |  | 
| is_default | bool | X | √ |  | 
| owner_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cidr_block_association_set | json | X | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 
| instance_tenancy | string | X | √ |  | 
| region | string | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 


