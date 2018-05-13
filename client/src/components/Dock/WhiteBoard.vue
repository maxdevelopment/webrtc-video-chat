<template>
  <div class="dock-btn screen">
    <button type="button" @click="showWhiteBoard()">
      <svg-icon src="line-chart"/>
    </button>
  </div>
</template>

<script>
  import SvgIcon from '@/components/SvgIcon'

  export default {
    name: "WhiteBoard",
    components: {
      SvgIcon
    },
    methods: {
      showWhiteBoard() {
        this.$store.dispatch('showWhiteBoard')
      }
    }
  }
</script>

<style scoped>

</style>
