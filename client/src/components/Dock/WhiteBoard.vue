<template>
  <div class="btn-lft">
    <a type="button" @click="showWhiteBoard()" class="button-lft" :class="{
    'btn-lft-active': !activeButton
    }">
      Camera
    </a>
    <a type="button" @click="showWhiteBoard()" class="button-lft" :class="{
    'btn-lft-active': activeButton
    }">
      Whiteboard
    </a>
    <a type="button" class="button-lft">
      Documents
    </a>
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
        },
        computed: {
            activeButton() {
                return this.$store.getters.whiteBoard.visibility
            }
        }
    }
</script>

<style lang="stylus">
  .btn-lft
    text-transform uppercase
    font-size 1.2em
    padding-top 8px

  .btn-lft-active
    color #87CEEB
    background-color #87CEEB

  .button-lft
    display inline-block
    padding 0.35em 1.2em
    margin 0 0.3em 0.3em 0
    border-radius 0.12em
    box-sizing border-box
    text-decoration none
    color #FFFFFF;
    text-align center
    transition all 0.2s
    &:hover
      color #000000
      background-color #FFFFFF

</style>
