# Configuration Example

This document will outline sample configuration to help configure to your desired configuration. You can find field 
functionality and descriptions in the [config.yml](../config.yml) file itself.

## Table of Contents
1. [GitHub Configuration sample](#github-configuration-sample)
   1. [Organisation based authentication](#github-organisation-based-authentication)
   2. [Team based authentication](#github-team-based-authentication)
2. [GitLab Configuration Sample](#gitlab-configuration-sample)
   1. [Organisation based authentication](#gitlab-organisation-based-authentication)
   2. [Group based authentication](#gitlab-group-based-authentication)
2. [SQL Configuration Sample](#sql-configuration-sample)


## GitHub Configuration sample

### GitHub organisation based authentication
```yaml
system_conf:
  name: Test
  method: org
  service: github
  admin_user: root
  default_user: ubuntu

github:
  access_token: 0ce2e0cc1187234bc802b..........................
  admin_role: admin
  default_role: all
  org: DummyAccount
  team_name:

gitlab:
  org_url:
  access_token:
  group_name:
  admin_role:
  default_role:
  inherit_permission:
    admin_user:
    default_user:
    admin_stack:
    default_stack:

sql:
  host:
  user:
  password:
  database:
  table:

alerts:
  slack: https://hooks.slack.com/services/T2Y...BV4/B6VB......................
  hipchat:
    url:
    token:
```

### GitHub team based authentication
```yaml
system_conf:
  name: Test
  method: team
  service: github
  admin_user: root
  default_user: ubuntu

github:
  access_token: 0ce2e0cc1187234bc802b..........................
  admin_role: maintainer
  default_role: all
  org: DummyAccount
  team_name: devops

gitlab:
  org_url:
  access_token:
  group_name:
  admin_role:
  default_role:
  inherit_permission:
    admin_user:
    default_user:
    admin_stack:
    default_stack:

sql:
  host:
  user:
  password:
  database:
  table:

alerts:
  slack: https://hooks.slack.com/services/T2Y...BV4/B6VB......................
  hipchat:
    url:
    token:
```

## GitLab Configuration sample

### GitLab organisation based authentication
```yaml
system_conf:
  name: Test
  method: org
  service: gitlab
  admin_user: root
  default_user: ubuntu

github:
  access_token:
  admin_role:
  default_role:
  org:
  team_name:

gitlab:
  org_url: https://gitlab.domain.com
  access_token: pcN2K9xJA4hGAy........
  group_name:
  admin_role: admin
  default_role: all
  inherit_permission:
    admin_user:
    default_user:
    admin_stack:
    default_stack:

sql:
  host:
  user:
  password:
  database:
  table:

alerts:
  slack: https://hooks.slack.com/services/T2Y...BV4/B6VB......................
  hipchat:
    url:
    token:
```

### GitLab group based authentication
```yaml
system_conf:
  name: Test
  method: group
  service: gitlab
  admin_user: root
  default_user: ubuntu

github:
  access_token:
  admin_role:
  default_role:
  org:
  team_name:

gitlab:
  org_url: https://gitlab.domain.com
  access_token: pcN2K9xJA4hGAy........
  group_name: Test
  admin_role: owner
  default_role: developer
  inherit_permission:
    admin_user: true
    default_user: true
    admin_stack: down
    default_stack: down

sql:
  host:
  user:
  password:
  database:
  table:

alerts:
  slack: https://hooks.slack.com/services/T2Y...BV4/B6VB......................
  hipchat:
    url:
    token:
```

## SQL Configuration sample


```yaml
system_conf:
  name: Test
  method: group
  service: gitlab
  admin_user: root
  default_user: ubuntu

github:
  access_token:
  admin_role:
  default_role:
  org:
  team_name:

gitlab:
  org_url:
  access_token:
  group_name:
  admin_role:
  default_role:
  inherit_permission:
    admin_user:
    default_user:
    admin_stack:
    default_stack:

sql:
  host: 0.0.0.0
  user: dummyAccount
  password: kjasnjIU891NoihnN
  database: authentication
  table: sshLogin

alerts:
  slack: https://hooks.slack.com/services/T2Y...BV4/B6VB......................
  hipchat:
    url:
    token:
```