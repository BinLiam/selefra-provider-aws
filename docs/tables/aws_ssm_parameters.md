# Table: aws_ssm_parameters

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_date | timestamp | X | √ |  | 
| allowed_pattern | string | X | √ |  | 
| name | string | X | √ | `The parameter name` | 
| policies | json | X | √ |  | 
| version | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| data_type | string | X | √ |  | 
| type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| key_id | string | X | √ |  | 
| last_modified_user | string | X | √ |  | 
| tier | string | X | √ |  | 
| region | string | X | √ |  | 


