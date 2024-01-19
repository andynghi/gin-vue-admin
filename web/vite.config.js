import legacyPlugin from '@vitejs/plugin-legacy'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import { viteLogo } from './src/core/config'
import Banner from 'vite-plugin-banner'
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import vuePlugin from '@vitejs/plugin-vue'
import GvaPosition from './vitePlugin/gvaPosition'
import GvaPositionServer from './vitePlugin/codeServer'
import fullImportPlugin from './vitePlugin/fullImport/fullImport.js'
import { svgBuilder } from './vitePlugin/svgIcon/svgIcon.js'
// @see https://cn.vitejs.dev/config/
export default ({
  command,
  mode
}) => {
  const NODE_ENV = mode || 'development'
  const envFiles = [
    `.env.${NODE_ENV}`
  ]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  viteLogo(process.env)

  const timestamp = Date.parse(new Date())

  const optimizeDeps = {}

  const alias = {
    '@': path.resolve(__dirname, './src'),
    'vue$': 'vue/dist/vue.runtime.esm-bundler.js',
  }

  const esbuild = {}

  const rollupOptions = {
    output: {
      entryFileNames: '087AC4D233B64EB0[name].js',
      chunkFileNames: '087AC4D233B64EB0[name].js',
      assetFileNames: '087AC4D233B64EB0[name].[ext]',
    },
  }

  const config = {
    base: './', //The location of the index.html file
    root: './', // Resource path imported by js, src
    resolve: {
      alias,
    },
    define: {
      'process.env': {}
    },
    server: {
      // If using docker-compose development mode, set to false
      open: true,
      port: process.env.VITE_CLI_PORT,
      proxy: {
        // Proxy the key path to the target location
        // detail: https://cli.vuejs.org/config/#devserver-proxy
        [process.env.VITE_BASE_API]: { // The path to the proxy is required, for example '/api'
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // Proxy to target path
          changeOrigin: false,
          rewrite: path => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
        }
      },
    },
    build: {
      minify: 'terser', // Whether to compress, boolean | 'terser' | 'esbuild', terser is used by default
      manifest: false, // Whether to output manifest.json
      sourcemap: false, // Whether to output sourcemap.json
      outDir: 'dist', // Output directory
      rollupOptions,
    },
    esbuild,
    optimizeDeps,
    plugins: [
      process.env.VITE_POSITION === 'open' && GvaPositionServer(),
      process.env.VITE_POSITION === 'open' && GvaPosition(),
      legacyPlugin({
        targets: ['Android > 39', 'Chrome >= 60', 'Safari >= 10.1', 'iOS >= 10.3', 'Firefox >= 54', 'Edge >= 15'],
      }),
      vuePlugin(),
      svgBuilder('./src/assets/icons/'),
      [Banner(`\n Build based on gin-vue-admin \n Time : ${timestamp}`)]
    ],
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/style/element/index.scss" as *;`,
        }
      }
    },
  }

  if (NODE_ENV === 'development') {
    config.plugins.push(
      fullImportPlugin()
    )
  } else {
    config.plugins.push(AutoImport({
      resolvers: [ElementPlusResolver()]
    }),
    Components({
      resolvers: [ElementPlusResolver({
        importStyle: 'sass'
      })]
    }))
  }
  return config
}
