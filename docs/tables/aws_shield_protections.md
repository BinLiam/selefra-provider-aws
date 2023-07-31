# Table: aws_shield_protections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| health_check_ids | string_array | X | √ |  | 
| protection_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_layer_automatic_response_configuration | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_arn | string | X | √ |  | 


