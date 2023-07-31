# Table: aws_docdb_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| thumbprint | string | X | √ |  | 
| valid_from | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| certificate_arn | string | X | √ |  | 
| certificate_type | string | X | √ |  | 
| valid_till | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_identifier | string | X | √ |  | 


