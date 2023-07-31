# Table: aws_ssm_inventory_schemas

## Primary Keys 

```
account_id, region, type_name, version
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attributes | json | X | √ |  | 
| type_name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


