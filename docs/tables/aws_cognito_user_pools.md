# Table: aws_cognito_user_pools

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| email_verification_message | string | X | √ |  | 
| sms_configuration | json | X | √ |  | 
| admin_create_user_config | json | X | √ |  | 
| sms_authentication_message | string | X | √ |  | 
| sms_configuration_failure | string | X | √ |  | 
| username_attributes | string_array | X | √ |  | 
| verification_message_template | json | X | √ |  | 
| deletion_protection | string | X | √ |  | 
| mfa_configuration | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| email_configuration | json | X | √ |  | 
| id | string | X | √ |  | 
| policies | json | X | √ |  | 
| status | string | X | √ |  | 
| auto_verified_attributes | string_array | X | √ |  | 
| email_configuration_failure | string | X | √ |  | 
| name | string | X | √ |  | 
| username_configuration | json | X | √ |  | 
| account_recovery_setting | json | X | √ |  | 
| estimated_number_of_users | big_int | X | √ |  | 
| schema_attributes | json | X | √ |  | 
| user_pool_tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| custom_domain | string | X | √ |  | 
| email_verification_subject | string | X | √ |  | 
| sms_verification_message | string | X | √ |  | 
| user_attribute_update_settings | json | X | √ |  | 
| region | string | X | √ |  | 
| alias_attributes | string_array | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| device_configuration | json | X | √ |  | 
| domain | string | X | √ |  | 
| lambda_config | json | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| user_pool_add_ons | json | X | √ |  | 


