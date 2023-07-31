# Table: aws_docdb_event_subscriptions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| event_categories_list | string_array | X | √ |  | 
| event_subscription_arn | string | X | √ |  | 
| source_type | string | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| status | string | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


