# Table: aws_cognito_identity_pools

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| identity_pool_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| identity_pool_tags | json | X | √ |  | 
| saml_provider_ar_ns | string_array | X | √ |  | 
| supported_login_providers | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| allow_unauthenticated_identities | bool | X | √ |  | 
| allow_classic_flow | bool | X | √ |  | 
| cognito_identity_providers | json | X | √ |  | 
| developer_provider_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| identity_pool_id | string | X | √ |  | 
| open_id_connect_provider_ar_ns | string_array | X | √ |  | 
| id | string | X | √ |  | 


