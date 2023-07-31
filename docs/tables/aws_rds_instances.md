# Table: aws_rds_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| preferred_backup_window | string | X | √ |  | 
| activity_stream_policy_status | string | X | √ |  | 
| character_set_name | string | X | √ |  | 
| master_username | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| processor_features | json | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| aws_backup_recovery_point_arn | string | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| storage_type | string | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| db_system_id | string | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| custom_iam_instance_profile | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| region | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| db_instance_automated_backups_replications | json | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| activity_stream_engine_native_audit_fields_included | bool | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| replica_mode | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| nchar_character_set_name | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| arn | string | √ | √ |  | 
| activity_stream_status | string | X | √ |  | 
| db_instance_arn | string | X | √ |  | 
| db_parameter_groups | json | X | √ |  | 
| engine | string | X | √ |  | 
| license_model | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| network_type | string | X | √ |  | 
| timezone | string | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| customer_owned_ip_enabled | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| resume_full_automation_mode_time | timestamp | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| automation_mode | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| tag_list | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| max_allocated_storage | big_int | X | √ |  | 
| performance_insights_retention_period | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| backup_target | string | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| db_instance_port | big_int | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| certificate_details | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| listener_endpoint | json | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| db_name | string | X | √ |  | 
| master_user_secret | json | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| status_infos | json | X | √ |  | 
| storage_throughput | big_int | X | √ |  | 
| activity_stream_mode | string | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 


