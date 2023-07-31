# Table: aws_ssm_documents

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| latest_version | string | X | √ |  | 
| pending_review_version | string | X | √ |  | 
| schema_version | string | X | √ |  | 
| status | string | X | √ |  | 
| status_information | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_date | timestamp | X | √ |  | 
| document_version | string | X | √ |  | 
| owner | string | X | √ |  | 
| target_type | string | X | √ |  | 
| approved_version | string | X | √ |  | 
| default_version | string | X | √ |  | 
| document_format | string | X | √ |  | 
| permissions | json | X | √ |  | 
| attachments_information | json | X | √ |  | 
| author | string | X | √ |  | 
| tags | json | X | √ |  | 
| version_name | string | X | √ |  | 
| category | string_array | X | √ |  | 
| display_name | string | X | √ |  | 
| document_type | string | X | √ |  | 
| parameters | json | X | √ |  | 
| requires | json | X | √ |  | 
| review_information | json | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| platform_types | string_array | X | √ |  | 
| review_status | string | X | √ |  | 
| sha1 | string | X | √ |  | 
| category_enum | string_array | X | √ |  | 
| description | string | X | √ |  | 
| hash | string | X | √ |  | 
| hash_type | string | X | √ |  | 
| arn | string | √ | √ |  | 


