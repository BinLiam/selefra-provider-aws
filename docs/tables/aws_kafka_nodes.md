# Table: aws_kafka_nodes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| broker_node_info | json | X | √ |  | 
| instance_type | string | X | √ |  | 
| node_type | string | X | √ |  | 
| zookeeper_node_info | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| added_to_cluster_time | string | X | √ |  | 
| node_arn | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| aws_kafka_clusters_selefra_id | string | X | X | fk to aws_kafka_clusters.selefra_id | 


