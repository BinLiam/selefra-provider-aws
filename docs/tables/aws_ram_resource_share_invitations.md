# Table: aws_ram_resource_share_invitations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_share_associations | json | X | √ |  | 
| resource_share_invitation_arn | string | X | √ |  | 
| resource_share_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| receiver_arn | string | X | √ |  | 
| receiver_account_id | string | X | √ |  | 
| resource_share_arn | string | X | √ |  | 
| sender_account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| invitation_timestamp | timestamp | X | √ |  | 


