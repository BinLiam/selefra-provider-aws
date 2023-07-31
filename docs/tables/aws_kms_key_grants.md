# Table: aws_kms_key_grants

## Primary Keys 

```
key_arn, grant_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| key_id | string | X | √ |  | 
| key_arn | string | X | √ |  | 
| constraints | json | X | √ |  | 
| retiring_principal | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_kms_keys_selefra_id | string | X | X | fk to aws_kms_keys.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| operations | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| grant_id | string | X | √ |  | 
| grantee_principal | string | X | √ |  | 
| issuing_account | string | X | √ |  | 
| name | string | X | √ |  | 


