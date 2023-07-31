# Table: aws_eks_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resources_vpc_config | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| connector_config | json | X | √ |  | 
| encryption_config | json | X | √ |  | 
| identity | json | X | √ |  | 
| tags | json | X | √ |  | 
| version | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| outpost_config | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| id | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| health | json | X | √ |  | 
| kubernetes_network_config | json | X | √ |  | 
| logging | json | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_authority | json | X | √ |  | 
| client_request_token | string | X | √ |  | 
| status | string | X | √ |  | 


