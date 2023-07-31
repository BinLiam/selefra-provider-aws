# Table: aws_elasticache_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cache_subnet_group_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cache_cluster_id | string | X | √ |  | 
| cache_parameter_group_name | string | X | √ |  | 
| num_cache_nodes | big_int | X | √ |  | 
| num_node_groups | big_int | X | √ |  | 
| replication_group_description | string | X | √ |  | 
| snapshot_source | string | X | √ |  | 
| topic_arn | string | X | √ |  | 
| cache_cluster_create_time | timestamp | X | √ |  | 
| engine | string | X | √ |  | 
| node_snapshots | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| port | big_int | X | √ |  | 
| snapshot_name | string | X | √ |  | 
| preferred_availability_zone | string | X | √ |  | 
| preferred_outpost_arn | string | X | √ |  | 
| snapshot_status | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| data_tiering | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_retention_limit | big_int | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| automatic_failover | string | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 


