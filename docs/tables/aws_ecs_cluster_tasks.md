# Table: aws_ecs_cluster_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| version | big_int | X | √ |  | 
| containers | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| memory | string | X | √ |  | 
| pull_started_at | timestamp | X | √ |  | 
| started_by | string | X | √ |  | 
| stopping_at | timestamp | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| inference_accelerators | json | X | √ |  | 
| overrides | json | X | √ |  | 
| pull_stopped_at | timestamp | X | √ |  | 
| started_at | timestamp | X | √ |  | 
| availability_zone | string | X | √ |  | 
| desired_status | string | X | √ |  | 
| stopped_reason | string | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| connectivity_at | timestamp | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| execution_stopped_at | timestamp | X | √ |  | 
| health_status | string | X | √ |  | 
| last_status | string | X | √ |  | 
| task_protection | json | X | √ |  | 
| group | string | X | √ |  | 
| platform_family | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| task_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| attachments | json | X | √ |  | 
| connectivity | string | X | √ |  | 
| enable_execute_command | bool | X | √ |  | 
| stopped_at | timestamp | X | √ |  | 
| attributes | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| launch_type | string | X | √ |  | 
| stop_code | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cpu | string | X | √ |  | 
| tags | json | X | √ |  | 
| task_definition_arn | string | X | √ |  | 


