module.exports = toolbox => {
  const showLogs = require('../toolbox/kubectl/showLogs').showLogs;
  toolbox.kubectl = {
    showLogs
  }
}