
import axios from 'axios'
import { ref } from 'vue'

const weatherInfo = ref('It will be sunny today, 0℃ - 10℃, the weather is cold, please bring more clothes.')
const amapKey = '8e8baa8a7317586c29ec694895de6e0a'

export const useWeatherInfo = () => {
  ip()
  return weatherInfo
}

export const ip = async() => {
  //Replace the key with your own https://console.amap.com/dev/index
  if (amapKey === '') {
    return false
  }
  const res = await axios.get('https://restapi.amap.com/v3/ip?key=' + amapKey)
  if (res.data.adcode) {
    getWeather(res.data.adcode)
  }
}

const getWeather = async(code) => {
  const response = await axios.get('https://restapi.amap.com/v3/weather/weatherInfo?key=' + amapKey + '&extensions=base&city=' + code)
  if (response.data.status === '1') {
    const s = response.data.lives[0]
    weatherInfo.value = s.city + 'Weather:' + s.weather + 'Temperature:' + s.temperature + 'Celsius Wind direction:' + s.winddirection + 'Wind power:' + s.windpower + 'Level air humidity: ' + s.humidity
  }
}

