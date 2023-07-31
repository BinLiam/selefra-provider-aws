# Table: aws_kinesis_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| stream_arn | string | X | √ |  | 
| retention_period_hours | big_int | X | √ |  | 
| stream_creation_timestamp | timestamp | X | √ |  | 
| stream_name | string | X | √ |  | 
| open_shard_count | big_int | X | √ |  | 
| stream_mode_details | json | X | √ |  | 
| stream_status | string | X | √ |  | 
| consumer_count | big_int | X | √ |  | 
| key_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| enhanced_monitoring | json | X | √ |  | 


