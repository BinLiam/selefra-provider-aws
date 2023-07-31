# Table: aws_appsync_graphql_apis

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| user_pool_config | json | X | √ |  | 
| waf_web_acl_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| log_config | json | X | √ |  | 
| lambda_authorizer_config | json | X | √ |  | 
| open_id_connect_config | json | X | √ |  | 
| uris | json | X | √ |  | 
| xray_enabled | bool | X | √ |  | 
| additional_authentication_providers | json | X | √ |  | 
| api_id | string | X | √ |  | 
| authentication_type | string | X | √ |  | 


