# Table: aws_neptune_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_parameter_groups | json | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| allocated_storage | big_int | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| account_id | string | X | √ |  | 
| timezone | string | X | √ |  | 
| db_instance_port | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_retention_period | big_int | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| engine | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| db_instance_arn | string | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| character_set_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| endpoint | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| monitoring_interval | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| db_name | string | X | √ |  | 
| iops | big_int | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| status_infos | json | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| availability_zone | string | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| license_model | string | X | √ |  | 


