# Table: aws_neptune_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_subnet_group_name | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| db_subnet_group_arn | string | X | √ |  | 
| db_subnet_group_description | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subnet_group_status | string | X | √ |  | 
| subnets | json | X | √ |  | 


