# Table: aws_dax_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_endpoint_encryption_type | string | X | √ |  | 
| description | string | X | √ |  | 
| node_ids_to_remove | string_array | X | √ |  | 
| active_nodes | big_int | X | √ |  | 
| cluster_name | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| nodes | json | X | √ |  | 
| security_groups | json | X | √ |  | 
| cluster_discovery_endpoint | json | X | √ |  | 
| node_type | string | X | √ |  | 
| notification_configuration | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| parameter_group | json | X | √ |  | 
| sse_description | json | X | √ |  | 
| status | string | X | √ |  | 
| subnet_group | string | X | √ |  | 
| total_nodes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 


