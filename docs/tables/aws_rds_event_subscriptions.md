# Table: aws_rds_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| source_type | string | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| event_subscription_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sns_topic_arn | string | X | √ |  | 


