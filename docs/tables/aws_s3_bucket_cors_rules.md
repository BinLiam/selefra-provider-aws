# Table: aws_s3_bucket_cors_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allowed_methods | string_array | X | √ |  | 
| allowed_origins | string_array | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| allowed_headers | string_array | X | √ |  | 
| expose_headers | string_array | X | √ |  | 
| id | string | X | √ |  | 
| max_age_seconds | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 


