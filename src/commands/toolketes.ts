module.exports = {
  name: 'toolketes',
  description: 'Toolketes CLI welcome screen',
  run: async toolbox => {
    const GluegunToolbox = require('gluegun').GluegunToolbox
    const { print } = toolbox as typeof GluegunToolbox

    print.info('Welcome to toolketes')
  }
}
