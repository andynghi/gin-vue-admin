## GVA email sending function plug-in
#### Developer: GIN-VUE-ADMIN Official

### Steps for usage

#### 1. Go to initialize/router.go under the GVA main program and add this plug-in at the end of the Routers method according to the security mode you need.
    example:
    This plug-in can use the configuration file of gva, or you can directly write the content as the configuration. It is recommended to add the configuration file structure to gva and then pass the configuration in.
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
		))

    You can also write it in when passing it in

    PluginInit(PrivateGroup, email.CreateEmailPlug(
    "a@qq.com",
    "b@qq.com",
    "smtp.qq.com",
    "global.GVA_CONFIG.Email.Secret",
    "Login Key",
    465,
    true,
    ))

### 2. Configuration instructions

#### 2-1 Global configuration structure description
    //The Form and Secret are usually the username and password

    type Email struct {
	    To string // Recipients: Multiple separated by English commas. Example: a@qq.com b@qq.com Please use this item as a parameter during official development. The configuration here is mainly used to send error monitoring emails.
	    From string // Sender: the email address you want to send the email to
	    Host string // Server address For example smtp.qq.com Please go to QQ or the email address you want to send emails to check its SMTP protocol
	    Secret string // Key The key used for login. It is best not to use the email password. Go to the email SMTP to apply for a key used for login.
	    Nickname string // Nickname Sender's nickname can be customized and can be left blank
	    Port int // Port. Please go to QQ or the email address you want to send emails to check its SMTP protocol. Most of them are 465.
	    IsSSL bool // Whether to SSL or whether to enable SSL
    }
#### 2-2 Input parameter structure description
    //The Form and Secret are usually the username and password

    type Email struct {
        To string `json:"to"` // Who to send the email to
        Subject string `json:"subject"` // Email title
        Body string `json:"body"` // Email content
    }


### 3. Method API

    utils.EmailTest(email title, email body) Send test email
    Example: utils.EmailTest("Test Email", "Test Email")
    utils.ErrorToEmail(mail title, mail body) error monitoring
    Example: utils.ErrorToEmail("Test Email", "Test Email")
    utils.Email (if there are multiple target mailboxes, separate them with commas, email title, email body) Send test email
    Example: utils.Email("a.qq.com,b.qq.com","Test Email","Test Email")

### 4. Directly callable interface

    Test interface: /email/emailTest [post] swagger configured

    Sending email interface: /email/emailSend [post] swagger configured
    Input parameters:
    type Email struct {
        To string `json:"to"` // Who to send the email to
        Subject string `json:"subject"` // Email title
        Body string `json:"body"` // Email content
    }
   
