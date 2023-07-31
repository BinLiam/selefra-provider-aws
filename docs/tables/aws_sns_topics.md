# Table: aws_sns_topics

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_master_key_id | string | X | √ |  | 
| fifo_topic | bool | X | √ |  | 
| content_based_deduplication | bool | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| subscriptions_confirmed | big_int | X | √ |  | 
| arn | string | √ | √ |  | 
| effective_delivery_policy | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| display_name | string | X | √ |  | 
| policy | json | X | √ |  | 
| subscriptions_pending | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| delivery_policy | json | X | √ |  | 
| owner | string | X | √ |  | 
| subscriptions_deleted | big_int | X | √ |  | 
| region | string | X | √ |  | 


