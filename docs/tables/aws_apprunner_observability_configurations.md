# Table: aws_apprunner_observability_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| latest | bool | X | √ |  | 
| observability_configuration_revision | big_int | X | √ |  | 
| status | string | X | √ |  | 
| trace_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| observability_configuration_arn | string | X | √ |  | 
| observability_configuration_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


