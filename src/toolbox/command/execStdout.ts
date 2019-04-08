export async function execStdout(
  cmd: string,
  args: string[],
  logger: (s: string) => void
) {
  return new Promise((resolve, reject) => {
    let spawn = require('child_process').spawn
    const exec = spawn(cmd, args)
    let buffer = ''

    exec.stdout.on('data', function(data) {
      logger(data.toString())
      buffer += data.toString()
    })

    exec.stderr.on('data', function(data) {
      logger(data.toString())
    })

    exec.on('error', function(err) {
      return reject(err)
    })

    exec.on('exit', function() {
      return resolve(buffer)
    })
  })
}
