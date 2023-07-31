# Table: aws_sns_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| protocol | string | X | √ |  | 
| subscription_arn | string | X | √ |  | 
| pending_confirmation | bool | X | √ |  | 
| redrive_policy | json | X | √ |  | 
| endpoint | string | X | √ |  | 
| owner | string | X | √ |  | 
| confirmation_was_authenticated | bool | X | √ |  | 
| effective_delivery_policy | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| topic_arn | string | X | √ |  | 
| filter_policy | json | X | √ |  | 
| subscription_role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| delivery_policy | json | X | √ |  | 
| raw_message_delivery | bool | X | √ |  | 
| unknown_fields | json | X | √ |  | 


