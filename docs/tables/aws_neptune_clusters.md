# Table: aws_neptune_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| endpoint | string | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| database_name | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| engine | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| multi_az | bool | X | √ |  | 
| percent_progress | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| serverless_v2_scaling_configuration | json | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| master_username | string | X | √ |  | 
| port | big_int | X | √ |  | 
| character_set_name | string | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| cross_account_clone | bool | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| region | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| status | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


