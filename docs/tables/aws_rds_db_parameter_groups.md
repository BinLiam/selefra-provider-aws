# Table: aws_rds_db_parameter_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_parameter_group_family | string | X | √ |  | 
| db_parameter_group_name | string | X | √ |  | 
| description | string | X | √ |  | 
| db_parameter_group_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 


