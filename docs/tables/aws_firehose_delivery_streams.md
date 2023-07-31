# Table: aws_firehose_delivery_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| delivery_stream_type | string | X | √ |  | 
| has_more_destinations | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| delivery_stream_status | string | X | √ |  | 
| version_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| source | json | X | √ |  | 
| delivery_stream_name | string | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| delivery_stream_encryption_configuration | json | X | √ |  | 
| last_update_timestamp | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| delivery_stream_arn | string | X | √ |  | 
| destinations | json | X | √ |  | 
| failure_description | json | X | √ |  | 


