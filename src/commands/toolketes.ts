module.exports = {
  name: 'toolketes',
  run: async toolbox => {
    const GluegunToolbox = require('gluegun').GluegunToolbox
    const { print } = toolbox as typeof GluegunToolbox

    print.info('Welcome to your CLI')
  }
}
