# Table: aws_ses_custom_verification_email_templates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| success_redirection_url | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| failure_redirection_url | string | X | √ |  | 
| from_email_address | string | X | √ |  | 
| template_subject | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| template_content | string | X | √ |  | 
| template_name | string | X | √ |  | 


