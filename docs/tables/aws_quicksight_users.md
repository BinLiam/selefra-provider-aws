# Table: aws_quicksight_users

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| principal_id | string | X | √ |  | 
| external_login_federation_provider_type | string | X | √ |  | 
| custom_permissions_name | string | X | √ |  | 
| email | string | X | √ |  | 
| role | string | X | √ |  | 
| user_name | string | X | √ |  | 
| active | bool | X | √ |  | 
| external_login_federation_provider_url | string | X | √ |  | 
| external_login_id | string | X | √ |  | 
| identity_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 


