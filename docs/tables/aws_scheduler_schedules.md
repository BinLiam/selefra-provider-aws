# Table: aws_scheduler_schedules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| kms_key_arn | string | X | √ |  | 
| last_modification_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| flexible_time_window | json | X | √ |  | 
| group_name | string | X | √ |  | 
| start_date | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| target | json | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| end_date | timestamp | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| schedule_expression_timezone | string | X | √ |  | 


