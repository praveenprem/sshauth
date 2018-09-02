# SSH Auth

 Is an authentication plugin for SSH, which allow system administrators to integrate third-party service such as GitHub and 
 GitLab as a authentication mechanism in conjunction with default `authorized_keys` file. This is stateless and therefore
 any changes made on third-party service side will take effect immediately across all systems.  
 
 This plugin will allow you to configure one of the following service as an authentication mechanism:
 - GitHub v1.0.0
 - GitLab (Self-Hosted only) v1.0.0
 - AWS IAM v2.0.0
 - SQL server v2.0.0
 ###### _Version attached to the service refer to the release version of this plugin, not the service itself._
 
_**WARNING:** This plugin is NOT a replacement option for default_ **authorized_keys**!
_Should always be used as a failover._
 
## Table of Contents
 
 1. [Getting Started](#getting-started)
    1. [Prerequisites](#prerequisites)
    2. [Installation](#installation)
    3. [Test installation and configuration](#test-installation-and-configuration)
 2. [Alerting](#alerting)
    1. [Example](#alerting-example)
 3. [Logging](#logging)
 4. [Resources](#resources)
 5. [License](#license)
 6. [Versioning](#versioning)
 7. [Authors](#authors)
 8. [Acknowledgments](#acknowledgments)
 
---
 
## Getting Started
 
 Following instruction will guide you through how to install and configure this plugin on a Unix/Linux server.
 
### Prerequisites

 - Root access to the server in question
 - Administrative access to one of third-party services listed above
 - Make sure you've an uninterrupted connection to the server in questions, especially when making changes to the SSH
   daemon.
 
### Installation

###### _It is recommended to run following installation instruction as `root` user of the system to avoid any permission issues during runtime and [testing](#test-installation-and-configuration) stage._

 1. Download a copy of the latest release from the [Releases tab](https://github.com/praveenprem/sshauth/releases). Which
    consist of SSHAuth binary and the latest [config.yml](./resources/config.yml).
 2. Unzip the TAR (Zip) and copy the `sshauth` binary file into `/usr/local/bin/` folder. <h6>This is not plugin requirement
    this location can be anywhere in the system as long as the path to file exist in the $PATH variable and binary can
    be executed from anywhere in the system.</h6>
 3. Change file permission to allow binary execution with `chmod +x /usr/local/bin/sshauth`.
 4. Create a folder `sshauth` in the `/etc/`. The path will look like this `/etc/sshauth/`
 5. Copy the [config.yml](./resources/config.yml) file to this `/etc/sshauth/` (`/etc/sshauth/config.yml`) and configure
    to use your desired third-party service. [Configuration examples](./resources/docs/CONFIGURE.md) <h6>Owner of the
    <i>config.yml</i> must be same as the <i>sshauth</i> binary owner.</h6>
 6. Update the `/etc/ssh/sshd_config` to reflect the following changes:
    - `AuthorizedKeysCommand /usr/local/bin/sshauth %u %k`.
    - `AuthorizedKeysCommandUser root`. This differ on which user own the `sshauth` binary execution file.
 7. Apply the changes made to the SSH daemon using system specific command. I.E. `service ssh restart` for Ubuntu.
    It is recommended to [test](#test-installation-and-configuration) the installation before applying these changes.
    
### Test installation and configuration

#### Test case

 For testing the installation and configuration, run the following command:
 
 ```
 sshauth <USERNAME> "<PUBLIC KEY>"
 ```
   - `<USERNAME>` is the admin or default user of the server, defined under `system_conf.admin_user` or 
     `system_conf.default_user`
   - `<PUBLIC KEY>` is a public key from any user who's public key can be retrieved from the third-party service.
     In order to mimic the SSH daemon place the public key with in `" "` and exclude the trailing comment of the key.
     ```
     sshauth ubuntu "ssh-rsa AAAAB3Nza.......................F7XR2rshD/imqQ6BmjJgw0ejsL+xGj74I62GM3JdTWEcj5OgtHvPcZ6NOb"
     ```

#### Result

 With successful installation and configuration, you can expect following results:
 
 - List of public keys from the users that match the configuration criteria.
 - An alert to the given alerting service of a public key match from a login.
   
## Alerting

 Current version support alerting with [Slack Incoming Webhooks](https://api.slack.com/incoming-webhooks).
 ([HipChat Server]() and [HipChat Cloud]() will be added soon.)
 
 Alert will have basic formatting to provide the following information:
 
   - `User: UserName` - User name of the third-party service that public key matched
   - `Host: ServerName` - Server name defined under `system_conf.name`
   - `Service: ServiceUrl` - Url of the third-party service that has been used to authenticate the user
   
#### Alerting Example
 ```
 User: praveenprem
 Host: Test
 Service: github.com
 ```
   
## Logging

 This plugin will log errors and warning for debugging purposes, such as bad configuration. These logs can be found in
 `/var/log/sshauth.log` when the plugin is run as `root`.
 
 If the plugin is run as any another user, logs will use system's default login mechanism such as `system.log`.
 
#### Logging Example

```
2018/06/25 21:55:55 ERROR: Get http://gitlab.example.local/api/v4/groups?private_token=hgascvgJLSVBYLP: dial tcp: lookup gitlab.example.local: no such host
2018/06/25 21:55:55 ERROR: unable to locate group "Test". Please check the configuration
2018/06/25 21:57:06 ERROR: arguments missing or not provided
2018/06/25 21:57:36 ERROR: invalid user "bob"
```
 
## Resources

 - [Configuration Manual](./resources/docs/CONFIGURE.md)
 - [Exit Code Dictionary](./resources/docs/DICTIONARY.md)
 
## License

SSH Auth plugin is licensed under the MIT License, please refer to the [LICENSE](LICENSE) for more details.

## Versioning

For versions available, please refer to the [tags on this repository](https://github.com/praveenprem/sshauth/tags).

Source code for each release version can be found under the tag branch with `v` prefix. I.E `v1.0.0` for the version `1.0.0`

## Authors

 | <div><a href="https://github.com/praveenprem"><img width="200" src="https://avatars3.githubusercontent.com/u/23165760"/><p></p><p>Praveen Premaratne</p></a></div> |
 | :-------: |
 
## Acknowledgments

 - This README file has been structured according to [Billie Thompson](https://gist.github.com/PurpleBooth)'s [README-Template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
