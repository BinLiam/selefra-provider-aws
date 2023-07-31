# Table: aws_wafv2_ipsets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ip_address_version | string | X | √ |  | 
| id | string | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| addresses | ip_array | X | √ |  | 
| name | string | X | √ |  | 


