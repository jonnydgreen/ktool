import { GluegunToolbox } from 'gluegun'

module.exports = {
  name: 'namespace',
  alias: ['n'],
  description: 'Changes the current namespace of your kubernetes config',
  run: async (toolbox: GluegunToolbox) => {
    // retrieve the tools from the toolbox that we will need
    const { parameters, print, prompt, system } = toolbox

    // check if there's a name provided on the command line first
    // if not, let's prompt the user for one and then assign that to `newNamespace`
    let newNamespace
    if (parameters) {
      newNamespace = parameters.first
    }

    let availableNamespaces: string[] = (await system.run(
      `kubectl get ns -o=jsonpath='{.items[*].metadata.name}'`,
      { trim: true }
    )).split(' ')
    if (availableNamespaces.length === 0) {
      print.error('No namespaces to choose from')
      return
    } else if (newNamespace && !availableNamespaces.includes(newNamespace)) {
      print.error('Invalid namespace specified, toolketes is going back to bed')
      return
    }

    if (!newNamespace) {
      const result: { namespace?: string } = await prompt.ask({
        type: 'select',
        name: 'namespace',
        message: 'Which namespaceeroonie?',
        choices: availableNamespaces,
        default: availableNamespaces[0]
      })
      if (result && result.namespace) {
        newNamespace = result.namespace
      }
    }

    try {
      const currentNamespace = await system.run(
        'kubectl config current-context',
        { trim: true }
      )
      await system.run(
        `kubectl config set-context ${currentNamespace} --namespace=${newNamespace}`
      )
    } catch (err) {
      print.error(`Invalid namespace specified probz: ${err.message}`)
    }

    // success!
    print.divider()
    print.fancy(`Namespace changed to '${newNamespace}' you maniac!`)
    print.divider()
  }
}
