# Table: aws_rds_db_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| storage_throughput | big_int | X | √ |  | 
| availability_zone | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| iops | big_int | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| option_group_name | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| attributes | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| db_snapshot_identifier | string | X | √ |  | 
| engine | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| timezone | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| port | big_int | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| original_snapshot_create_time | timestamp | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| snapshot_database_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| snapshot_target | string | X | √ |  | 
| source_region | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| db_snapshot_arn | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| license_model | string | X | √ |  | 
| processor_features | json | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| source_db_snapshot_identifier | string | X | √ |  | 
| arn | string | √ | √ |  | 


