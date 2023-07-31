# Table: aws_iam_virtual_mfa_devices

## Primary Keys 

```
serial_number
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| serial_number | string | √ | √ |  | 
| base32_string_seed | byte_array | X | √ |  | 
| enable_date | timestamp | X | √ |  | 
| qr_code_png | byte_array | X | √ |  | 
| tags | json | X | √ |  | 
| user | json | X | √ |  | 
| account_id | string | X | √ |  | 


