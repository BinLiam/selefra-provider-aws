# Table: aws_docdb_global_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deletion_protection | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| global_cluster_identifier | string | X | √ |  | 
| region | string | X | √ |  | 
| database_name | string | X | √ |  | 
| global_cluster_arn | string | X | √ |  | 
| global_cluster_members | json | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| engine | string | X | √ |  | 
| global_cluster_resource_id | string | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


