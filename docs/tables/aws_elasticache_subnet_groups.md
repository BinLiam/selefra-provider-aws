# Table: aws_elasticache_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpc_id | string | X | √ |  | 
| subnets | json | X | √ |  | 
| cache_subnet_group_description | string | X | √ |  | 
| cache_subnet_group_name | string | X | √ |  | 
| supported_network_types | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 


