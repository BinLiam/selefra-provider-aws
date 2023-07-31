# Table: aws_appstream_users

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_time | timestamp | X | √ |  | 
| first_name | string | X | √ |  | 
| last_name | string | X | √ |  | 
| user_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| enabled | bool | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| authentication_type | string | X | √ |  | 


