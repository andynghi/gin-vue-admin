# gin-vue-admin web 

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```


Organize code structure
``` lua
web
 ├── babel.config.js
 ├── Dockerfile
 ├── favicon.ico
 ├── index.html -- main page
 ├── limit.js -- helper code
 ├── package.json -- package manager code
 ├── src -- source code
 │ ├── api -- api group
 │ ├── App.vue -- Main page
 │ ├── assets -- static resources
 │ ├── components -- global components
 │ ├── core -- gva component package
 │ │ ├── config.js -- gva website configuration file
 │ │ ├── gin-vue-admin.js -- Registration welcome file
 │ │ └── global.js -- unified import file
 │ ├── directive -- v-auth registration file
 │ ├── main.js -- main file
 │ ├── permission.js -- routing middleware
 │ ├── pinia -- pinia state manager, replacing vuex
 │ │ ├── index.js -- entry file
 │   │   └── modules            -- modules
 │   │       ├── dictionary.js
 │   │       ├── router.js
 │   │       └── user.js
 │ ├── router -- routing declaration file
 │   │   └── index.js
 │ ├── style -- global style
 │   │   ├── base.scss
 │   │   ├── basics.scss
 │ │ ├── element_visiable.scss -- The element-plus style can be globally overridden here.
 │ │ ├── iconfont.css -- style files for the top icons
 │   │   ├── main.scss
 │   │   ├── mobile.scss
 │   │   └── newLogin.scss
 │ ├── utils -- method package library
 │ │ ├── asyncRouter.js -- dynamic routing related
 │ │ ├── bus.js -- global mitt declaration file
 │ │ ├── date.js -- date related
 │ │ ├── dictionary.js -- Get dictionary method
 │ │ ├── downloadImg.js -- How to download images
 │ │ ├── format.js -- formatting related
 │ │ ├── image.js -- Image related methods
 │ │ ├── page.js -- Set page title
 │ │ ├── request.js -- Request
 │ │ └── stringFun.js -- string file
 | ├── view -- main view code
 | | ├── about -- About us
 | | ├── dashboard -- panel
 | | ├── error -- error
 | | ├── example --upload case
 | | ├── iconList -- icon list
 | | ├── init -- initialization data
 | | | ├── index -- new version
 | | | ├── init -- old version
 | | ├── layout -- layout constraint page
 |   |   |   ├── aside 
 |   |   |   ├── bottomInfo     -- bottomInfo
 | | | ├── screenfull -- full screen settings
 | | | ├── setting -- system settings
 | | | └── index.vue -- base constraints
 | | ├── login --log in
 | | ├── person --personal center
 | | ├── superAdmin -- super administrator operation
 | | ├── system -- system detection page
 | | ├── systemTools -- System configuration related pages
 | | └── routerHolder.vue -- page entry page
 ├── vite.config.js -- vite configuration file
 └── yarn.lock

```
