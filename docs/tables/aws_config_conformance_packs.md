# Table: aws_config_conformance_packs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| delivery_s3_key_prefix | string | X | √ |  | 
| template_ssm_document_details | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| conformance_pack_arn | string | X | √ |  | 
| conformance_pack_name | string | X | √ |  | 
| created_by | string | X | √ |  | 
| delivery_s3_bucket | string | X | √ |  | 
| last_update_requested_time | timestamp | X | √ |  | 
| conformance_pack_id | string | X | √ |  | 
| conformance_pack_input_parameters | json | X | √ |  | 


