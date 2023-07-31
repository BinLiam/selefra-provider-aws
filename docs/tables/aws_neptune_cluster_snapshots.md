# Table: aws_neptune_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| license_model | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| master_username | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| port | big_int | X | √ |  | 
| vpc_id | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 


