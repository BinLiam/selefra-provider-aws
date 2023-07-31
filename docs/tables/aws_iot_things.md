# Table: aws_iot_things

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| version | big_int | X | √ |  | 
| region | string | X | √ |  | 
| principals | string_array | X | √ |  | 
| thing_arn | string | X | √ |  | 
| thing_type_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| attributes | json | X | √ |  | 
| thing_name | string | X | √ |  | 


