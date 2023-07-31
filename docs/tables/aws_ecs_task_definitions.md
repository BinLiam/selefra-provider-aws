# Table: aws_ecs_task_definitions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| execution_role_arn | string | X | √ |  | 
| ipc_mode | string | X | √ |  | 
| region | string | X | √ |  | 
| inference_accelerators | json | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| family | string | X | √ |  | 
| memory | string | X | √ |  | 
| account_id | string | X | √ |  | 
| cpu | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| task_definition_arn | string | X | √ |  | 
| proxy_configuration | json | X | √ |  | 
| registered_at | timestamp | X | √ |  | 
| compatibilities | string_array | X | √ |  | 
| placement_constraints | json | X | √ |  | 
| requires_attributes | json | X | √ |  | 
| runtime_platform | json | X | √ |  | 
| arn | string | √ | √ |  | 
| container_definitions | json | X | √ |  | 
| volumes | json | X | √ |  | 
| requires_compatibilities | string_array | X | √ |  | 
| network_mode | string | X | √ |  | 
| pid_mode | string | X | √ |  | 
| registered_by | string | X | √ |  | 
| revision | big_int | X | √ |  | 
| task_role_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| deregistered_at | timestamp | X | √ |  | 


