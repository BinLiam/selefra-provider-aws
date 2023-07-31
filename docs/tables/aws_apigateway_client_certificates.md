# Table: aws_apigateway_client_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| client_certificate_id | string | X | √ |  | 
| expiration_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| pem_encoded_certificate | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 


