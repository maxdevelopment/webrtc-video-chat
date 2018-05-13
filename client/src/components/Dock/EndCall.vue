<template>
  <div class="dock-btn end-call">
    <button type="button" @click="destroyPeer()">
      <icon name="power-off"></icon>
    </button>
  </div>
</template>

<script>
    import 'vue-awesome/icons/power-off'
    import Icon from 'vue-awesome/components/Icon'

    import SvgIcon from '@/components/SvgIcon'

    export default {
        name: 'EndCall',
        components: {
            SvgIcon,
            Icon
        },
        methods: {
            destroyPeer() {
                window.location.href = '/'
                localStorage.removeItem('room')
            }
        }
    }
</script>

<style lang="stylus">
  @import '../../styles'

  .end-call
    background transparentify($amour, .8)

    &:hover
      background transparentify($amour, .6) !important
</style>
