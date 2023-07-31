# Table: aws_dynamodb_tables

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| table_size_bytes | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| key_schema | json | X | √ |  | 
| latest_stream_arn | string | X | √ |  | 
| creation_date_time | timestamp | X | √ |  | 
| item_count | big_int | X | √ |  | 
| region | string | X | √ |  | 
| attribute_definitions | json | X | √ |  | 
| local_secondary_indexes | json | X | √ |  | 
| account_id | string | X | √ |  | 
| table_id | string | X | √ |  | 
| global_secondary_indexes | json | X | √ |  | 
| global_table_version | string | X | √ |  | 
| provisioned_throughput | json | X | √ |  | 
| table_arn | string | X | √ |  | 
| table_status | string | X | √ |  | 
| archival_summary | json | X | √ |  | 
| restore_summary | json | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| billing_mode_summary | json | X | √ |  | 
| latest_stream_label | string | X | √ |  | 
| sse_description | json | X | √ |  | 
| stream_specification | json | X | √ |  | 
| replicas | json | X | √ |  | 
| table_name | string | X | √ |  | 
| table_class_summary | json | X | √ |  | 


