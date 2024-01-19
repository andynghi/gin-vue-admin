export const downloadImage = (imgsrc, name) => { // Download image address and image name
  var image = new Image()
  image.setAttribute('crossOrigin', 'anonymous')
  image.onload = function() {
    var canvas = document.createElement('canvas')
    canvas.width = image.width
    canvas.height = image.height
    var context = canvas.getContext('2d')
    context.drawImage(image, 0, 0, image.width, image.height)
    var url = canvas.toDataURL('image/png') // Get the base64 encoded data of the image

    var a = document.createElement('a') // Generate an a element
    var event = new MouseEvent('click') // Create a click event
    a.download = name || 'photo' //Set the picture name
    a.href = url // Set the generated URL to the a.href attribute
    a.dispatchEvent(event) // Trigger the click event of a
  }
  image.src = imgsrc
}
