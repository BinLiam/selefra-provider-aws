# Table: aws_ecs_cluster_container_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| agent_connected | bool | X | √ |  | 
| health_status | json | X | √ |  | 
| remaining_resources | json | X | √ |  | 
| running_tasks_count | big_int | X | √ |  | 
| status | string | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| agent_update_status | string | X | √ |  | 
| attributes | json | X | √ |  | 
| tags | json | X | √ |  | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| attachments | json | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| registered_resources | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| ec2_instance_id | string | X | √ |  | 
| pending_tasks_count | big_int | X | √ |  | 
| registered_at | timestamp | X | √ |  | 
| status_reason | string | X | √ |  | 
| version_info | json | X | √ |  | 


