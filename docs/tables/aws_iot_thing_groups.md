# Table: aws_iot_thing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| thing_group_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| things_in_group | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| index_name | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| thing_group_name | string | X | √ |  | 
| thing_group_properties | json | X | √ |  | 
| query_string | string | X | √ |  | 
| arn | string | √ | √ |  | 
| thing_group_id | string | X | √ |  | 
| version | big_int | X | √ |  | 
| result_metadata | json | X | √ |  | 
| policies | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| query_version | string | X | √ |  | 
| thing_group_arn | string | X | √ |  | 


