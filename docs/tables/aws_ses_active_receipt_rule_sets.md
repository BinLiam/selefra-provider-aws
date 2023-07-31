# Table: aws_ses_active_receipt_rule_sets

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| metadata | json | X | √ |  | 
| rules | json | X | √ |  | 


