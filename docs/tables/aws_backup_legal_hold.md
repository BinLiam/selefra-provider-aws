# Table: aws_backup_legal_hold

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| legal_hold_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_date | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| cancellation_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| retain_record_until | timestamp | X | √ |  | 
| recovery_point_selection | json | X | √ |  | 


