const {abortFailedE2eCommand} = require("./constants");

/**
 *
 * @param {string[]} parts - slash command strings
 * @return {object}
 */
function checkAbortE2eCluster(parts){
  const command = parts[0];
  if (command !== abortFailedE2eCommand) {
    return null;
  }

  if (parts.length !== 6) {
    let err = 'clean failed e2e cluster should have 6 arguments'
    switch (parts.length){
      case 5:
        err = 'state dir is required';
        break;
      case 4:
        err = 'artifact name and state dir are required';
        break;
      case 3:
        err = 'run id, artifact name and state dir are required';
        break;
      case 2:
        err = 'ran for (provider, layout, cri, k8s version), run id, artifact name, and state dir are required';
        break;
      case 1:
        err = 'full_ref, ran for (provider, layout, cri, k8s version), run id, artifact name, and state dir are required';
        break;
    }
    return {notFoundMsg: err};
  }

  const ranForSplit = parts[2].split(',').map(v => v.trim()).filter(v => !!v);
  if (ranForSplit.length !== 4) {
    let err = '"ran for" argument should have 4 parts';
    switch (ranForSplit.length) {
      case 3:
        err = 'k8s version is required';
        break;
      case 2:
        err = 'cri and k8s version are required';
        break;
      case 1:
        err = 'layout, cri and k8s version are required';
        break;
      case 0:
        err = 'provider, layout, cri and k8s version are required';
        break;
    }

    return {notFoundMsg: err};
  }

  const provider = ranForSplit[0];

  return {
    isDestroyFailedE2e: true,
    workflow_id: `e2e-clean-${provider}.yml`,
    targetRef: parts[1],
    inputs: {
      run_id: parts[3],
      state_artifact_name: parts[4],
      state_dir: parts[5],

      layout: ranForSplit[1],
      cri: ranForSplit[2],
      k8s_version: ranForSplit[3],
    },
  };
}


module.exports = {
  checkAbortE2eCluster,
}
