<!-- Modification of this file is prohibited! If you have not purchased a license, please contact wx:shouzi_1994 to purchase a license. If you are not authorized, just keep this code and it will not affect any normal use -->

<template>
  <div class="flex gap-4 items-center text-sm text-gray-500 justify-center mb-4">
    <span>Powered by</span>
    <span>
      <a
        class="text-gray-800 font-bold"
        href="https://github.com/flipped-aurora/gin-vue-admin"
      >Gin-Vue-Admin</a>
    </span>
    <slot />
    <span>Copyright</span>
    <span>
      <a
        class="text-gray-800 font-bold"
        href="https://github.com/flipped-aurora"
      >flipped-aurora team</a>
    </span>
  </div>
</template>

<script  setup>
// The content of this file is copyright information. If you need to change it, please contact wx:shouzi_1994 to purchase authorization. If it is not authorized, just keep this code. It will not affect any normal use.
// The project is under the apatch protocol. Please abide by the copyright agreement.
defineOptions({
  name: 'BottomInfo'
})

console.log(
  `%c powered by %c flipped-aurorae %c`,
  'background:#0081ff; padding: 1px; border-radius: 3px 0 0 3px; color: #fff',
  'background:#354855; padding: 1px 5px; border-radius: 0 3px 3px 0; color: #fff; font-weight: bold;',
  'background:transparent'
)
</script>

