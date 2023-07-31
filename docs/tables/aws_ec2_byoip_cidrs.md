# Table: aws_ec2_byoip_cidrs

## Primary Keys 

```
account_id, region, cidr
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| state | string | X | √ |  | 
| status_message | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cidr | string | X | √ |  | 


