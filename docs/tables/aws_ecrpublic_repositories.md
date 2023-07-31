# Table: aws_ecrpublic_repositories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| repository_uri | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| registry_id | string | X | √ |  | 
| repository_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| repository_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 


