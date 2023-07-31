# Table: aws_rds_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_type | string | X | √ |  | 
| thumbprint | string | X | √ |  | 
| valid_from | timestamp | X | √ |  | 
| valid_till | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| certificate_identifier | string | X | √ |  | 
| customer_override | bool | X | √ |  | 
| customer_override_valid_till | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| certificate_arn | string | X | √ |  | 


