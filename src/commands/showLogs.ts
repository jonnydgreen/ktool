module.exports = {
  name: 'logs',
  alias: ['l'],
  description: 'Shows the logs of a selected pod',
  run: async (toolbox) => {
    // retrieve the tools from the toolbox that we will need
    const GluegunToolbox = require('gluegun').GluegunToolbox
    const { parameters, print, prompt, system,  } = toolbox as typeof GluegunToolbox

    // check if there's a name provided on the command line first
    // if not, let's prompt the user for one and then assign that to `newContext`
    let selectedPod
    if (parameters) {
      selectedPod = parameters.first
    }

    const availablePods = (await system.run(
      `kubectl get pods -o jsonpath={.items[*].metadata.name}`
    , { trim: true }))
      .split(' ')

    if (availablePods.length === 0) {
      print.error('No contexts to choose from')
      return
    } else if (selectedPod && !availablePods.includes(selectedPod)) {
      print.error('Invalid context specified, toolketes is going back to bed')
      return
    }

    if (!selectedPod) {
      const result: { pod?: string } = await prompt.ask({
        type: 'autocomplete',
        name: 'pod',
        message: 'Which poderoonie?',
        choices: availablePods,
        default: availablePods[0]
      })
      if (result && result.pod) {
        selectedPod = result.pod
      }
    }

    const result: { rollingLogs?: boolean } = await prompt.ask({
      type: 'confirm',
      name: 'rollingLogs',
      message: 'Gloucester logs rolling?',
      initial: 'true'
    })

    try {

      console.log(selectedPod)
      await toolbox.kubectl.showLogs(selectedPod, result.rollingLogs, print.info);
    } catch (err) {
      print.error(`Pod not found probz: ${err.message}`)
    }

    // success!
    print.divider()
  }
}
