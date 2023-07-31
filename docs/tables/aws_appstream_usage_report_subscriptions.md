# Table: aws_appstream_usage_report_subscriptions

## Primary Keys 

```
account_id, region, s3_bucket_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| last_generated_report_date | timestamp | X | √ |  | 
| s3_bucket_name | string | X | √ |  | 
| schedule | string | X | √ |  | 
| subscription_errors | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


