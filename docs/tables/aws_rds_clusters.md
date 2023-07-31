# Table: aws_rds_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| backup_retention_period | big_int | X | √ |  | 
| endpoint | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| network_type | string | X | √ |  | 
| percent_progress | string | X | √ |  | 
| custom_endpoints | string_array | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| backtrack_consumed_change_records | big_int | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| engine_mode | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| status | string | X | √ |  | 
| cross_account_clone | bool | X | √ |  | 
| db_cluster_instance_class | string | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| serverless_v2_scaling_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_system_id | string | X | √ |  | 
| earliest_backtrack_time | timestamp | X | √ |  | 
| master_username | string | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| global_write_forwarding_status | string | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| backtrack_window | big_int | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| activity_stream_mode | string | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| global_write_forwarding_requested | bool | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| region | string | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| database_name | string | X | √ |  | 
| performance_insights_retention_period | big_int | X | √ |  | 
| associated_roles | json | X | √ |  | 
| http_endpoint_enabled | bool | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| iops | big_int | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| master_user_secret | json | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| activity_stream_status | string | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| engine | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| capacity | big_int | X | √ |  | 
| character_set_name | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| port | big_int | X | √ |  | 
| scaling_configuration_info | json | X | √ |  | 
| tags | json | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 


