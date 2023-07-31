# Table: aws_servicecatalog_provisioned_products

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| physical_id | string | X | √ |  | 
| type | string | X | √ |  | 
| user_arn_session | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| provisioning_artifact_id | string | X | √ |  | 
| status | string | X | √ |  | 
| status_message | string | X | √ |  | 
| tags | json | X | √ |  | 
| user_arn | string | X | √ |  | 
| idempotency_token | string | X | √ |  | 
| product_id | string | X | √ |  | 
| last_record_id | string | X | √ |  | 
| provisioning_artifact_name | string | X | √ |  | 
| last_successful_provisioning_record_id | string | X | √ |  | 
| product_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| last_provisioning_record_id | string | X | √ |  | 


