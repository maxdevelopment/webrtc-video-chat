DOM = require '../reactGUI/ReactDOMFactories-shim'
createReactClass = require '../reactGUI/createReactClass-shim'
createSetStateOnEventMixin = require './createSetStateOnEventMixin'
{_} = require '../core/localization'
{classSet} = require '../core/util'

ClearButton = createReactClass
  displayName: 'ClearButton'
  getState: -> {isEnabled: @props.lc.canUndo()}
  getInitialState: -> @getState()
  mixins: [createSetStateOnEventMixin('drawingChange')]

  render: ->
    {div} = DOM
    {lc} = @props

    className = classSet
      'lc-clear': true
      'toolbar-button': true
      'fat-button': true
      'disabled' : false
      ###'disabled': not @state.isEnabled###
    onClick = (=> lc.clear())

    (div {className, onClick}, _('Clear'))


module.exports = ClearButton
