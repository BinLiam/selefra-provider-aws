# Table: aws_kms_key_policies

## Primary Keys 

```
key_arn, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| policy | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| key_arn | string | X | √ |  | 
| aws_kms_keys_selefra_id | string | X | X | fk to aws_kms_keys.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


