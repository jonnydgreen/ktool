export async function showLogs(
  podName: string,
  watch: boolean,
  logger: (s: string) => void,
  podNamespace?: string
) {
  const execStdout = require('../command/execStdout').execStdout
  const args = `logs ${watch ? '-f' : ''} ${podName} ${
    podNamespace ? '-n ' + podNamespace : ''
  }`
    .split(' ')
    .filter((a: string) => a.trim() !== '')
  return execStdout('kubectl', args, logger)
}
