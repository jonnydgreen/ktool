
import { GluegunToolbox } from 'gluegun'


module.exports = {
  name: 'toolketes',
  run: async (toolbox: GluegunToolbox) => {
    const { print } = toolbox

    print.info('Welcome to your CLI')
  },
}
