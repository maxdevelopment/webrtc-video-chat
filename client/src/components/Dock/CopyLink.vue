<template>
  <div class="dock-btn copy-link">
    <button type="button" @click="copyLink()">
      <icon name="copy"></icon>
      <input ref="link" type="text" :value="link" />
    </button>
  </div>
</template>

<script>
    import 'vue-awesome/icons/copy'
    import Icon from 'vue-awesome/components/Icon'

    import SvgIcon from '@/components/SvgIcon'

    export default {
        name: 'CopyLink',
        data: () => ({
            link: null
        }),
        components: {
            SvgIcon,
            Icon
        },
        beforeMount() {
            this.link = window.location.href
        },
        methods: {
            copyLink() {
                this.$refs.link.value = this.link
                this.$refs.link.select()
                document.execCommand('copy')

                this.$message({
                    showClose: true,
                    message: 'Room link copied',
                    type: 'success'
                })
            }
        }
    }
</script>

<style lang="stylus">
  @import '../../styles'

  .copy-link
    background transparentify(#F4A460, .8)

    &:hover
     background transparentify(#DAA520, .6) !important

    input
      position absolute
      opacity 0
      bottom -100%
</style>
