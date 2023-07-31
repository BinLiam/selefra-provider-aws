# Table: aws_docdb_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| engine | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| vpc_id | string | X | √ |  | 
| aws_docdb_clusters_selefra_id | string | X | X | fk to aws_docdb_clusters.selefra_id | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| master_username | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| attributes | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| percent_progress | big_int | X | √ |  | 
| port | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| db_cluster_snapshot_arn | string | X | √ |  | 
| engine_version | string | X | √ |  | 


