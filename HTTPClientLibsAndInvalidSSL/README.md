# Testing HTTP Client Libraries and Default Behavior with Invalid / Insecure SSL

| Environment | Library/Command | Behavior with Invalid/Insecure SSL |
|-------------|-----------------|-------------------------------------|
| Python      | requests        | Aborts on bad SSL                   |
| Python      | urllib          | Proceeds despite bad SSL            |
| Node        | axios           | Aborts on bad SSL                   |
| Node        | node-fetch      | Aborts on bad SSL                   |
| Bash        | curl/wget       | Aborts on bad SSL                   |
| Ruby        | net-http        | Aborts on bad SSL                   |



