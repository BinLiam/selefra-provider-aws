# Table: aws_amp_workspaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| logging_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| status | json | X | √ |  | 
| alert_manager_definition | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| workspace_id | string | X | √ |  | 
| alias | string | X | √ |  | 
| prometheus_endpoint | string | X | √ |  | 
| tags | json | X | √ |  | 


