<template>
  <div 
    class="dock-btn mute-stream"
    :class="{
      'dock-btn--active': active
    }"
  >
    <button type="button" @click="muteStrem()">
      <div style="margin-left: 3px">
        <icon v-if="!active" name="microphone"></icon>
      </div>
      <icon v-if="active" name="microphone-slash"></icon>
    </button>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

import 'vue-awesome/icons/microphone'
import 'vue-awesome/icons/microphone-slash'
import Icon from 'vue-awesome/components/Icon'

export default {
  name: 'MuteStream',
  data: () => ({
    active: false
  }),
  components: {
    SvgIcon,
      Icon
  },
  methods: {
    muteStrem() {
      this.active = !this.active
      this.currentPeer.stream.getAudioTracks()[0].enabled = !this.active
    }
  },
  computed: {
    currentPeer() {
      return this.$store.getters.currentPeer
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .mute-stream
    background transparentify(#00FFFF, .8)

    &:hover
     background transparentify(#00BFFF, .6) !important
</style>
