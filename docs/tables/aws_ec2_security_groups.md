# Table: aws_ec2_security_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| group_name | string | X | √ |  | 
| ip_permissions | json | X | √ |  | 
| ip_permissions_egress | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| group_id | string | X | √ |  | 


