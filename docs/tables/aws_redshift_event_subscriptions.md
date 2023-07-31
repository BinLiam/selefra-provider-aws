# Table: aws_redshift_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cust_subscription_id | string | X | √ |  | 
| source_type | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| arn | string | √ | √ | `ARN of the event subscription.` | 
| event_categories_list | string_array | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| subscription_creation_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| severity | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


