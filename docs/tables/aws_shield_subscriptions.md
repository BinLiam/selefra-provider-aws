# Table: aws_shield_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_renew | string | X | √ |  | 
| limits | json | X | √ |  | 
| proactive_engagement_status | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| time_commitment_in_seconds | big_int | X | √ |  | 
| subscription_limits | json | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| subscription_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 


