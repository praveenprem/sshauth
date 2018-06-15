#**Exit Code Dictionary**

###

Following table lists exit codes used by this program when a fatal error occurs.


Exit Code   | Description
:---------: | :-----------:
2           | No such file or directory
5           | I/O error
7           | Arg list too long
22          | Invalid argument
61          | No data available
70          | Communication error on send

##

### Scenarios these codes are used

#### Exit Code `2`

- Fail to open the `config.yml` file
- Fail to parse the content of `config.yml`. Such as an invalid syntax issue

#### Exit Code `5`

- Fail to parse JSON body from Http reposes
- Unable to open log file for logging.

#### Exit Code `7`

- Too many arguments passed on execution

#### Exit Code `22`

- Arguments are missing on execution
- Argument not matching the values in configuration `config.yam` file.

   Invalid system username
    
#### Exit Code `61`

- Invalid configuration in the `config.yam` file.
- Missing configuration in `config.yam` file.
- Service response mismatch with configuration in `config.yam` file.

    - No matching Team in GitHub

#### Exit Code `70`

- Https connection error

    - GitHub authentication service API returned error