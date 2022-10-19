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
    return {err};
  }

  const ranForSplit = parts[2].split(';').map(v => v.trim()).filter(v => !!v);
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

    return {err};
  }

  const provider = ranForSplit[0];

  return {
    isDestroyFailedE2e: true,
    workflow: {
      ID: `e2e-clean-${provider}.yml`,
      targetRef: parts[1],
    },
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

/**
 * Build additional info about failed e2e test
 * Contains information about
 *
 * @param {object} jobs - GitHub needsContext context
 * @returns {string}
 */
function buildFailedE2eTestAdditionalInfo({ needsContext, core }){
  const connectStrings = Object.getOwnPropertyNames(needsContext).
  filter((k) => k.startsWith('run_')).
  map((key, _i, _a) => {
    const result = needsContext[key].result;
    if (result === 'failure' || result === 'cancelled') {
      if (needsContext[key].outputs){
        const outputs = needsContext[key].outputs;

        if(!outputs['failed_cluster_stayed']){
          return null;
        }

        // ci_commit_branch
        const connectStr = outputs['ssh_master_connection_string'] || '';
        const ranFor = outputs['ran_for'] || '';
        const runId = outputs['run_id'] || '';
        const artifactName = outputs['state_artifact_name'] || '';
        const stateDir = needsContext[key].outputs['state_dir'] || '';
        const branch = needsContext[key].outputs['ref'] || '';

        if (!branch || !stateDir || !ranFor || !connectStr || !artifactName || !runId) {
          core.warn(`Incorrect outputs for ${key}: ${JSON.stringify(outputs)}`)
        }

        const splitRunFor = ranFor.replace(';', ' ');

        return `E2e for ${splitRunFor} was failed. Use:
  \`ssh -i ~/.ssh/e2e-id-rsa ${connectStr}\` - connect for debugging;

  \`${abortFailedE2eCommand} ${branch} ${ranFor} ${runId} ${artifactName} ${stateDir}\` - for abort failed cluster
`
      }
    }

    return null;
  }).filter((v) => !!v)

  if (connectStrings.length === 0) {
    return "";
  }

  return "\r\n" + "#failed_clusters_start\r\n" + connectStrings.join("\r\n") + "\r\n#failed_clusters_end";
}

module.exports = {
  checkAbortE2eCluster,
  buildFailedE2eTestAdditionalInfo,
}
