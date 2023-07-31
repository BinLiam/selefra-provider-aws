# Table: aws_sqs_queues

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| visibility_timeout | big_int | X | √ |  | 
| fifo_throughput_limit | string | X | √ |  | 
| deduplication_scope | string | X | √ |  | 
| receive_message_wait_time_seconds | big_int | X | √ |  | 
| approximate_number_of_messages | big_int | X | √ |  | 
| approximate_number_of_messages_delayed | big_int | X | √ |  | 
| approximate_number_of_messages_not_visible | big_int | X | √ |  | 
| created_timestamp | big_int | X | √ |  | 
| sqs_managed_sse_enabled | bool | X | √ |  | 
| fifo_queue | bool | X | √ |  | 
| redrive_allow_policy | json | X | √ |  | 
| url | string | X | √ |  | 
| tags | json | X | √ |  | 
| policy | json | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| account_id | string | X | √ |  | 
| message_retention_period | big_int | X | √ |  | 
| maximum_message_size | big_int | X | √ |  | 
| redrive_policy | json | X | √ |  | 
| kms_data_key_reuse_period_seconds | big_int | X | √ |  | 
| delay_seconds | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| content_based_deduplication | bool | X | √ |  | 
| last_modified_timestamp | big_int | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kms_master_key_id | string | X | √ |  | 


