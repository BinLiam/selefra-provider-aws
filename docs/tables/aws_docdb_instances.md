# Table: aws_docdb_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| availability_zone | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| engine | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_retention_period | big_int | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| promotion_tier | big_int | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| status_infos | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_instance_identifier | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| tags | json | X | √ |  | 
| aws_docdb_clusters_selefra_id | string | X | X | fk to aws_docdb_clusters.selefra_id | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| arn | string | √ | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| db_instance_arn | string | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| db_instance_status | string | X | √ |  | 


