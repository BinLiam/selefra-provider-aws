# Table: aws_lightsail_bucket_access_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| access_key_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| last_used | json | X | √ |  | 
| secret_access_key | string | X | √ |  | 
| status | string | X | √ |  | 
| aws_lightsail_buckets_selefra_id | string | X | X | fk to aws_lightsail_buckets.selefra_id | 
| selefra_id | string | √ | √ | random id | 


