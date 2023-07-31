# Table: aws_rds_cluster_parameter_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| db_cluster_parameter_group_arn | string | X | √ |  | 
| db_cluster_parameter_group_name | string | X | √ |  | 
| description | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


