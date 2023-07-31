# Table: aws_iam_groups

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| group_id | string | X | √ |  | 
| group_name | string | X | √ |  | 
| path | string | X | √ |  | 
| policies | json | X | √ |  | 
| arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_date | timestamp | X | √ |  | 


