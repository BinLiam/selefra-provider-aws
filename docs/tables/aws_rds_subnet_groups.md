# Table: aws_rds_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subnet_group_status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| db_subnet_group_arn | string | X | √ |  | 
| db_subnet_group_description | string | X | √ |  | 
| db_subnet_group_name | string | X | √ |  | 
| subnets | json | X | √ |  | 
| supported_network_types | string_array | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 


