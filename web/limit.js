//Execute this script through node before running the project (this script is at the same level as the node_modules directory)
const fs = require('fs')
const path = require('path')
const wfPath = path.resolve(__dirname, './node_modules/.bin')

fs.readdir(wfPath, (err, files) => {
  if (err) {
    console.log(err)
  } else {
    if (files.length !== 0) {
      files.forEach((item) => {
        if (item.split('.')[1] === 'cmd') {
          replaceStr(`${wfPath}/${item}`, /"%_prog%"/, '%_prog%')
        }
      })
    }
  }
})

// Parameters: [file path, string to be modified, modified string] (replace the public function corresponding to the string in the file)
function replaceStr(filePath, sourceRegx, targetSrt) {
  fs.readFile(filePath, (err, data) => {
    if (err) {
      console.log(err)
    } else {
      let str = data.toString()
      str = str.replace(sourceRegx, targetSrt)
      fs.writeFile(filePath, str, (err) => {
        if (err) {
          console.log(err)
        } else {
          console.log('\x1B[42m%s\x1B[0m', 'File modified successfully')
        }
      })
    }
  })
}
