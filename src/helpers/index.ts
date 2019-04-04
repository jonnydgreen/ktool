import * as cp from 'child_process'

export function exec(cmd: string): Promise<string> {
  return new Promise((resolve, reject) => {
    cp.exec(cmd, (err, stdout, stderr) => {
      if (err) {
        return reject(err)
      }
      return resolve(stdout)
    })
  })
}
