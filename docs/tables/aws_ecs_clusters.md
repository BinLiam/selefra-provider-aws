# Table: aws_ecs_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| capacity_providers | string_array | X | √ |  | 
| configuration | json | X | √ |  | 
| pending_tasks_count | big_int | X | √ |  | 
| running_tasks_count | big_int | X | √ |  | 
| service_connect_defaults | json | X | √ |  | 
| account_id | string | X | √ |  | 
| active_services_count | big_int | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| statistics | json | X | √ |  | 
| attachments_status | string | X | √ |  | 
| registered_container_instances_count | big_int | X | √ |  | 
| status | string | X | √ |  | 
| tags | json | X | √ |  | 
| attachments | json | X | √ |  | 
| default_capacity_provider_strategy | json | X | √ |  | 
| settings | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


