# Table: aws_s3_bucket_grants

## Primary Keys 

```
bucket_arn, grantee_type, grantee_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| grantee_type | string | X | √ |  | 
| grantee_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| grantee | json | X | √ |  | 
| permission | string | X | √ |  | 
| account_id | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 


