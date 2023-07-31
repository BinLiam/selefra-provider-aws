# Table: aws_iot_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| stream_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| role_arn | string | X | √ |  | 
| stream_id | string | X | √ |  | 
| stream_version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| files | json | X | √ |  | 


