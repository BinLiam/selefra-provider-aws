# Table: aws_transfer_servers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| domain | string | X | √ |  | 
| endpoint_details | json | X | √ |  | 
| identity_provider_type | string | X | √ |  | 
| protocol_details | json | X | √ |  | 
| workflow_details | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| endpoint_type | string | X | √ |  | 
| host_key_fingerprint | string | X | √ |  | 
| logging_role | string | X | √ |  | 
| security_policy_name | string | X | √ |  | 
| tags | json | X | √ | `Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described` | 
| user_count | big_int | X | √ |  | 
| certificate | string | X | √ |  | 
| identity_provider_details | json | X | √ |  | 
| post_authentication_login_banner | string | X | √ |  | 
| server_id | string | X | √ |  | 
| pre_authentication_login_banner | string | X | √ |  | 
| protocols | string_array | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 


