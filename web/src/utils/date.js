// Extension of Date, convert Date into String in specified format
// Month (M), day (d), hour (h), minute (m), second (s), quarter (q) can use 1-2 placeholders,
// Year (y) can use 1-4 placeholders, milliseconds (S) can only use 1 placeholder (a number with 1-3 digits)
// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
// eslint-disable-next-line no-extend-native
Date.prototype.Format = function(fmt) {
  var o = {
    'M+': this.getMonth() + 1, // month
    'd+': this.getDate(), // day
    'h+': this.getHours(), // hours
    'm+': this.getMinutes(), // minutes
    's+': this.getSeconds(), // seconds
    'q+': Math.floor((this.getMonth() + 3) / 3), // quarter
    'S': this.getMilliseconds() // Milliseconds
  }
  if (/(y+)/.test(fmt)) { fmt = fmt.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length)) }
  for (var k in o) {
    if (new RegExp('(' + k + ')').test(fmt)) { fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length))) }
  }
  return fmt
}

export function formatTimeToStr(times, pattern) {
  var d = new Date(times).Format('yyyy-MM-dd hh:mm:ss')
  if (pattern) {
    d = new Date(times).Format(pattern)
  }
  return d.toLocaleString()
}
