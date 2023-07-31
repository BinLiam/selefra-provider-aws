# Table: aws_lightsail_static_ips

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| ip_address | string | X | √ |  | 
| location | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| attached_to | string | X | √ |  | 
| is_attached | bool | X | √ |  | 
| name | string | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


