# Table: aws_backup_report_plan

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| report_plan_name | string | X | √ |  | 
| description | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| deployment_status | string | X | √ |  | 
| last_attempted_execution_time | timestamp | X | √ |  | 
| last_successful_execution_time | timestamp | X | √ |  | 
| report_delivery_channel | json | X | √ |  | 
| report_setting | json | X | √ |  | 


