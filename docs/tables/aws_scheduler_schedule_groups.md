# Table: aws_scheduler_schedule_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_date | timestamp | X | √ |  | 
| last_modification_date | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


