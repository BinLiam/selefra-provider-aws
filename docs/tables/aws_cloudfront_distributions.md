# Table: aws_cloudfront_distributions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| active_trusted_key_groups | json | X | √ |  | 
| active_trusted_signers | json | X | √ |  | 
| alias_icp_recordals | json | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| distribution_config | json | X | √ |  | 
| domain_name | string | X | √ |  | 
| id | string | X | √ |  | 
| in_progress_invalidation_batches | big_int | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 


