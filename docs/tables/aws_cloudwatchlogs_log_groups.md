# Table: aws_cloudwatchlogs_log_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ |  | 
| metric_filter_count | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| creation_time | big_int | X | √ |  | 
| retention_in_days | big_int | X | √ |  | 
| stored_bytes | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| data_protection_status | string | X | √ |  | 
| log_group_name | string | X | √ |  | 


