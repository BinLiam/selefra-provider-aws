# Table: aws_kafka_cluster_operations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| operation_arn | string | X | √ |  | 
| operation_state | string | X | √ |  | 
| tags | json | X | √ |  | 
| client_request_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| error_info | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| end_time | timestamp | X | √ |  | 
| operation_type | string | X | √ |  | 
| aws_kafka_clusters_selefra_id | string | X | X | fk to aws_kafka_clusters.selefra_id | 
| target_cluster_info | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cluster_arn | string | X | √ |  | 
| operation_steps | json | X | √ |  | 
| source_cluster_info | json | X | √ |  | 


