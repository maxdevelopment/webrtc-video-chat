<template>
    <div class="dock">
        <el-tooltip id="tooltip5" effect="dark" content="Open whiteboard" placement="top">
            <white-board/>
        </el-tooltip>

        <div class="btn-right">
            <div class="btn-block">
                <el-tooltip id="tooltip1" effect="dark" content="Copy room link" placement="top">
                    <copy-link/>
                </el-tooltip>
            </div>

            <div class="btn-block">
                <el-tooltip id="tooltip2" effect="dark" content="Mute audio" placement="top">
                    <mute-stream/>
                </el-tooltip>
            </div>

            <div class="btn-block">
                <el-tooltip id="tooltip3" effect="dark" content="End call" placement="top">
                    <end-call/>
                </el-tooltip>
            </div>

            <div class="btn-block">
                <el-tooltip id="tooltip4" effect="dark" content="Make screenshot" placement="top">
                    <screen-shot/>
                </el-tooltip>
            </div>
        </div>
    </div>
</template>

<script>
import EndCall from './EndCall'
import MuteStream from './MuteStream'
import CopyLink from './CopyLink'
import ScreenShot from './ScreenShot'
import WhiteBoard from './WhiteBoard'

export default {
    name: 'Dock',
    components: {
        EndCall,
        MuteStream,
        CopyLink,
        ScreenShot,
        WhiteBoard
    }
}
</script>

<style lang="stylus">
  .dock
    font-weight bold
    position absolute
    width 100%
    bottom 12px
    left 0
    right 0
    margin auto
    padding 8px
    display flex
    flex-flow row nowrap
    background rgba(0, 0, 0, .5)
    border-radius 8px
    z-index 6

  .btn-block
    display inline-block

  .btn-right
    margin-left auto
    padding-right 20px

</style>
