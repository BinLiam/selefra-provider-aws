# Table: aws_eventbridge_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| connection_arn | string | X | √ |  | 
| last_authorized_time | timestamp | X | √ |  | 
| state_reason | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| authorization_type | string | X | √ |  | 
| connection_state | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 


