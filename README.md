
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.20-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.1-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.3.4-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.3.8-green"/>
<img src="https://img.shields.io/badge/gorm-1.25.2-red"/>
</div>

[English](./README-en.md) | Simplified Chinese

## different versions

We will continue to maintain the following four versions, please choose the version that suits you. The latest technology stack is a combined API version that supports multiple languages (I18N)

[Combined API version (main)](https://github.com/flipped-aurora/gin-vue-admin) |
[Combined API multi-language (i18n) version](https://github.com/flipped-aurora/gin-vue-admin/tree/i18n-dev-new) |
[Declarative API version](https://github.com/flipped-aurora/gin-vue-admin/tree/v2.4.x) |
[Declarative API multi-language (i18n) version](https://github.com/flipped-aurora/gin-vue-admin/tree/i18n-dev)

#Project documentation
[Online Documentation](https://www.gin-vue-admin.com): https://www.gin-vue-admin.com

[Initialization](https://www.gin-vue-admin.com/guide/start-quickly/initialization.html)
						       
[From environment to deployment teaching video](https://www.bilibili.com/video/BV1Rg411u7xH)

[Development Tutorial](https://www.gin-vue-admin.com/guide/start-quickly/env.html) (Contributor: <a href="https://github.com/LLemonGreen">LLemonGreen </a> And <a href="https://github.com/fkk0509">Fann</a>)

[Communication Community](https://support.qq.com/products/371961)

[Plugin Market](https://plugin.gin-vue-admin.com/)

# important hint

1. This project has documentation and detailed video tutorials from start-up to development to deployment.

2. This project requires you to have a certain foundation in golang and vue

3. You can complete all operations through our tutorials and documents, so we no longer provide free technical services. If you need services, please go to [paid support] (https://www.gin-vue-admin.com/coffee /payment.html)

4. If you use this project for commercial purposes, please comply with the Apache2.0 agreement and retain the author's technical support statement. You need to retain the following copyright statement information, and there are no restrictions on other information functions. If you need to remove it, please [Purchase Authorization](https://www.gin-vue-admin.com/empower/index.html)

<img src="https://qmplusimg.henrongyi.top/%E6%8E%88%E6%9D%83.png" width="1000">

## 1. Basic introduction

### 1.1 Project Introduction

> Gin-vue-admin is a full-stack front-end and back-end development basic platform developed based on [vue](https://vuejs.org) and [gin](https://gin-gonic.com), integrating jwt Authentication, dynamic routing, dynamic menu, casbin authentication, form generator, code generator and other functions provide a variety of sample files, allowing you to focus more time on business development.

[Online preview](http://demo.gin-vue-admin.com): http://demo.gin-vue-admin.com

Test username: admin

Test password: 123456

### 1.2 Contribution Guidelines
Hi! First of all, thank you for using gin-vue-admin.

Gin-vue-admin is a set of open source frameworks with a front-end and back-end separation architecture prepared for rapid research and development. It is designed to quickly build small and medium-sized projects.

The growth of Gin-vue-admin cannot be separated from everyone's support. If you are willing to contribute code or provide suggestions for gin-vue-admin, please read the following content.

#### 1.2.1 Issue specification
- Issues are only used to submit bugs or features and design-related content. Other content may be closed directly.
									      
- Before submitting an issue, please search to see if the relevant content has been raised.

#### 1.2.2 Pull Request Specification
- Please fork a copy to your own project first, do not create a branch directly under the warehouse.

- The commit information should be filled in in the form of `[file name]: description information`, for example `README.md: fix xxx bug`.

- If you are fixing a bug, please provide a description in the PR.

- Merging the code requires the participation of two maintainers: one person reviews and then approves, and the other person reviews again, and can be merged after approval.

## 2. Instructions for use

```
- node version > v16.8.3
- golang version >= v1.16
- IDE recommendation: Goland
```

### 2.1 server project

Use editing tools such as `Goland` to open the server directory. Do not open the gin-vue-admin root directory.

```bash

# Clone project
git clone https://github.com/flipped-aurora/gin-vue-admin.git
# Enter the server folder
cd server

# Use go mod and install go dependency packages
go generate

# compile
go build -o server main.go (the windows compilation command is go build -o server.exe main.go)

# Run binary
./server (Windows running command is server.exe)
```

### 2.2 web project

```bash
# Enter the web folder
cd web

# Install dependencies
npm install

# Start web project
npm run serve
```

### 2.3 swagger automation API documentation

#### 2.3.1 Install swagger

``` shell
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 2.3.2 Generate API documentation

```` shell
cd server
swag init
````

> After executing the above command, the three files `docs.go`, `swagger.json`, and `swagger.yaml` in the docs folder will appear in the server directory. After starting the go service, enter [http in the browser ://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) to view the swagger documentation

### 2.4 VSCode workspace

#### 2.4.1 Development

Use `VSCode` to open the workspace file `gin-vue-admin.code-workspace` in the root directory. You can see three virtual directories in the sidebar: `backend`, `frontend`, and `root`.

#### 2.4.2 Run/Debug

You can also see three tasks in running and debugging: `Backend`, `Frontend`, and `Both (Backend & Frontend)`. Run `Both (Backend & Frontend)` to start the front-end and front-end projects at the same time.

#### 2.4.3 settings

There is a `go.toolsEnvVars` field in the workspace configuration file, which is the go tool environment variable used for `VSCode` itself. In addition, in systems with multiple go versions, you can specify the running version through `gopath` and `go.goroot`.

```json
    "go.gopath": null,
    "go.goroot": null,
```

## 3. Technology selection

- Front-end: Build the basic page with [Element](https://github.com/ElemeFE/element) based on [Vue](https://vuejs.org).
- Backend: Use [Gin](https://gin-gonic.com/) to quickly build a basic restful style API. [Gin](https://gin-gonic.com/) is a web framework written in go language. .
- Database: Use `MySql` > (5.7) version database engine InnoDB, use [gorm](http://gorm.cn) to implement basic operations on the database.
- Cache: Use `Redis` to record the `jwt` token of the current active user and implement multi-point login restrictions.
- API documentation: Use `Swagger` to build automation documentation.
- Configuration file: Use [fsnotify](https://github.com/fsnotify/fsnotify) and [viper](https://github.com/spf13/viper) to implement the `yaml` format configuration file.
- Log: Use [zap](https://github.com/uber-go/zap) to implement logging.

## 4. Project structure

### 4.1 System Architecture Diagram

![System Architecture Diagram](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 Front-end detailed design drawing (provided by: <a href="https://github.com/baobeisuper">baobeisuper</a>)

![Front-end detailed design drawing](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 Directory structure

```
    ├── server
        ├── api (api layer)
        │ └── v1 (v1 version interface)
        ├── config (configuration package)
        ├── core (core file)
        ├── docs (swagger document directory)
        ├── global (global object)
        ├── initialize (initialization)
        │ └── internal (initialization internal function)
        ├── middleware (middleware layer)
        ├── model (model layer)
        │ ├── request (input parameter structure)
        │ └── response (parameter structure)
        ├── packfile (static file packaging)
        ├── resource (static resource folder)
        │ ├── excel (default path for excel import and export)
        │ ├── page (form builder)
        │ └── template
        ├── router (routing layer)
        ├── service (service layer)
        ├── source (source layer)
        └── utils (toolkit)
            ├── timer (timer interface package)
            └── upload (oss interface encapsulation)
    
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
        │ │ ├── btnAuth.js -- Dynamic permission button related
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

## 5. Main functions

- Permission management: Permission management based on `jwt` and `casbin`.
- File upload and download: implement file upload operations based on `Qiniu Cloud`, `Alibaba Cloud`, `Tencent Cloud` (please develop your own application for `token` or `key` for each platform).
- Paging encapsulation: The front end uses `mixins` to encapsulate paging, and the paging method calls `mixins`.
- User management: System administrator assigns user roles and role permissions.
- Role management: Create the main object for permission control, and assign different API permissions and menu permissions to roles.
- Menu management: realize user dynamic menu configuration and realize different menus for different roles.
- API management: Different users have different permissions on the API interfaces they can call.
- Configuration management: Configuration files can be modified at the front desk (this function is not available on the online experience site).
- Conditional search: Add conditional search examples.
- restful example: You can refer to the sample API in the user management module.
	- Front-end file reference: [web/src/view/superAdmin/api/api.vue](https://github.com/flipped-aurora/gin-vue-admin/blob/master/web/src/view/superAdmin /api/api.vue)
    - Backend file reference: [server/router/sys_api.go](https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/router/sys_api.go)
- Multipoint login restrictions: You need to change `use-multipoint` in `system` to true in `config.yaml` (you need to configure Redis and Redis parameters in Config yourself. During the testing phase, please report any bugs in time).
- Multipart upload: Provides examples of multipart file upload and multipart upload of large files.
- Form builder: The form builder uses [@Variant Form](https://github.com/vform666/variant-form).
- Code generator: background basic logic and simple curd code generator.

## 6. Knowledge Base

## 6.1 Team Blog

> https://www.yuque.com/flipped-aurora
>
>Contains front-end framework teaching videos. If you think the project is helpful to you, you can add my personal WeChat: shouzi_1994. You are welcome to make valuable requests.

## 6.2 Teaching video

(1) Step-by-step instruction video

> https://www.bilibili.com/video/BV1Rg411u7xH/

(2) Introduction to back-end directory structure adjustment and how to use it

> https://www.bilibili.com/video/BV1x44y117TT/

(3) golang basic teaching video

> bilibili：https://space.bilibili.com/322210472/channel/detail?cid=108884

(4) Basic teaching of gin framework

> bilibili：https://space.bilibili.com/322210472/channel/detail?cid=126418&ctype=0

(5) gin-vue-admin version update introduction video

> bilibili：https://www.bilibili.com/video/BV1kv4y1g7nT

## 7. Contact information

### 7.1 Technical Group

### QQ communication group: 622360840
| QQ group |
|  :---:  |
| <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

### WeChat communication group
| WeChat |
|  :---:  | 
| <img width="150" src="http://qmplusimg.henrongyi.top/qrjjz.png"> 

To prevent advertisements from entering the group, add WeChat and enter the following code execution results (do not convert to string)

```
str := "5Yqg5YWlR1ZB5Lqk5rWB576k"
decodeBytes, err := base64.StdEncoding.DecodeString(str)
fmt.Println(decodeBytes, err)
```

### [About us](https://www.gin-vue-admin.com/about/join.html)

## 8. Contributors

Thank you for your contribution to gin-vue-admin!

<a href="https://github.com/flipped-aurora/gin-vue-admin/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=flipped-aurora/gin-vue-admin" />
</a>

## 9. Donate

If you think this project is helpful to you, you can buy the author a drink :tropical_drink: [Click me](https://www.gin-vue-admin.com/coffee/index.html)

## 10. Precautions for commercial use

If you use this project for commercial purposes, please comply with the Apache2.0 license and retain the author's technical support statement.
