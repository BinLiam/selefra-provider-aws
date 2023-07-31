# Table: aws_glue_classifiers

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| grok_classifier | json | X | √ |  | 
| json_classifier | json | X | √ |  | 
| xml_classifier | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| csv_classifier | json | X | √ |  | 


