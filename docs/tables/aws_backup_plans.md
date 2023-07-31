# Table: aws_backup_plans

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| backup_plan_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| version_id | string | X | √ |  | 
| backup_plan_arn | string | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| last_execution_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| advanced_backup_settings | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_plan | json | X | √ |  | 


