# Table: aws_elasticbeanstalk_environments

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| environment_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| health | string | X | √ |  | 
| status | string | X | √ |  | 
| tier | json | X | √ |  | 
| version_label | string | X | √ |  | 
| abortable_operation_in_progress | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| listeners | json | X | √ |  | 
| application_name | string | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| environment_id | string | X | √ |  | 
| solution_stack_name | string | X | √ |  | 
| resources | json | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | X | √ |  | 
| description | string | X | √ |  | 
| endpoint_url | string | X | √ |  | 
| environment_name | string | X | √ |  | 
| platform_arn | string | X | √ |  | 
| cname | string | X | √ |  | 
| environment_links | json | X | √ |  | 
| template_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| date_created | timestamp | X | √ |  | 
| health_status | string | X | √ |  | 
| operations_role | string | X | √ |  | 


