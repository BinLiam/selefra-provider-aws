# Table: aws_xray_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| filter_expression | string | X | √ |  | 
| group_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| group_arn | string | X | √ |  | 
| insights_configuration | json | X | √ |  | 
| tags | json | X | √ |  | 


