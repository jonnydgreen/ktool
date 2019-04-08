module.exports = {
  name: 'all-pods',
  alias: ['a'],
  description: 'Get all pods',
  run: async (toolbox) => {
    const GluegunToolbox = require('gluegun').GluegunToolbox
    const { print, prompt, system } = toolbox as typeof GluegunToolbox

    const result: { watchAllPods?: boolean } = await prompt.ask({
      type: 'confirm',
      name: 'watchAllPods',
      message: 'Watch podtoon network?',
      initial: 'true'
    })
    let watchAllPods = false
    if (result && result.watchAllPods) {
      watchAllPods = result.watchAllPods
    }

    let allPods = await system.run(`kubectl get pods --all-namespaces`, {
      trim: true
    })
    const spinner = print.spin(allPods)
    try {
      while (watchAllPods) {
        await toolbox.system.run('sleep 1')
        allPods = await system.run(`kubectl get pods --all-namespaces`, {
          trim: true
        })
        spinner.text = allPods
        spinner.render()
      }
      spinner.stop()
      print.divider()
      print.info(allPods)
      print.divider()
    } catch (err) {
      spinner.fail(`Something went wrong: ${err.message}`)
    }
  }
}
