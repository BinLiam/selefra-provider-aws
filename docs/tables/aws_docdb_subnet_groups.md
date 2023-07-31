# Table: aws_docdb_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_subnet_group_description | string | X | √ |  | 
| subnet_group_status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_subnet_group_arn | string | X | √ |  | 
| db_subnet_group_name | string | X | √ |  | 
| subnets | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 


