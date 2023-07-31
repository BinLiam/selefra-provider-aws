# Table: aws_elasticache_replication_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cache_node_type | string | X | √ |  | 
| network_type | string | X | √ |  | 
| status | string | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| log_delivery_configurations | json | X | √ |  | 
| multi_az | string | X | √ |  | 
| auth_token_last_modified_date | timestamp | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| automatic_failover | string | X | √ |  | 
| cluster_enabled | bool | X | √ |  | 
| configuration_endpoint | json | X | √ |  | 
| account_id | string | X | √ |  | 
| snapshotting_cluster_id | string | X | √ |  | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| data_tiering | string | X | √ |  | 
| global_replication_group_info | json | X | √ |  | 
| member_clusters | string_array | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| ip_discovery | string | X | √ |  | 
| transit_encryption_mode | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auth_token_enabled | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| member_clusters_outpost_arns | string_array | X | √ |  | 
| snapshot_retention_limit | big_int | X | √ |  | 
| user_group_ids | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| node_groups | json | X | √ |  | 
| replication_group_create_time | timestamp | X | √ |  | 
| replication_group_id | string | X | √ |  | 


