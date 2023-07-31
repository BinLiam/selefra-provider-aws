# Table: aws_lightsail_container_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| current_deployment | json | X | √ |  | 
| power | string | X | √ |  | 
| power_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| container_service_name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| state_detail | json | X | √ |  | 
| tags | json | X | √ |  | 
| scale | big_int | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| is_disabled | bool | X | √ |  | 
| location | json | X | √ |  | 
| private_domain_name | string | X | √ |  | 
| private_registry_access | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| region | string | X | √ |  | 
| next_deployment | json | X | √ |  | 
| principal_arn | string | X | √ |  | 
| public_domain_names | json | X | √ |  | 
| url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


