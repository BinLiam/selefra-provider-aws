# Table: aws_apprunner_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| network_configuration | json | X | √ |  | 
| service_name | string | X | √ |  | 
| source_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| encryption_configuration | json | X | √ |  | 
| auto_scaling_configuration_summary | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| instance_configuration | json | X | √ |  | 
| service_arn | string | X | √ |  | 
| service_id | string | X | √ |  | 
| status | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| service_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| deleted_at | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| health_check_configuration | json | X | √ |  | 
| observability_configuration | json | X | √ |  | 


