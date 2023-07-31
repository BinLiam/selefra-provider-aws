# Table: aws_docdb_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| account_id | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_retention_period | big_int | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| percent_progress | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| db_cluster_arn | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| region | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| engine | string | X | √ |  | 
| port | big_int | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| engine_version | string | X | √ |  | 


