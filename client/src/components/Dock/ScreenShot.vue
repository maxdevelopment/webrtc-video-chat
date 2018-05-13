<template>
  <div class="dock-btn screen">
    <button type="button" @click="makeScreen()">
      <icon name="camera"></icon>
      <canvas ref="screenShow" class="screen__canvas" />
    </button>
  </div>
</template>

<script>
    import 'vue-awesome/icons/camera'
    import Icon from 'vue-awesome/components/Icon'

    import SvgIcon from '@/components/SvgIcon'

    export default {
        name: 'ScreenShot',
        components: {
            SvgIcon,
            Icon
        },
        methods: {
            makeScreen() {
                const canvas = this.$refs.screenShow

                canvas.width = window.innerWidth
                canvas.height = window.innerWidth / 1.4

                const video = document.getElementById('otherPic')
                const context = canvas.getContext('2d')

                context.drawImage(video, 0, 0, canvas.width, canvas.height)

                this.$store.dispatch('setImage', canvas.toDataURL())
            }
        }
    }
</script>

<style lang="stylus">
  @import '../../styles'

  .screen
    background transparentify(#ADFF2F, .8)

    &:hover
     background transparentify(#32CD32, .6) !important

    &__canvas
      position absolute
      width 100vw
      height 100vh
      top 0
      left 0
      opacity 0
      z-index -1
</style>
