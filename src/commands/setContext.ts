import { GluegunToolbox } from 'gluegun'

module.exports = {
  name: 'context',
  alias: ['c'],
  description: 'Changes the current context of your kubernetes config',
  run: async (toolbox: GluegunToolbox) => {
    // retrieve the tools from the toolbox that we will need
    const { parameters, print, prompt, system } = toolbox

    // check if there's a name provided on the command line first
    // if not, let's prompt the user for one and then assign that to `newContext`
    let newContext
    if (parameters) {
      newContext = parameters.first
    }

    const availableContexts = (await system.run(
      'kubectl config get-contexts -o name'
    , { trim: true }))
      .split('\n')

    if (availableContexts.length === 0) {
      print.error('No contexts to choose from')
      return
    } else if (newContext && !availableContexts.includes(newContext)) {
      print.error('Invalid context specified, toolketes is going back to bed')
      return
    }

    if (!newContext) {
      const result: { context?: string } = await prompt.ask({
        type: 'select',
        name: 'context',
        message: 'Which contexteroonie?',
        choices: availableContexts,
        default: availableContexts[0]
      })
      if (result && result.context) {
        newContext = result.context
      }
    }

    try {
      await system.run(`kubectl config use-context ${newContext}`)
    } catch (err) {
      print.error(`Invalid context specified probz: ${err.message}`)
    }

    // success!
    print.divider()
    print.fancy(`Context changed to '${newContext}' you maniac!`)
    print.divider()
  }
}
