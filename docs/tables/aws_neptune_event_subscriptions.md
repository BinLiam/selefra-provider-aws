# Table: aws_neptune_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| sns_topic_arn | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| source_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| customer_aws_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| status | string | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| event_subscription_arn | string | X | √ |  | 


