# Table: aws_ses_contact_lists

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_timestamp | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| topics | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| contact_list_name | string | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_timestamp | timestamp | X | √ |  | 
| result_metadata | json | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


