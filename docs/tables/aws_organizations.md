# Table: aws_organizations

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| feature_set | string | X | √ |  | 
| id | string | X | √ |  | 
| master_account_arn | string | X | √ |  | 
| available_policy_types | json | X | √ |  | 
| master_account_email | string | X | √ |  | 
| master_account_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


