## TL:DR 

```shell
open https://id.atlassian.com/manage-profile/security/api-tokens
# export JIRA_API_TOKEN=...

mkdir -p $HOME/.jira.d
cat <<EOF > $HOME/.jira.d/config.yml
---
authentication-method: api-token
endpoint: https://yourcompany.atlassian.net
user: Your Name
login: your-email@yourcompany.com
project: YOURPROJECT
EOF

brew install go
git clone /somewhere
cd /somewhere
make build
mv ./jira /on/your/path
jira list
```
