# Table: aws_rds_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| engine | string | X | √ |  | 
| engine_mode | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| master_username | string | X | √ |  | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| db_system_id | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| tag_list | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cluster_create_time | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| license_model | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| tags | json | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| port | big_int | X | √ |  | 
| status | string | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| region | string | X | √ |  | 


