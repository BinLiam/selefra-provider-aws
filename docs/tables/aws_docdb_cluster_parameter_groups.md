# Table: aws_docdb_cluster_parameter_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_cluster_parameter_group_name | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_cluster_parameter_group_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| parameters | json | X | √ |  | 
| description | string | X | √ |  | 


