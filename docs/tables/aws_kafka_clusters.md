# Table: aws_kafka_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| serverless | json | X | √ |  | 
| active_operation_arn | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_type | string | X | √ |  | 
| current_version | string | X | √ |  | 
| provisioned | json | X | √ |  | 
| state_info | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cluster_name | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 


