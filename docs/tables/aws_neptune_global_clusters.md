# Table: aws_neptune_global_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| global_cluster_arn | string | X | √ |  | 
| global_cluster_resource_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| global_cluster_members | json | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| engine_version | string | X | √ |  | 
| arn | string | √ | √ |  | 
| deletion_protection | bool | X | √ |  | 
| global_cluster_identifier | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| region | string | X | √ |  | 


