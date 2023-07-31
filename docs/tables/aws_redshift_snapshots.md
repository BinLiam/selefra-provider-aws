# Table: aws_redshift_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| port | big_int | X | √ |  | 
| restorable_node_types | string_array | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| elapsed_time_in_seconds | big_int | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| manual_snapshot_retention_period | big_int | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| snapshot_retention_start_time | timestamp | X | √ |  | 
| source_region | string | X | √ |  | 
| cluster_version | string | X | √ |  | 
| current_backup_rate_in_mega_bytes_per_second | float | X | √ |  | 
| encrypted_with_hsm | bool | X | √ |  | 
| engine_full_version | string | X | √ |  | 
| estimated_seconds_to_completion | big_int | X | √ |  | 
| node_type | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| encrypted | bool | X | √ |  | 
| manual_snapshot_remaining_days | big_int | X | √ |  | 
| master_username | string | X | √ |  | 
| status | string | X | √ |  | 
| backup_progress_in_mega_bytes | float | X | √ |  | 
| account_id | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| total_backup_size_in_mega_bytes | float | X | √ |  | 
| aws_redshift_clusters_selefra_id | string | X | X | fk to aws_redshift_clusters.selefra_id | 
| availability_zone | string | X | √ |  | 
| db_name | string | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| snapshot_identifier | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| arn | string | √ | √ | `ARN of the snapshot.` | 
| accounts_with_restore_access | json | X | √ |  | 
| actual_incremental_backup_size_in_mega_bytes | float | X | √ |  | 
| tags | json | X | √ | `Tags consisting of a name/value pair for a resource.` | 
| region | string | X | √ |  | 


