# Table: aws_ecr_repositories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_configuration | json | X | √ |  | 
| image_tag_mutability | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| registry_id | string | X | √ |  | 
| repository_name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| repository_arn | string | X | √ |  | 
| policy_text | json | X | √ |  | 
| image_scanning_configuration | json | X | √ |  | 
| repository_uri | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 


