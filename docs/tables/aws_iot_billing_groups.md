# Table: aws_iot_billing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| billing_group_metadata | json | X | √ |  | 
| billing_group_name | string | X | √ |  | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| billing_group_arn | string | X | √ |  | 
| billing_group_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| things_in_group | string_array | X | √ |  | 
| billing_group_properties | json | X | √ |  | 
| result_metadata | json | X | √ |  | 


