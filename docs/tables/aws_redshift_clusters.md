# Table: aws_redshift_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_parameter_groups | json | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| iam_roles | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| tags | json | X | √ |  | 
| encrypted | bool | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| expected_next_snapshot_schedule_time_status | string | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| deferred_maintenance_windows | json | X | √ |  | 
| elastic_resize_number_of_node_options | string | X | √ |  | 
| snapshot_schedule_identifier | string | X | √ |  | 
| cluster_revision_number | string | X | √ |  | 
| default_iam_role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| pending_actions | string_array | X | √ |  | 
| allow_version_upgrade | bool | X | √ |  | 
| cluster_availability_status | string | X | √ |  | 
| cluster_snapshot_copy_status | json | X | √ |  | 
| elastic_ip_status | json | X | √ |  | 
| expected_next_snapshot_schedule_time | timestamp | X | √ |  | 
| automated_snapshot_retention_period | big_int | X | √ |  | 
| master_username | string | X | √ |  | 
| region | string | X | √ |  | 
| aqua_configuration | json | X | √ |  | 
| total_storage_capacity_in_mega_bytes | big_int | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| pending_modified_values | json | X | √ |  | 
| cluster_public_key | string | X | √ |  | 
| cluster_status | string | X | √ |  | 
| cluster_version | string | X | √ |  | 
| data_transfer_progress | json | X | √ |  | 
| hsm_status | json | X | √ |  | 
| node_type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| cluster_subnet_group_name | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| cluster_nodes | json | X | √ |  | 
| availability_zone_relocation_status | string | X | √ |  | 
| cluster_namespace_arn | string | X | √ |  | 
| cluster_security_groups | json | X | √ |  | 
| db_name | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| restore_status | json | X | √ |  | 
| manual_snapshot_retention_period | big_int | X | √ |  | 
| next_maintenance_window_start_time | timestamp | X | √ |  | 
| reserved_node_exchange_status | json | X | √ |  | 
| snapshot_schedule_state | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| modify_status | string | X | √ |  | 
| resize_info | json | X | √ |  | 
| logging_status | json | X | √ | `Describes the status of logging for a cluster.` | 


