# Table: aws_cloudhsmv2_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_message | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| account_id | string | X | √ |  | 
| backup_policy | string | X | √ |  | 
| certificates | json | X | √ |  | 
| security_group | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_retention_policy | json | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 
| cluster_id | string | X | √ |  | 
| hsm_type | string | X | √ |  | 
| state | string | X | √ |  | 
| subnet_mapping | json | X | √ |  | 
| hsms | json | X | √ |  | 
| pre_co_password | string | X | √ |  | 
| source_backup_id | string | X | √ |  | 


