/*
                    Commercial code company’s own products do not require authorization
    This script must be retained if the product is sold as code (anything involving delivery of code to a third party for subsequent development)
                         Or mark the original author information
                          Otherwise, rights will be safeguarded according to law
*/

var child_process = require('child_process')

var url = 'https://www.gin-vue-admin.com'
var cmd = ''
console.log(process.platform)
switch (process.platform) {
  case 'win32':
    cmd = 'start'
    child_process.exec(cmd + ' ' + url)
    break

  case 'darwin':
    cmd = 'open'
    child_process.exec(cmd + ' ' + url)
    break
}
