# Table: aws_config_configuration_recorders

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| configuration_recorder | json | X | √ |  | 
| status_last_error_code | string | X | √ |  | 
| status_last_status | string | X | √ |  | 
| status_last_stop_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| status_last_error_message | string | X | √ |  | 
| status_last_start_time | timestamp | X | √ |  | 
| status_last_status_change_time | timestamp | X | √ |  | 
| status_recording | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


