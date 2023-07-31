# Table: aws_eventbridge_api_destinations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| invocation_endpoint | string | X | √ |  | 
| invocation_rate_limit_per_second | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| api_destination_state | string | X | √ |  | 
| connection_arn | string | X | √ |  | 
| http_method | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_destination_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


